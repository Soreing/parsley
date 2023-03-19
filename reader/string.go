package reader

import (
	"time"
)

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
	if r.pos >= len(r.dat) {
		return nil, NewEndOfFileError()
	} else if r.dat[r.pos] != '"' {
		return nil, NewInvalidCharacterError(r.dat[r.pos], r.pos)
	}

	r.pos++
	beg, end, esc := r.pos, 0, false

	for {
		if r.pos == len(r.dat) {
			return nil, NewEndOfFileError()
		} else {
			switch r.dat[r.pos] {
			case '\\':
				esc = !esc
			case '"':
				if !esc {
					end = r.pos
					r.pos++
					r.skipWhiteSpace()
					return (r.dat)[beg:end], nil
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

func (r *Reader) stringSeq(idx int) (res []string, err error) {
	var bs []byte
	if bs, err = r.GetByteArray(); err == nil {
		if r.Next() {
			res, err = r.stringSeq(idx + 1)
		} else {
			res = make([]string, idx+1)
		}
		if err == nil {
			res[idx] = string(bs)
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
