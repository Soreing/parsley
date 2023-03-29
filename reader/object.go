package reader

func (r *Reader) skipObject() (err error) {
	if err = r.OpenObject(); err == nil {
		if r.GetType() != TerminatorToken {
			for err == nil {
				if _, err = r.GetKey(); err == nil {
					if err = r.Skip(); err != nil || !r.Next() {
						break
					}
				}
			}
		}
		if err == nil {
			err = r.CloseObject()
		}
	}
	return
}

func (r *Reader) OpenObject() error {
	if r.pos >= len(r.dat) {
		return NewEndOfFileError()
	}
	if r.dat[r.pos] != '{' {
		return NewInvalidCharacterError(r.dat[r.pos], r.pos)
	}

	r.pos++
	r.SkipWhiteSpace()
	return nil
}

func (r *Reader) CloseObject() error {
	if r.pos >= len(r.dat) {
		return NewEndOfFileError()
	} else if r.dat[r.pos] != '}' {
		return NewInvalidCharacterError(r.dat[r.pos], r.pos)
	}

	r.pos++
	r.SkipWhiteSpace()
	return nil
}

func (r *Reader) GetKey() ([]byte, error) {
	if key, err := r.GetByteArray(); err != nil {
		return nil, err
	} else {
		if r.dat[r.pos] != ':' {
			return nil, NewInvalidCharacterError(r.dat[r.pos], r.pos)
		}

		r.pos++
		r.SkipWhiteSpace()
		return key, nil
	}
}
