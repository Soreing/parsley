package reader

// TokenKind is an enum that represent the type of a JSON token
type TokenKind int

// TokenKind can be one of these values
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

// IsNull directly checks if the next token is a null
func (r *Reader) IsNull() bool {
	dat := r.dat[r.pos:]
	return len(dat) >= 4 &&
		dat[0] == 'n' && dat[1] == 'u' && dat[2] == 'l' && dat[3] == 'l'
}

// Token returns the TokenKind of the next token in the JSON.
func (r *Reader) Token() TokenKind {
	if r.pos >= len(r.dat) {
		return InvalidToken
	} else if c := r.dat[r.pos]; c == ',' || c == ':' {
		return SeparatorToken
	} else if c == ']' || c == '}' {
		return TerminatorToken
	} else if c == '{' {
		return ObjectToken
	} else if c == '[' {
		return ArrayToken
	} else if c == '"' {
		return StringToken
	} else if c-'0' <= 9 || c == '-' {
		return NumberToken
	} else if c == 't' || c == 'f' {
		return BooleanToken
	} else if c == 'n' {
		return NullToken
	} else {
		return InvalidToken
	}
}
