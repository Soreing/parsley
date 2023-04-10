package writer

const CHUNK_SIZE = 256

type Writer struct {
	Cursor  int
	Buffer  []byte
	Storage [][]byte
}

func NewWriter(length int) *Writer {
	return &Writer{
		Cursor:  0,
		Buffer:  make([]byte, length),
		Storage: [][]byte{},
	}
}

func (w *Writer) Build() (res []byte) {
	ln, beg := w.Cursor, 0
	bf, st := w.Buffer[:ln], w.Storage

	if len(st) == 0 {
		return bf
	}

	for _, e := range st {
		ln += len(e)
	}

	res = make([]byte, ln)
	for _, e := range st {
		copy(res[beg:], e)
		beg += len(e)
	}
	copy(res[beg:], bf)
	return res
}
