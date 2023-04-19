package reader

// Bool extracts the next boolean "true" or "false" value from the data and
// skips all whitespace after it.
func (r *Reader) Bool() (bool, error) {
	dat, pos := r.dat, r.pos
	if pos < len(dat) {
		switch dat[pos] {
		case 't':
			if pos+3 >= len(dat) {
				return false, newEndOfFileError()
			} else if dat[pos+1] != 'r' {
				return false, newInvalidCharacterError(dat[pos+1], pos+1)
			} else if dat[pos+2] != 'u' {
				return false, newInvalidCharacterError(dat[pos+2], pos+2)
			} else if dat[pos+3] != 'e' {
				return false, newInvalidCharacterError(dat[pos+3], pos+3)
			} else {
				r.pos += 4
				r.SkipWhiteSpace()
				return true, nil
			}
		case 'f':
			if pos+4 >= len(dat) {
				return false, newEndOfFileError()
			} else if dat[pos+1] != 'a' {
				return false, newInvalidCharacterError(dat[pos+1], pos+1)
			} else if dat[pos+2] != 'l' {
				return false, newInvalidCharacterError(dat[pos+2], pos+2)
			} else if dat[pos+3] != 's' {
				return false, newInvalidCharacterError(dat[pos+3], pos+3)
			} else if dat[pos+4] != 'e' {
				return false, newInvalidCharacterError(dat[pos+4], pos+4)
			} else {
				r.pos += 5
				r.SkipWhiteSpace()
				return false, nil
			}
		default:
			return false, newInvalidCharacterError(dat[pos], pos)
		}
	} else {
		return false, newEndOfFileError()
	}
}

// boolSeq extracts booleans recursively untill the closing bracket is found,
// then assigns the elements to the allocated slice.
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

// Bools extracts an array of boolean values from the data and skips all
// whitespace after it. The values must be enclosed in square brackets "[...]"
// and the values must be separated by commas.
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

// Bool extracts the next boolean value and returns a pointer variable.
func (r *Reader) Boolp() (res *bool, err error) {
	if v, err := r.Bool(); err == nil {
		res = &v
	}
	return
}
