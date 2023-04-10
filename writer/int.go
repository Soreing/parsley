package writer

import (
	"strconv"
)

func Int8Len(n int8) (ln int) {
	if n < 0 {
		return ui8dc(uint8(-n)) + 1
	} else {
		return ui8dc(uint8(n))
	}
}

func Int8sLen(ns []int8) (ln int) {
	for _, n := range ns {
		ln += Int8Len(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

// Writes "null" to the buffer when nil, otherwise writes an 8 bit integer.
func (w *Writer) Int8p(n *int8) {
	if n == nil {
		w.Raw("null")
	} else {
		w.Int8(*n)
	}
}

// Writes an 8 bit integer to the buffer.
func (w *Writer) Int8(n int8) {
	var vln int
	var dst []byte

	if n < 0 {
		vln = ui8dc(uint8(-n)) + 1
	} else {
		vln = ui8dc(uint8(n))
	}

	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
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

// Writes an array of 8 bit integer values separated by commas and enclosed by
// square brackets to the buffer. When the slice is nil, writes "null".
func (w *Writer) Int8s(ns []int8) {
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
	} else if 1+len(ns)*5 <= ln-cr {
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
			vln = Int8Len(n)
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

func Int16Len(n int16) (ln int) {
	if n < 0 {
		return ui16dc(uint16(-n)) + 1
	} else {
		return ui16dc(uint16(n))
	}
}

func Int16sLen(ns []int16) (ln int) {
	for _, n := range ns {
		ln += Int16Len(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

// Writes "null" to the buffer when nil, otherwise writes an 16 bit integer.
func (w *Writer) Int16p(n *int16) {
	if n == nil {
		w.Raw("null")
	} else {
		w.Int16(*n)
	}
}

// Writes an 16 bit integer to the buffer.
func (w *Writer) Int16(n int16) {
	var vln int
	var dst []byte

	if n < 0 {
		vln = ui16dc(uint16(-n)) + 1
	} else {
		vln = ui16dc(uint16(n))
	}

	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
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

// Writes an array of 16 bit integer values separated by commas and enclosed by
// square brackets to the buffer. When the slice is nil, writes "null".
func (w *Writer) Int16s(ns []int16) {
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
	} else if 1+len(ns)*7 <= ln-cr {
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
			vln = Int16Len(n)
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

func Int32Len(n int32) (ln int) {
	if n < 0 {
		return ui32dc(uint32(-n)) + 1
	} else {
		return ui32dc(uint32(n))
	}
}

func Int32sLen(ns []int32) (ln int) {
	for _, n := range ns {
		ln += Int32Len(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

// Writes "null" to the buffer when nil, otherwise writes an 32 bit integer.
func (w *Writer) Int32p(n *int32) {
	if n == nil {
		w.Raw("null")
	} else {
		w.Int32(*n)
	}
}

// Writes an 32 bit integer to the buffer.
func (w *Writer) Int32(n int32) {
	var vln int
	var dst []byte

	if n < 0 {
		vln = ui32dc(uint32(-n)) + 1
	} else {
		vln = ui32dc(uint32(n))
	}

	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
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

// Writes an array of 32 bit integer values separated by commas and enclosed by
// square brackets to the buffer. When the slice is nil, writes "null".
func (w *Writer) Int32s(ns []int32) {
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
			vln = Int32Len(n)
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

func Int64Len(n int64) (ln int) {
	if n < 0 {
		return ui64dc(-uint64(n)) + 1
	} else {
		return ui64dc(-uint64(n))
	}
}

func Int64sLen(ns []int64) (ln int) {
	for _, n := range ns {
		ln += Int64Len(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

// Writes "null" to the buffer when nil, otherwise writes an 64 bit integer.
func (w *Writer) Int64p(n *int64) {
	if n == nil {
		w.Raw("null")
	} else {
		w.Int64(*n)
	}
}

// Writes an 64 bit integer to the buffer.
func (w *Writer) Int64(n int64) {
	var vln int
	var dst []byte

	if n < 0 {
		vln = ui64dc(uint64(-n)) + 1
	} else {
		vln = ui64dc(uint64(n))
	}

	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
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

// Writes an array of 64 bit integer values separated by commas and enclosed by
// square brackets to the buffer. When the slice is nil, writes "null".
func (w *Writer) Int64s(ns []int64) {
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
			vln = Int64Len(n)
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

func IntLen(n int) (ln int) {
	if n < 0 {
		return ui32dc(uint32(-n)) + 1
	} else {
		return ui32dc(uint32(n))
	}
}

func IntsLen(ns []int) (ln int) {
	for _, n := range ns {
		ln += IntLen(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

func WriteInt(dst []byte, n int) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[0] = '0'
		return 1
	}
}

func WriteIntPtr(dst []byte, n *int) (ln int) {
	if n != nil {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(*n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		return copy(dst, "null")
	}
}

func WriteInts(dst []byte, ns []int) (ln int) {
	tmp, res := make([]byte, 0, 32), ([]byte)(nil)
	if len(ns) > 0 {
		ln = 1
		for _, n := range ns {
			res = strconv.AppendInt(tmp, int64(n), 10)
			ln += copy(dst[ln:], res)
			dst[ln] = ','
			ln++
		}

		dst[0], dst[ln-1] = '[', ']'
		return ln
	} else if ns != nil {
		return copy(dst, "[]")
	} else {
		return copy(dst, "null")
	}
}

// Writes "null" to the buffer when nil, otherwise writes an integer.
func (w *Writer) Intp(n *int) {
	if n == nil {
		w.Raw("null")
	} else {
		w.Int(*n)
	}
}

// Writes an integer to the buffer.
func (w *Writer) Int(n int) {
	var vln int
	var dst []byte

	if n < 0 {
		vln = ui32dc(uint32(-n)) + 1
	} else {
		vln = ui32dc(uint32(n))
	}

	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
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

// Writes an array of integer values separated by commas and enclosed by
// square brackets to the buffer. When the slice is nil, writes "null".
func (w *Writer) Ints(ns []int) {
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
			vln = IntLen(n)
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
