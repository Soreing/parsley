package writer

import (
	"encoding/base64"
	"strconv"
)

func UInt8Len(n uint8) (bytes int) {
	return ui8dc(uint8(n))
}

func UInt8pLen(n *uint8) (bytes int) {
	if n == nil {
		return 4
	} else {
		return ui8dc(uint8(*n))
	}
}

func UInt8sLen(ns []uint8) (bytes int) {
	if ns == nil {
		return 4
	} else {
		return (len(ns)+2)/3*4 + 2
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
	vln := (len(ns)+2)/3*4 + 2

	if ns == nil {
		w.Raw("null")
		return
	} else if len(ns) == 0 {
		w.Raw("\"\"")
		return
	} else if vln <= ln-cr {
		base64.StdEncoding.Encode(bf[cr+1:], ns)
		bf[cr], bf[cr+vln-1] = '"', '"'
		w.Cursor += vln
		return
	} else if vln-1 == ln-cr {
		bf[cr] = '"'
		base64.StdEncoding.Encode(bf[cr+1:], ns)
		w.Storage = append(w.Storage, bf)
		bf = make([]byte, CHUNK_SIZE)
		bf[0] = '"'
		w.Buffer, w.Cursor = bf, 1
		return
	} else if ln == cr {
		w.Storage = append(w.Storage, bf)
		bf = make([]byte, vln+CHUNK_SIZE)
		base64.StdEncoding.Encode(bf[1:], ns)
		bf[0], bf[vln-1] = '"', '"'
		w.Buffer, w.Cursor = bf, vln
		return
	} else if ln == cr+1 {
		bf[cr] = '"'
		w.Storage = append(w.Storage, bf)
		bf = make([]byte, vln+CHUNK_SIZE)
		base64.StdEncoding.Encode(bf, ns)
		bf[vln-2] = '"'
		w.Buffer, w.Cursor = bf, vln
		return
	} else {
		cap := (ln - cr) / 4
		bcap, dcap := cap*3, cap*4
		ovf := (len(ns) - bcap) * 4

		bf[cr] = '"'
		base64.StdEncoding.Encode(bf[cr+1:], ns[:bcap])
		w.Storage = append(w.Storage, bf[:cr+1+dcap])
		bf = make([]byte, ovf+CHUNK_SIZE)
		base64.StdEncoding.Encode(bf, ns[bcap:])
		bf[ovf] = '"'
		w.Buffer, w.Cursor = bf, ovf+1
		return
	}
}

func UInt16Len(n uint16) (bytes int) {
	return ui16dc(uint16(n))
}

func UInt16pLen(n *uint16) (bytes int) {
	if n == nil {
		return 4
	} else {
		return ui16dc(uint16(*n))
	}
}

func UInt16sLen(ns []uint16) (bytes int) {
	if ns == nil {
		return 4
	} else if len(ns) == 0 {
		return 2
	} else {
		bytes++
		for _, n := range ns {
			bytes += ui16dc(uint16(n)) + 1
		}
		return
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
		w.Raw("null")
		return
	} else if len(ns) == 0 {
		w.Raw("[]")
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
				cr, ln = vln, vln-cap+CHUNK_SIZE
				bf = make([]byte, ln)
				strconv.AppendInt(bf[:0], int64(n), 10)
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

func UInt32Len(n uint32) (bytes int) {
	return ui32dc(uint32(n))
}

func UInt32pLen(n *uint32) (bytes int) {
	if n == nil {
		return 4
	} else {
		return ui32dc(uint32(*n))
	}
}

func UInt32sLen(ns []uint32) (bytes int) {
	if ns == nil {
		return 4
	} else if len(ns) == 0 {
		return 2
	} else {
		bytes++
		for _, n := range ns {
			bytes += ui32dc(uint32(n)) + 1
		}
		return
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
		w.Raw("null")
		return
	} else if len(ns) == 0 {
		w.Raw("[]")
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
				cr, ln = vln, vln-cap+CHUNK_SIZE
				bf = make([]byte, ln)
				strconv.AppendInt(bf[:0], int64(n), 10)
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

func UInt64Len(n uint64) (bytes int) {
	return ui64dc(uint64(n))
}

func UInt64pLen(n *uint64) (bytes int) {
	if n == nil {
		return 4
	} else {
		return ui64dc(uint64(*n))
	}
}

func UInt64sLen(ns []uint64) (bytes int) {
	if ns == nil {
		return 4
	} else if len(ns) == 0 {
		return 2
	} else {
		bytes++
		for _, n := range ns {
			bytes += ui64dc(uint64(n)) + 1
		}
		return
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
		w.Raw("null")
		return
	} else if len(ns) == 0 {
		w.Raw("[]")
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
				cr, ln = vln, vln-cap+CHUNK_SIZE
				bf = make([]byte, ln)
				strconv.AppendInt(bf[:0], int64(n), 10)
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

func UIntLen(n uint) (bytes int) {
	return ui32dc(uint32(n))
}

func UIntpLen(n *uint) (bytes int) {
	if n == nil {
		return 4
	} else {
		return ui32dc(uint32(*n))
	}
}

func UIntsLen(ns []uint) (bytes int) {
	if ns == nil {
		return 4
	} else if len(ns) == 0 {
		return 2
	} else {
		bytes++
		for _, n := range ns {
			bytes += ui32dc(uint32(n)) + 1
		}
		return
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
		w.Raw("null")
		return
	} else if len(ns) == 0 {
		w.Raw("[]")
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
				cr, ln = vln, vln-cap+CHUNK_SIZE
				bf = make([]byte, ln)
				strconv.AppendInt(bf[:0], int64(n), 10)
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
