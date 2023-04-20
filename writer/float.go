package writer

import (
	"strconv"
)

func Float32Len(n float32) (bytes int) {
	return 24
}

func Float32pLen(n *float32) (bytes int) {
	if n == nil {
		return 4
	} else {
		return 24
	}
}

func Float32sLen(ns []float32) (bytes int) {
	if ns == nil {
		return 4
	} else if len(ns) == 0 {
		return 2
	} else {
		bytes++
		for range ns {
			bytes += 24 + 1
		}
		return
	}
}

// Writes "null" to the buffer when nil, otherwise writes an 32 bit float.
func (w *Writer) Float32p(n *float32) {
	if n == nil {
		w.Raw("null")
	} else {
		w.Float32(*n)
	}
}

// Writes an 32 bit float to the buffer.
func (w *Writer) Float32(n float32) {
	var vln int
	var dst []byte

	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln = Float32Len(n)
	if vln <= ln-cr {
		dst = bf[cr:]
	} else {
		w.Storage = append(w.Storage, bf[:cr])
		dst = make([]byte, vln+CHUNK_SIZE)
		w.Cursor, w.Buffer = 0, dst
	}
	w.Cursor += len(strconv.AppendFloat(dst[:0], float64(n), 'g', -1, 32))
}

// Writes an array of 32 bit float values separated by commas and enclosed
// by square brackets to the buffer. When the slice is nil, writes "null".
func (w *Writer) Float32s(ns []float32) {
	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln, cap := 0, ln-cr

	if ns == nil {
		w.Raw("null")
		return
	} else if len(ns) == 0 {
		w.Raw("[]")
		return
	} else if 1+len(ns)*25 <= ln-cr {
		bf[cr] = '['
		for _, n := range ns {
			cr++
			tb := strconv.AppendFloat(bf[:cr], float64(n), 'g', -1, 32)
			cr = len(tb)
			bf[cr] = ','
		}

		bf[cr] = ']'
		w.Cursor = cr + 1
		return
	} else {
		if ln != cr {
			bf[cr] = '['
			cr++
		} else {
			w.Storage = append(w.Storage, bf)
			cr, ln = 1, CHUNK_SIZE
			bf = make([]byte, CHUNK_SIZE)
			bf[0] = '['
		}

		for _, n := range ns {
			vln = Float32Len(n)
			if cap = ln - cr; vln <= cap {
				strconv.AppendFloat(bf[:cr], float64(n), 'g', -1, 32)
				cr += vln
			} else {
				w.Storage = append(w.Storage, bf[:cr])
				cr, ln = vln-cap, vln-cap+CHUNK_SIZE
				bf = make([]byte, ln)
				strconv.AppendFloat(bf, float64(n), 'g', -1, 32)
			}

			if ln != cr {
				bf[cr] = ','
				cr++
			} else {
				w.Storage = append(w.Storage, bf)
				cr, ln = 1, CHUNK_SIZE
				bf = make([]byte, CHUNK_SIZE)
				bf[0] = ','
			}
		}

		bf[cr-1] = ']'
		w.Cursor, w.Buffer = cr, bf
		return
	}
}

func Float64Len(n float64) (bytes int) {
	return 24
}

func Float64pLen(n *float64) (bytes int) {
	if n == nil {
		return 4
	} else {
		return 24
	}
}

func Float64sLen(ns []float64) (bytes int) {
	if ns == nil {
		return 4
	} else if len(ns) == 0 {
		return 2
	} else {
		bytes++
		for range ns {
			bytes += 24 + 1
		}
		return
	}
}

// Writes "null" to the buffer when nil, otherwise writes an 64 bit float.
func (w *Writer) Float64p(n *float64) {
	if n == nil {
		w.Raw("null")
	} else {
		w.Float64(*n)
	}
}

// Writes an 64 bit float to the buffer.
func (w *Writer) Float64(n float64) {
	var vln int
	var dst []byte

	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln = Float64Len(n)
	if vln <= ln-cr {
		dst = bf[cr:]
	} else {
		w.Storage = append(w.Storage, bf[:cr])
		dst = make([]byte, vln+CHUNK_SIZE)
		w.Cursor, w.Buffer = 0, dst
	}
	w.Cursor += len(strconv.AppendFloat(dst[:0], float64(n), 'g', -1, 64))
}

// Writes an array of 64 bit float values separated by commas and enclosed
// by square brackets to the buffer. When the slice is nil, writes "null".
func (w *Writer) Float64s(ns []float64) {
	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln, cap := 0, ln-cr

	if ns == nil {
		w.Raw("null")
		return
	} else if len(ns) == 0 {
		w.Raw("[]")
		return
	} else if 1+len(ns)*25 <= ln-cr {
		bf[cr] = '['
		for _, n := range ns {
			cr++
			tb := strconv.AppendFloat(bf[:cr], float64(n), 'g', -1, 64)
			cr = len(tb)
			bf[cr] = ','
		}

		bf[cr] = ']'
		w.Cursor = cr + 1
		return
	} else {
		if ln != cr {
			bf[cr] = '['
			cr++
		} else {
			w.Storage = append(w.Storage, bf)
			cr, ln = 1, CHUNK_SIZE
			bf = make([]byte, CHUNK_SIZE)
			bf[0] = '['
		}

		for _, n := range ns {
			vln = Float64Len(n)
			if cap = ln - cr; vln <= cap {
				strconv.AppendFloat(bf[:cr], float64(n), 'g', -1, 64)
				cr += vln
			} else {
				w.Storage = append(w.Storage, bf[:cr])
				cr, ln = vln-cap, vln-cap+CHUNK_SIZE
				bf = make([]byte, ln)
				strconv.AppendFloat(bf, float64(n), 'g', -1, 64)
			}

			if ln != cr {
				bf[cr] = ','
				cr++
			} else {
				w.Storage = append(w.Storage, bf)
				cr, ln = 1, CHUNK_SIZE
				bf = make([]byte, CHUNK_SIZE)
				bf[0] = ','
			}
		}

		bf[cr-1] = ']'
		w.Cursor, w.Buffer = cr, bf
		return
	}
}
