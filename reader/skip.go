package reader

func (r *Reader) SkipWhiteSpace() {
	dat, pos, c := r.dat, r.pos, byte(0)
	for pos < len(dat) {
		c = dat[pos]
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			pos++
		} else {
			break
		}
	}
	r.pos = pos
}

func (r *Reader) Next() bool {
	if r.pos < len(r.dat) && r.dat[r.pos] == ',' {
		r.pos++
		r.SkipWhiteSpace()
		return true
	}
	return false
}

func (r *Reader) Skip() error {
	dat, pos := r.dat, r.pos
	if pos == len(dat) {
		return NewEndOfFileError()
	}

	switch dat[pos] {
	case '{':
		if err := r.skipObject(); err != nil {
			return err
		}
	case '[':
		if err := r.skipArray(); err != nil {
			return err
		}
	case '"':
		if err := r.skipString(); err != nil {
			return err
		}
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		if err := r.skipNumber(); err != nil {
			return err
		}
	case 't':
		if pos+3 >= len(dat) {
			return NewEndOfFileError()
		} else if dat[pos+1] != 'r' {
			return NewInvalidCharacterError(dat[pos+1], pos+1)
		} else if dat[pos+2] != 'u' {
			return NewInvalidCharacterError(dat[pos+2], pos+2)
		} else if dat[pos+3] != 'e' {
			return NewInvalidCharacterError(dat[pos+3], pos+3)
		} else {
			r.pos += 4
		}
	case 'f':
		if pos+4 >= len(dat) {
			return NewEndOfFileError()
		} else if dat[pos+1] != 'a' {
			return NewInvalidCharacterError(dat[pos+1], pos+1)
		} else if dat[pos+2] != 'l' {
			return NewInvalidCharacterError(dat[pos+2], pos+2)
		} else if dat[pos+3] != 's' {
			return NewInvalidCharacterError(dat[pos+3], pos+3)
		} else if dat[pos+4] != 'e' {
			return NewInvalidCharacterError(dat[pos+4], pos+4)
		} else {
			r.pos += 5
		}
	case 'n':
		if pos+3 >= len(dat) {
			return NewEndOfFileError()
		} else if dat[pos+1] != 'u' {
			return NewInvalidCharacterError(dat[pos+1], pos+1)
		} else if dat[pos+2] != 'l' {
			return NewInvalidCharacterError(dat[pos+2], pos+2)
		} else if dat[pos+3] != 'l' {
			return NewInvalidCharacterError(dat[pos+3], pos+3)
		} else {
			r.pos += 4
		}
	default:
		return NewInvalidCharacterError(dat[pos], pos)
	}

	r.SkipWhiteSpace()
	return nil
}

func (r *Reader) SkipNull() {
	r.pos += 4
	r.SkipWhiteSpace()
}
