package reader

func (r *Reader) skipObject() (err error) {
	if err = r.OpenObject(); err == nil {
		if r.Token() != TerminatorToken {
			for err == nil {
				if _, err = r.Key(); err == nil {
					if err = r.Skip(); err != nil || !r.Next() {
						break
					}
				}
			}
		}
		if err == nil {
			dat, pos := r.dat, r.pos
			if pos == len(dat) {
				return NewEndOfFileError()
			} else if dat[pos] != '}' {
				return NewInvalidCharacterError(dat[pos], pos)
			} else {
				r.pos++
			}
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

func (r *Reader) Key() ([]byte, error) {
	if key, err := r.Bytes(); err != nil {
		return nil, err
	} else {
		if r.pos == len(r.dat) {
			return nil, NewEndOfFileError()
		} else if r.dat[r.pos] != ':' {
			return nil, NewInvalidCharacterError(r.dat[r.pos], r.pos)
		}

		r.pos++
		r.SkipWhiteSpace()
		return key, nil
	}
}
