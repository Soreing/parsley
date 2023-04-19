package reader

type Reader struct {
	pos int
	dat []byte
	buf []byte
}

// NewReader creates a new Reader obect with some JSON data.
func NewReader(dat []byte) *Reader {
	return &Reader{
		pos: 0,
		dat: dat,
	}
}

// Reset sets the cursor position of the reader to the beginning of the data.
func (r *Reader) Reset() {
	r.pos = 0
}

// GetPosition returns the current cursor position in the reader.
func (r *Reader) GetPosition() int {
	return r.pos
}

// SetPosition sets the cursor position int the reader.
func (r *Reader) SetPosition(pos int) {
	r.pos = pos
}
