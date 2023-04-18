package reader

func (r *Reader) skipArray() (err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() != TerminatorToken {
			for err == nil {
				if err = r.Skip(); err != nil || !r.Next() {
					break
				}
			}
		}
		if err == nil {
			dat, pos := r.dat, r.pos
			if pos == len(dat) {
				return NewEndOfFileError()
			} else if dat[pos] != ']' {
				return NewInvalidCharacterError(dat[pos], pos)
			} else {
				r.pos++
			}
		}
	}
	return
}

func (r *Reader) OpenArray() error {
	if r.pos >= len(r.dat) {
		return NewEndOfFileError()
	}
	if r.dat[r.pos] != '[' {
		return NewInvalidCharacterError(r.dat[r.pos], r.pos)
	}

	r.pos++
	r.SkipWhiteSpace()
	return nil
}

func (r *Reader) CloseArray() error {
	if r.pos >= len(r.dat) {
		return NewEndOfFileError()
	}
	if r.dat[r.pos] != ']' {
		return NewInvalidCharacterError(r.dat[r.pos], r.pos)
	}

	r.pos++
	r.SkipWhiteSpace()
	return nil
}
