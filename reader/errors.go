package reader

import "fmt"

func NewEndOfFileError() error {
	return fmt.Errorf("unexpected end of file")
}

func NewInvalidCharacterError(b byte, pos int) error {
	return fmt.Errorf("invalid character '%c' at offset %d", b, pos)
}

func NewUnknownTimeFormatError(s string, pos int) error {
	return fmt.Errorf("unknown time format \"%s\" at offset %d", s, pos)
}
