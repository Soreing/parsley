package reader

func (r *Reader) Bool() (bool, error) {
	dat, pos := r.dat, r.pos
	if pos < len(dat) {
		switch dat[pos] {
		case 't':
			if pos+3 >= len(dat) {
				return false, NewEndOfFileError()
			} else if dat[pos+1] != 'r' {
				return false, NewInvalidCharacterError(dat[pos+1], pos+1)
			} else if dat[pos+2] != 'u' {
				return false, NewInvalidCharacterError(dat[pos+2], pos+2)
			} else if dat[pos+3] != 'e' {
				return false, NewInvalidCharacterError(dat[pos+3], pos+3)
			} else {
				r.pos += 4
				r.SkipWhiteSpace()
				return true, nil
			}
		case 'f':
			if pos+4 >= len(dat) {
				return false, NewEndOfFileError()
			} else if dat[pos+1] != 'a' {
				return false, NewInvalidCharacterError(dat[pos+1], pos+1)
			} else if dat[pos+2] != 'l' {
				return false, NewInvalidCharacterError(dat[pos+2], pos+2)
			} else if dat[pos+3] != 's' {
				return false, NewInvalidCharacterError(dat[pos+3], pos+3)
			} else if dat[pos+4] != 'e' {
				return false, NewInvalidCharacterError(dat[pos+4], pos+4)
			} else {
				r.pos += 5
				r.SkipWhiteSpace()
				return false, nil
			}
		default:
			return false, NewInvalidCharacterError(dat[pos], pos)
		}
	} else {
		return false, NewEndOfFileError()
	}
}

func (r *Reader) boolSeq(idx int) (res []bool, err error) {
	var tf bool
	if tf, err = r.Bool(); err == nil {
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

func (r *Reader) Bools() (res []bool, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []bool{}
			err = r.CloseArray()
		} else if res, err = r.boolSeq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) Boolp() (res *bool, err error) {
	if v, err := r.Bool(); err == nil {
		res = &v
	}
	return
}
