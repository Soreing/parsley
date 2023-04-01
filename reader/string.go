package reader

var hexValue = [128]int{
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	0., 1., 2., 3., 4., 5., 6., 7., 8., 9., -1, -1, -1, -1, -1, -1,
	-1, 10, 11, 12, 13, 14, 15, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, 10, 11, 12, 13, 14, 15, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
}

func strTokLen(dat []byte) int {
	esc, cnt, e := false, 0, byte(0)
	for i := 0; i < len(dat); i++ {
		if e = dat[i]; esc {
			if e != 'u' {
				cnt++
			} else if i+4 >= len(dat) {
				return -1
			} else if dat[i+2] >= '8' {
				cnt += 3
			} else if dat[i+3] >= '8' {
				cnt += 4
			} else {
				cnt += 5
			}
			esc = false
		} else if e == '\\' {
			esc = true
		} else if e == '"' {
			return i - cnt
		}
	}
	return -1
}

func (r *Reader) skipString() error {
	pos, dat := r.pos+1, r.dat
	esc, unc, dgt := false, false, 0

	for {
		if pos >= len(dat) {
			return NewEndOfFileError()
		} else if c := dat[pos]; unc {
			if !(c-'0' < 10) && !(c-'7' < 16) && !(c-'W' < 16) {
				NewInvalidCharacterError(c, pos)
			} else if dgt == 3 {
				esc, unc, dgt = false, false, 0
			} else {
				dgt++
			}
		} else if esc {
			switch c {
			case '"', '\\', '/', 'b',
				'f', 'n', 'r', 't':
				esc = false
			case 'u':
				unc = true
			default:
				NewInvalidCharacterError(c, pos)
			}
		} else if c == '\\' {
			esc = !esc
		} else if c == '"' {
			r.pos = pos + 1
			return nil
		}
		pos++
	}
}

func Utf8(dst, src []byte) (ln int) {
	if len(src) == 4 {
		n := uint(0)
		for i, e := range src {
			if e >= 'A' && e <= 'F' {
				n = n<<4 | uint(e-'7')
			} else if e >= 'a' && e <= 'f' {
				n = n<<4 | uint(e-'W')
			} else if e >= '0' && e <= '9' {
				n = n<<4 | uint(e-'0')
			} else {
				return -i
			}
		}

		if n < 0x80 {
			dst[0] = byte(n)
			return 1
		} else if n < 0x800 {
			dst[1] = 0x80 | (byte(n) & 0x3F)
			dst[0] = 0xC0 | (byte(n>>6) & 0x3F)
			return 2
		} else {
			dst[2] = 0x80 | (byte(n) & 0x3F)
			dst[1] = 0x80 | (byte(n>>6) & 0x3F)
			dst[0] = 0xE0 | (byte(n>>12) & 0x3F)
			return 3
		}
	} else {
		return 0
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
