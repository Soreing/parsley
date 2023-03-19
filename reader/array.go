package reader

func (r *Reader) skipArray() (err error) {
	if err = r.OpenArray(); err == nil {
		if r.GetType() != TerminatorToken {
			for err == nil {
				if err = r.Skip(); err != nil || !r.Next() {
					break
				}
			}
		}
		if err == nil {
			err = r.CloseArray()
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
	r.skipWhiteSpace()
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
	r.skipWhiteSpace()
	return nil
}
