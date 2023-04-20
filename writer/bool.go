package writer

// Gets the encoded byte length of a boolean.
func BoolLen(b bool) (bytes int) {
	if b {
		return 4
	} else {
		return 5
	}
}

// Gets the encoded byte length of a boolean pointer.
func BoolpLen(b *bool) (bytes int) {
	if b == nil {
		return 4
	} else if *b {
		return 4
	} else {
		return 5
	}
}

// Gets the encoded byte length of a boolean slice with brackets and commas.
func BoolsLen(bs []bool) (bytes int) {
	if bs == nil {
		return 4
	} else if len(bs) == 0 {
		return 2
	} else {
		bytes = 6*len(bs) + 1
		for _, b := range bs {
			if b {
				bytes--
			}
		}
		return
	}
}

// Writes "null" to the buffer when nil, otherwise writes a boolean.
func (w *Writer) Boolp(b *bool) {
	if b == nil {
		w.Raw("null")
	} else {
		w.Bool(*b)
	}
}

// Writes  "true" or "false" to the buffer.
func (w *Writer) Bool(b bool) {
	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	v, vln, cap := "false", 5, ln-cr

	if b {
		v, vln = "true", 4
	}

	if vln <= cap {
		w.Cursor += copy(bf[cr:], v)
	} else {
		copy(bf[cr:], v[:cap])
		w.Storage = append(w.Storage, bf)
		bf = make([]byte, vln-cap+CHUNK_SIZE)
		w.Cursor = copy(bf, v[cap:])
		w.Buffer = bf
	}
}

// Writes an array of true / false values separated by commas and enclosed by
// square brackets to the buffer. When the slice is nil, writes "null".
func (w *Writer) Bools(bs []bool) {
	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	v, vln, cap := "", 0, ln-cr

	if bs == nil {
		w.Raw("null")
		return
	} else if len(bs) == 0 {
		w.Raw("[]")
		return
	} else if 1+len(bs)*6 <= ln-cr {
		bf[cr] = '['
		for _, b := range bs {
			cr++
			if b {
				v = "true"
			} else {
				v = "false"
			}
			cr += copy(bf[cr:], v)
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

		for _, b := range bs {
			if b {
				v, vln = "true", 4
			} else {
				v, vln = "false", 5
			}

			if cap = ln - cr; vln <= cap {
				copy(bf[cr:], v)
				cr += vln
			} else {
				copy(bf[cr:], v[:cap])
				w.Storage = append(w.Storage, bf)
				ln = vln - cap + CHUNK_SIZE
				bf = make([]byte, ln)
				cr = copy(bf, v[cap:])
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
