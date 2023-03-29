package reader

func (r *Reader) GetBool() (bool, error) {
	if r.pos < len(r.dat) {
		switch r.dat[r.pos] {
		case 't':
			if r.pos+3 >= len(r.dat) {
				return false, NewEndOfFileError()
			} else if r.dat[r.pos+1] != 'r' {
				return false, NewInvalidCharacterError(r.dat[r.pos+1], r.pos+1)
			} else if r.dat[r.pos+2] != 'u' {
				return false, NewInvalidCharacterError(r.dat[r.pos+2], r.pos+2)
			} else if r.dat[r.pos+3] != 'e' {
				return false, NewInvalidCharacterError(r.dat[r.pos+3], r.pos+3)
			} else {
				r.pos += 4
				r.skipWhiteSpace()
				return true, nil
			}
		case 'f':
			if r.pos+4 >= len(r.dat) {
				return false, NewEndOfFileError()
			} else if r.dat[r.pos+1] != 'a' {
				return false, NewInvalidCharacterError(r.dat[r.pos+1], r.pos+1)
			} else if r.dat[r.pos+2] != 'l' {
				return false, NewInvalidCharacterError(r.dat[r.pos+2], r.pos+2)
			} else if r.dat[r.pos+3] != 's' {
				return false, NewInvalidCharacterError(r.dat[r.pos+3], r.pos+3)
			} else if r.dat[r.pos+4] != 'e' {
				return false, NewInvalidCharacterError(r.dat[r.pos+4], r.pos+4)
			} else {
				r.pos += 5
				r.skipWhiteSpace()
				return false, nil
			}
		default:
			return false, NewInvalidCharacterError(r.dat[r.pos], r.pos)
		}
	} else {
		return false, NewEndOfFileError()
	}
}

func (r *Reader) boolSeq(idx int) (res []bool, err error) {
	var tf bool
	if tf, err = r.GetBool(); err == nil {
		if r.Next() {
			res, err = r.boolSeq(idx + 1)
		} else {
			res = make([]bool, idx+1)
		}

		if err == nil {
			res[idx] = tf
		}
	}
	return
}

func (r *Reader) GetBools() (res []bool, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.boolSeq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetBoolPtr() (res *bool, err error) {
	if v, err := r.GetBool(); err == nil {
		res = &v
	}
	return
}
