package reader

// SkipWhiteSpace consumes any valid whitespace characters. Valid characters are
// spaces ' ', tabs '\t', new lines '\n' and carriage returns '\r'.
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

// Next checks if the next character is a comma ',' implying more data exists.
// If the character is a comma, it skips all whitespaces after it.
func (r *Reader) Next() bool {
	if r.pos < len(r.dat) && r.dat[r.pos] == ',' {
		r.pos++
		r.SkipWhiteSpace()
		return true
	}
	return false
}

// Skip examines the next token type and skip the value.
func (r *Reader) Skip() (err error) {
	dat, pos := r.dat, r.pos
	if pos == len(dat) {
		return newEndOfFileError()
	}

	if c := dat[pos]; c == '{' {
		err = r.skipObject()
	} else if c == '[' {
		err = r.skipArray()
	} else if c == '"' {
		err = r.skipString()
	} else if c-'0' <= 9 || c == '-' {
		err = r.skipNumber()
	} else if c == 't' {
		if pos+3 >= len(dat) {
			return newEndOfFileError()
		} else if dat[pos+1] != 'r' {
			return newInvalidCharacterError(dat[pos+1], pos+1)
		} else if dat[pos+2] != 'u' {
			return newInvalidCharacterError(dat[pos+2], pos+2)
		} else if dat[pos+3] != 'e' {
			return newInvalidCharacterError(dat[pos+3], pos+3)
		} else {
			r.pos += 4
		}
	} else if c == 'f' {
		if pos+4 >= len(dat) {
			return newEndOfFileError()
		} else if dat[pos+1] != 'a' {
			return newInvalidCharacterError(dat[pos+1], pos+1)
		} else if dat[pos+2] != 'l' {
			return newInvalidCharacterError(dat[pos+2], pos+2)
		} else if dat[pos+3] != 's' {
			return newInvalidCharacterError(dat[pos+3], pos+3)
		} else if dat[pos+4] != 'e' {
			return newInvalidCharacterError(dat[pos+4], pos+4)
		} else {
			r.pos += 5
		}
	} else if c == 'n' {
		if pos+3 >= len(dat) {
			return newEndOfFileError()
		} else if dat[pos+1] != 'u' {
			return newInvalidCharacterError(dat[pos+1], pos+1)
		} else if dat[pos+2] != 'l' {
			return newInvalidCharacterError(dat[pos+2], pos+2)
		} else if dat[pos+3] != 'l' {
			return newInvalidCharacterError(dat[pos+3], pos+3)
		} else {
			r.pos += 4
		}
	} else {
		return newInvalidCharacterError(dat[pos], pos)
	}

	if err == nil {
		r.SkipWhiteSpace()
	}
	return
}

// SkipNull moves the cursor forward by 4 places and skips all whitespaces after
// the token. Only user when it's been confirmed that the next value is null.
func (r *Reader) SkipNull() {
	r.pos += 4
	r.SkipWhiteSpace()
}
