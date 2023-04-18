package reader

type Reader struct {
	pos int
	dat []byte
	buf []byte
}

func NewReader(dat []byte) *Reader {
	return &Reader{
		pos: 0,
		dat: dat,
	}
}

func (r *Reader) Reset() {
	r.pos = 0
}

func (r *Reader) GetPosition() int {
	return r.pos
}

func (r *Reader) SetPosition(pos int) {
	r.pos = pos
}
