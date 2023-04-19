package reader

// skipObject skips an entire object enclosed by curly braces "{...}" and the
// whitespace after the object.
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
				return newEndOfFileError()
			} else if dat[pos] != '}' {
				return newInvalidCharacterError(dat[pos], pos)
			} else {
				r.pos++
			}
		}
	}
	return
}

// OpenObject consumes an opening curly brace '{' character and skips all
// whitespaces after it.
func (r *Reader) OpenObject() error {
	if r.pos >= len(r.dat) {
		return newEndOfFileError()
	}
	if r.dat[r.pos] != '{' {
		return newInvalidCharacterError(r.dat[r.pos], r.pos)
	}

	r.pos++
	r.SkipWhiteSpace()
	return nil
}

// CloseObject consumes a closing curly brace '}' character and skips all
// whitespaces after it.
func (r *Reader) CloseObject() error {
	if r.pos >= len(r.dat) {
		return newEndOfFileError()
	} else if r.dat[r.pos] != '}' {
		return newInvalidCharacterError(r.dat[r.pos], r.pos)
	}

	r.pos++
	r.SkipWhiteSpace()
	return nil
}

// Key extracts the next string and checks if it is followed by a colon ':'
// character. It returns the key if successful.
func (r *Reader) Key() ([]byte, error) {
	if key, err := r.Bytes(); err != nil {
		return nil, err
	} else {
		if r.pos == len(r.dat) {
			return nil, newEndOfFileError()
		} else if r.dat[r.pos] != ':' {
			return nil, newInvalidCharacterError(r.dat[r.pos], r.pos)
		}

		r.pos++
		r.SkipWhiteSpace()
		return key, nil
	}
}
