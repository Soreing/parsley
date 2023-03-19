package reader

type TokenKind int

const (
	InvalidToken TokenKind = iota
	SeparatorToken
	TerminatorToken
	NullToken
	NumberToken
	BooleanToken
	StringToken
	ObjectToken
	ArrayToken
)

func (r *Reader) IsNull() bool {
	if r.pos+3 < len(r.dat) &&
		r.dat[r.pos+1] == 'u' &&
		r.dat[r.pos+2] == 'l' &&
		r.dat[r.pos+3] == 'l' {
		return true
	}
	return false
}

func (r *Reader) GetType() TokenKind {
	if r.pos < len(r.dat) {
		switch (r.dat)[r.pos] {
		case ',', ':':
			return SeparatorToken
		case '}', ']':
			return TerminatorToken
		case '{':
			return ObjectToken
		case '[':
			return ArrayToken
		case '"':
			return StringToken
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return NumberToken
		case 't', 'f':
			return BooleanToken
		case 'n':
			if r.pos+3 < len(r.dat) &&
				r.dat[r.pos+1] == 'u' &&
				r.dat[r.pos+2] == 'l' &&
				r.dat[r.pos+3] == 'l' {
				return NullToken
			}
		}
	}
	return InvalidToken
}
