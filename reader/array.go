package reader

// skipArray skips an entire array enclosed by square brackets "[...]" and the
// whitespace after the array. The content of the array is evaluated to make
// sure that the JSON is valid.
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
				return newEndOfFileError()
			} else if dat[pos] != ']' {
				return newInvalidCharacterError(dat[pos], pos)
			} else {
				r.pos++
			}
		}
	}
	return
}

// OpenArray consumes an opening square bracket '[' character and skips all
// whitespaces after it.
func (r *Reader) OpenArray() error {
	if r.pos >= len(r.dat) {
		return newEndOfFileError()
	}
	if r.dat[r.pos] != '[' {
		return newInvalidCharacterError(r.dat[r.pos], r.pos)
	}

	r.pos++
	r.SkipWhiteSpace()
	return nil
}

// CloseArray consumes a closing square bracket ']' character and skips all
// whitespaces after it.
func (r *Reader) CloseArray() error {
	if r.pos >= len(r.dat) {
		return newEndOfFileError()
	}
	if r.dat[r.pos] != ']' {
		return newInvalidCharacterError(r.dat[r.pos], r.pos)
	}

	r.pos++
	r.SkipWhiteSpace()
	return nil
}
