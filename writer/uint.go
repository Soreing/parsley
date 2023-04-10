package writer

import (
	"encoding/base64"
	"strconv"
)

func UInt8Len(n uint8) (ln int) {
	return ui8dc(uint8(n))
}

func UInt8sLen(ns []uint8) (ln int) {
	for _, n := range ns {
		ln += ui8dc(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

// Writes "null" to the buffer when nil, otherwise writes an 8 bit unsigned integer.
func (w *Writer) UInt8p(n *uint8) {
	if n == nil {
		w.Raw("null")
	} else {
		w.UInt8(*n)
	}
}

// Writes an 8 bit unsigned integer to the buffer.
func (w *Writer) UInt8(n uint8) {
	var vln int
	var dst []byte

	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln = ui8dc(uint8(n))
	if vln <= ln-cr {
		dst = bf[cr:]
		w.Cursor += vln
	} else {
		w.Storage = append(w.Storage, bf[:cr])
		dst = make([]byte, vln+CHUNK_SIZE)
		w.Cursor, w.Buffer = vln, dst
	}
	strconv.AppendInt(dst[:0], int64(n), 10)
}

func (w *Writer) UInt8s(ns []uint8) {
	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln, cap := (len(ns)+2)/3*4+2, ln-cr

	if ns == nil {
		if 4 <= cap {
			copy(bf[cr:], "null")
			w.Cursor += 4
		} else {
			copy(bf[cr:], "null"[:cap])
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, 4-cap+CHUNK_SIZE)
			w.Cursor = copy(bf, "null"[cap:])
			w.Buffer = bf
		}
		return
	} else if len(ns) == 0 {
		if 2 <= cap {
			bf[cr], bf[cr+1] = '"', '"'
			w.Cursor += 2
		} else if cap == 1 {
			bf[cr] = '"'
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, CHUNK_SIZE)
			bf[0] = '"'
			w.Cursor, w.Buffer = 1, bf
		} else {
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, CHUNK_SIZE)
			bf[0], bf[1] = '"', '"'
			w.Cursor, w.Buffer = 2, bf
		}
		return
	} else if vln <= ln-cr {
		dst := bf[cr+1:]
		base64.StdEncoding.Encode(dst, ns)
		bf[cr], bf[cr+vln-1] = '"', '"'
		w.Cursor += vln
		return
	} else {
		return
	}
}

func UInt16Len(n uint16) (ln int) {
	return ui16dc(uint16(n))
}

