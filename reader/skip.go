package reader

func (r *Reader) skipWhiteSpace() {
	for {
		if r.pos >= len(r.dat) {
			return
		} else {
			switch r.dat[r.pos] {
			case '\t', '\n', '\r', ' ':
				r.pos++
			default:
				return
			}
		}
	}
}

func (r *Reader) Skip() error {
	switch r.dat[r.pos] {
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
		if r.pos+3 >= len(r.dat) {
			return NewEndOfFileError()
		} else if r.dat[r.pos+1] != 'r' {
			return NewInvalidCharacterError(r.dat[r.pos+1], r.pos+1)
		} else if r.dat[r.pos+2] != 'u' {
			return NewInvalidCharacterError(r.dat[r.pos+2], r.pos+2)
		} else if r.dat[r.pos+3] != 'e' {
			return NewInvalidCharacterError(r.dat[r.pos+3], r.pos+3)
		} else {
			r.pos += 4
		}
	case 'f':
		if r.pos+4 >= len(r.dat) {
			return NewEndOfFileError()
		} else if r.dat[r.pos+1] != 'a' {
			return NewInvalidCharacterError(r.dat[r.pos+1], r.pos+1)
		} else if r.dat[r.pos+2] != 'l' {
			return NewInvalidCharacterError(r.dat[r.pos+2], r.pos+2)
		} else if r.dat[r.pos+3] != 's' {
			return NewInvalidCharacterError(r.dat[r.pos+3], r.pos+3)
		} else if r.dat[r.pos+4] != 'e' {
			return NewInvalidCharacterError(r.dat[r.pos+4], r.pos+4)
		} else {
			r.pos += 5
		}
	case 'n':
		if r.pos+3 >= len(r.dat) {
			return NewEndOfFileError()
		} else if r.dat[r.pos+1] != 'u' {
			return NewInvalidCharacterError(r.dat[r.pos+1], r.pos+1)
		} else if r.dat[r.pos+2] != 'l' {
			return NewInvalidCharacterError(r.dat[r.pos+2], r.pos+2)
		} else if r.dat[r.pos+3] != 'l' {
			return NewInvalidCharacterError(r.dat[r.pos+3], r.pos+3)
		} else {
			r.pos += 4
		}
	default:
		return NewInvalidCharacterError(r.dat[r.pos], r.pos)
	}

	r.skipWhiteSpace()
	return nil
}

func (r *Reader) SkipNull() {
	r.pos += 4
	r.skipWhiteSpace()
}
