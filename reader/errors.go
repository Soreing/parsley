package reader

import "fmt"

func NewEndOfFileError() error {
	return fmt.Errorf("unexpected end of file")
}

func NewInvalidCharacterError(b byte, pos int) error {
	return fmt.Errorf("invalid character '%c' at offset %d", b, pos)
}
