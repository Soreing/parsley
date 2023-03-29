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

func (r *Reader) Next() bool {
	if r.pos < len(r.dat) && r.dat[r.pos] == ',' {
		r.pos++
		r.SkipWhiteSpace()
		return true
	}
	return false
}
