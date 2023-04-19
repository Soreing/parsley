package reader

import "fmt"

func newEndOfFileError() error {
	return fmt.Errorf("unexpected end of file")
}

func newInvalidCharacterError(b byte, pos int) error {
	return fmt.Errorf("invalid character '%c' at offset %d", b, pos)
}

func newUnknownTimeFormatError(s string, pos int) error {
	return fmt.Errorf("unknown time format \"%s\" at offset %d", s, pos)
}

func newBase64PaddingError(pos int) error {
	return fmt.Errorf("invalid base64 padding at offset %d", pos)
}

func newNumberOutOfRangeError(num []byte, pos int) error {
	return fmt.Errorf("number %s out of range at offset %d", num, pos)
}