func UInt16sLen(ns []uint16) (ln int) {
	for _, n := range ns {
		ln += ui16dc(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

// Writes "null" to the buffer when nil, otherwise writes an 16 bit unsigned integer.
func (w *Writer) UInt16p(n *uint16) {
	if n == nil {
		w.Raw("null")
	} else {
		w.UInt16(*n)
	}
}

// Writes an 16 bit unsigned integer to the buffer.
func (w *Writer) UInt16(n uint16) {
	var vln int
	var dst []byte

	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln = ui16dc(uint16(n))
	if vln <= ln-cr {
		dst = bf[cr:]
		w.Cursor += vln
	} else {
		w.Storage = append(w.Storage, bf[:cr])
		dst = make([]byte, vln+CHUNK_SIZE)
		w.Cursor, w.Buffer = vln, dst
	}
	strconv.AppendInt(dst[:0], int64(n), 10)
}

// Writes an array of 16 bit unsigned integer values separated by commas and enclosed
// by square brackets to the buffer. When the slice is nil, writes "null".
func (w *Writer) UInt16s(ns []uint16) {
	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln, cap := 0, ln-cr

	if ns == nil {
		if 4 <= cap {
			copy(bf[cr:], "null")
			w.Cursor += 4
		} else {
			copy(bf[cr:], "null"[:cap])
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, 4-cap+CHUNK_SIZE)
			w.Cursor = copy(bf, "null"[cap:])
			w.Buffer = bf
		}
		return
	} else if len(ns) == 0 {
		if 2 <= cap {
			bf[cr], bf[cr+1] = '[', ']'
			w.Cursor += 2
		} else if cap == 1 {
			bf[cr] = '['
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, CHUNK_SIZE)
			w.Cursor, bf[0] = 1, ']'
			w.Buffer = bf
		} else {
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, CHUNK_SIZE)
			w.Cursor, bf[0], bf[1] = 2, '[', ']'
			w.Buffer = bf
		}
		return
	} else if 1+len(ns)*6 <= ln-cr {
		bf[cr] = '['
		for _, n := range ns {
			cr++
			tb := strconv.AppendInt(bf[:cr], int64(n), 10)
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
			vln = UInt16Len(n)
			if cap = ln - cr; vln <= cap {
				strconv.AppendInt(bf[:cr], int64(n), 10)
				cr += vln
			} else {
				w.Storage = append(w.Storage, bf[:cr])
				cr, ln = vln-cap, vln-cap+CHUNK_SIZE
				bf = make([]byte, ln)
				strconv.AppendInt(bf, int64(n), 10)
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

func UInt32Len(n uint32) (ln int) {
	return ui32dc(uint32(n))
}

func UInt32sLen(ns []uint32) (ln int) {
	for _, n := range ns {
		ln += ui32dc(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

// Writes "null" to the buffer when nil, otherwise writes an 32 bit unsigned integer.
func (w *Writer) UInt32p(n *uint32) {
	if n == nil {
		w.Raw("null")
	} else {
		w.UInt32(*n)
	}
}

// Writes an 32 bit unsigned integer to the buffer.
func (w *Writer) UInt32(n uint32) {
	var vln int
	var dst []byte

	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln = ui32dc(uint32(n))
	if vln <= ln-cr {
		dst = bf[cr:]
		w.Cursor += vln
	} else {
		w.Storage = append(w.Storage, bf[:cr])
		dst = make([]byte, vln+CHUNK_SIZE)
		w.Cursor, w.Buffer = vln, dst
	}
	strconv.AppendInt(dst[:0], int64(n), 10)
}

// Writes an array of 32 bit unsigned integer values separated by commas and enclosed
// by square brackets to the buffer. When the slice is nil, writes "null".
func (w *Writer) UInt32s(ns []uint32) {
	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln, cap := 0, ln-cr

	if ns == nil {
		if 4 <= cap {
			copy(bf[cr:], "null")
			w.Cursor += 4
		} else {
			copy(bf[cr:], "null"[:cap])
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, 4-cap+CHUNK_SIZE)
			w.Cursor = copy(bf, "null"[cap:])
			w.Buffer = bf
		}
		return
	} else if len(ns) == 0 {
		if 2 <= cap {
			bf[cr], bf[cr+1] = '[', ']'
			w.Cursor += 2
		} else if cap == 1 {
			bf[cr] = '['
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, CHUNK_SIZE)
			w.Cursor, bf[0] = 1, ']'
			w.Buffer = bf
		} else {
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, CHUNK_SIZE)
			w.Cursor, bf[0], bf[1] = 2, '[', ']'
			w.Buffer = bf
		}
		return
	} else if 1+len(ns)*4 <= ln-cr {
		bf[cr] = '['
		for _, n := range ns {
			cr++
			tb := strconv.AppendInt(bf[:cr], int64(n), 10)
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
			vln = UInt32Len(n)
			if cap = ln - cr; vln <= cap {
				strconv.AppendInt(bf[:cr], int64(n), 10)
				cr += vln
			} else {
				w.Storage = append(w.Storage, bf[:cr])
				cr, ln = vln-cap, vln-cap+CHUNK_SIZE
				bf = make([]byte, ln)
				strconv.AppendInt(bf, int64(n), 10)
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

func UInt64Len(n uint64) (ln int) {
	return ui64dc(uint64(n))
}

func UInt64sLen(ns []uint64) (ln int) {
	for _, n := range ns {
		ln += ui64dc(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

// Writes "null" to the buffer when nil, otherwise writes an 64 bit unsigned integer.
func (w *Writer) UInt64p(n *uint64) {
	if n == nil {
		w.Raw("null")
	} else {
		w.UInt64(*n)
	}
}

// Writes an 64 bit unsigned integer to the buffer.
func (w *Writer) UInt64(n uint64) {
	var vln int
	var dst []byte

	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln = ui64dc(uint64(n))
	if vln <= ln-cr {
		dst = bf[cr:]
		w.Cursor += vln
	} else {
		w.Storage = append(w.Storage, bf[:cr])
		dst = make([]byte, vln+CHUNK_SIZE)
		w.Cursor, w.Buffer = vln, dst
	}
	strconv.AppendInt(dst[:0], int64(n), 10)
}

// Writes an array of 64 bit unsigned integer values separated by commas and enclosed
// by square brackets to the buffer. When the slice is nil, writes "null".
func (w *Writer) UInt64s(ns []uint64) {
	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln, cap := 0, ln-cr

	if ns == nil {
		if 4 <= cap {
			copy(bf[cr:], "null")
			w.Cursor += 4
		} else {
			copy(bf[cr:], "null"[:cap])
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, 4-cap+CHUNK_SIZE)
			w.Cursor = copy(bf, "null"[cap:])
			w.Buffer = bf
		}
		return
	} else if len(ns) == 0 {
		if 2 <= cap {
			bf[cr], bf[cr+1] = '[', ']'
			w.Cursor += 2
		} else if cap == 1 {
			bf[cr] = '['
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, CHUNK_SIZE)
			w.Cursor, bf[0] = 1, ']'
			w.Buffer = bf
		} else {
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, CHUNK_SIZE)
			w.Cursor, bf[0], bf[1] = 2, '[', ']'
			w.Buffer = bf
		}
		return
	} else if 1+len(ns)*21 <= ln-cr {
		bf[cr] = '['
		for _, n := range ns {
			cr++
			tb := strconv.AppendInt(bf[:cr], int64(n), 10)
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
			vln = UInt64Len(n)
			if cap = ln - cr; vln <= cap {
				strconv.AppendInt(bf[:cr], int64(n), 10)
				cr += vln
			} else {
				w.Storage = append(w.Storage, bf[:cr])
				cr, ln = vln-cap, vln-cap+CHUNK_SIZE
				bf = make([]byte, ln)
				strconv.AppendInt(bf, int64(n), 10)
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

func UIntLen(n uint) (ln int) {
	return ui32dc(uint32(n))
}

func UIntsLen(ns []uint) (ln int) {
	for _, n := range ns {
		ln += ui32dc(uint32(n)) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

// Writes "null" to the buffer when nil, otherwise writes an unsigned integer.
func (w *Writer) UIntp(n *uint) {
	if n == nil {
		w.Raw("null")
	} else {
		w.UInt(*n)
	}
}

// Writes an unsigned integer to the buffer.
func (w *Writer) UInt(n uint) {
	var vln int
	var dst []byte

	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln = ui32dc(uint32(n))
	if vln <= ln-cr {
		dst = bf[cr:]
		w.Cursor += vln
	} else {
		w.Storage = append(w.Storage, bf[:cr])
		dst = make([]byte, vln+CHUNK_SIZE)
		w.Cursor, w.Buffer = vln, dst
	}
	strconv.AppendInt(dst[:0], int64(n), 10)
}

// Writes an array of unsigned integer values separated by commas and enclosed
// by square brackets to the buffer. When the slice is nil, writes "null".
func (w *Writer) UInts(ns []uint) {
	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln, cap := 0, ln-cr

	if ns == nil {
		if 4 <= cap {
			copy(bf[cr:], "null")
			w.Cursor += 4
		} else {
			copy(bf[cr:], "null"[:cap])
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, 4-cap+CHUNK_SIZE)
			w.Cursor = copy(bf, "null"[cap:])
			w.Buffer = bf
		}
		return
	} else if len(ns) == 0 {
		if 2 <= cap {
			bf[cr], bf[cr+1] = '[', ']'
			w.Cursor += 2
		} else if cap == 1 {
			bf[cr] = '['
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, CHUNK_SIZE)
			w.Cursor, bf[0] = 1, ']'
			w.Buffer = bf
		} else {
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, CHUNK_SIZE)
			w.Cursor, bf[0], bf[1] = 2, '[', ']'
			w.Buffer = bf
		}
		return
	} else if 1+len(ns)*12 <= ln-cr {
		bf[cr] = '['
		for _, n := range ns {
			cr++
			tb := strconv.AppendInt(bf[:cr], int64(n), 10)
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
			vln = UIntLen(n)
			if cap = ln - cr; vln <= cap {
				strconv.AppendInt(bf[:cr], int64(n), 10)
				cr += vln
			} else {
				w.Storage = append(w.Storage, bf[:cr])
				cr, ln = vln-cap, vln-cap+CHUNK_SIZE
				bf = make([]byte, ln)
				strconv.AppendInt(bf, int64(n), 10)
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
