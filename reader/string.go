package reader

import (
	"time"
)

func strTokLen(dat []byte) int {
	esc, ecnt := false, 0
	for i, e := range dat {
		if esc {
			esc = false
			continue
		}
		if e == '\\' {
			ecnt++
			esc = true
			continue
		}
		if e == '"' {
			return i - ecnt
		}
	}
	return -1
}

func (r *Reader) skipString() error {
	esc := false
	r.pos++

	for {
		if r.pos >= len(r.dat) {
			return NewEndOfFileError()
		} else {
			switch r.dat[r.pos] {
			case '\\':
				esc = !esc
			case '"':
				if !esc {
					r.pos++
					return nil
				} else {
					esc = false
				}
			default:
				esc = false
			}
		}
		r.pos++
	}
}

func (r *Reader) GetByteArray() ([]byte, error) {
	dat, pos := r.dat[r.pos:], 0
	buf, bi := r.buf, 0
	esc := false

	if len(dat) == 0 {
		return nil, NewEndOfFileError()
	} else if dat[0] != '"' {
		return nil, NewInvalidCharacterError(dat[pos], r.pos+pos)
	}

	for pos++; ; pos++ {
		if pos >= len(dat) {
			r.pos += pos + 1
			return nil, NewEndOfFileError()

		} else if dat[pos] == '"' {
			r.pos += pos + 1
			r.SkipWhiteSpace()
			return dat[1:pos], nil

		} else if dat[pos] == '\\' {
			if len(buf) < pos {
				if tlen := strTokLen(dat[pos:]); tlen == -1 {
					return nil, NewEndOfFileError()
				} else {
					len := pos + tlen - 1
					if len < 256 {
						len = 256
					}

					r.buf = make([]byte, len)
					buf = r.buf
				}
			}
			copy(buf, dat[1:pos])
			esc, bi = true, pos-1
			break
		}
	}

	for pos++; ; {
		if esc {
			esc = false
			switch dat[pos] {
			case '"':
				buf[bi] = '"'
			case '\\':
				buf[bi] = '\\'
			case '/':
				buf[bi] = '/'
			case 'b':
				buf[bi] = '\b'
			case 'f':
				buf[bi] = '\f'
			case 'n':
				buf[bi] = '\n'
			case 'r':
				buf[bi] = '\r'
			case 't':
				buf[bi] = '\t'
			case 'u':
				buf[bi] = ' ' // TODO
			}
			bi++
		} else {
			if pos >= len(dat) {
				r.pos += pos + 1
				return nil, NewEndOfFileError()
			} else if dat[pos] == '"' {
				r.pos += pos + 1
				r.SkipWhiteSpace()
				return buf[:bi], nil
			} else if bi == len(buf) {
				if tlen := strTokLen(dat[pos:]); tlen == -1 {
					return nil, NewEndOfFileError()
				} else {
					r.buf = make([]byte, bi+tlen)
					copy(r.buf, buf[:bi])
					buf = r.buf
					continue
				}
			} else if dat[pos] == '\\' {
				esc = true
			} else {
				buf[bi] = dat[pos]
				bi++
			}
		}
		pos++
	}
}

func (r *Reader) stringSeq(idx int) (res []string, err error) {
	var bs []byte
	if bs, err = r.GetByteArray(); err == nil {
		s := string(bs)
		if !r.Next() {
			res = make([]string, idx+1)
			res[idx] = s
		} else if res, err = r.stringSeq(idx + 1); err == nil {
			res[idx] = s
		}
	}
	return
}

func (r *Reader) GetStrings() (res []string, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.stringSeq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetString() (string, error) {
	if bs, err := r.GetByteArray(); err != nil {
		return "", err
	} else {
		return string(bs), nil
	}
}

func (r *Reader) GetStringPtr() (res *string, err error) {
	if v, err := r.GetString(); err == nil {
		res = &v
	}
	return
}

func (r *Reader) timeSeq(idx int) (res []time.Time, err error) {
	var bs []byte
	if bs, err = r.GetByteArray(); err == nil {
		if r.Next() {
			res, err = r.timeSeq(idx + 1)
		} else {
			res = make([]time.Time, idx+1)
		}
		if err == nil {
			res[idx], err = time.Parse(time.RFC3339, string(bs))
		}
	}
	return
}

func (r *Reader) GetTimes() (res []time.Time, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.timeSeq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetTime() (time.Time, error) {
	if bs, err := r.GetByteArray(); err != nil {
		return time.Time{}, err
	} else {
		return time.Parse(time.RFC3339, string(bs))
	}
}

func (r *Reader) GetTimePtr() (res *time.Time, err error) {
	if v, err := r.GetTime(); err == nil {
		res = &v
	}
	return
}
