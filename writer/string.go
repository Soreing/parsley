package writer

var escape = []string{
	`\u0000`, `\u0001`, `\u0002`, `\u0003`, `\u0004`, `\u0005`, `\u0006`, `\u0007`,
	`\u0008`, `\t`, `\n`, `\u000B`, `\u000C`, `\r`, `\u000E`, `\u000F`,
	`\u0010`, `\u0011`, `\u0012`, `\u0013`, `\u0014`, `\u0015`, `\u0016`, `\u0017`,
	`\u0018`, `\u0019`, `\u001A`, `\u001B`, `\u001C`, `\u001D`, `\u001E`, `\u001F`,
	"", "", `\"`, "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", `\\`, "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
}

//go:noinline
func (w *Writer) swrite(bfp *[]byte, crp *int, lnp *int, v string) {
	vln, bf, cr, ln := len(v), *bfp, *crp, *lnp
	if vln == 0 {
		return
	} else if cap := ln - cr; vln <= cap {
		cr += copy(bf[cr:], v)
	} else {
		copy(bf[cr:], v[:cap])
		w.Storage = append(w.Storage, bf)
		ln = vln - cap + CHUNK_SIZE
		bf = make([]byte, ln)
		cr = copy(bf, v[cap:])
	}
	*bfp, *crp, *lnp = bf, cr, ln
}

// Gets the encoded byte length of a string with quotes.
func StringLen(str string) (bytes int, volatile int) {
	return len(str) + 2, len(str)
}

// Gets the encoded byte length of a string pointer with quotes.
func StringpLen(str *string) (bytes int, volatile int) {
	if str == nil {
		return 4, 0
	} else {
		return len(*str) + 2, len(*str)
	}
}

// Gets the encoded byte length of a string slice with brackets, commas and quotes.
func StringsLen(strs []string) (bytes int, volatile int) {
	if strs == nil {
		return 4, 0
	} else if len(strs) == 0 {
		return 2, 0
	} else {
		bytes++
		for _, s := range strs {
			bytes += len(s) + 3
			volatile += len(s)
		}
		return
	}
}

// Writes a sequence of bytes to the buffer without escaping.
func (w *Writer) Raw(s string) {
	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln, cap := len(s), ln-cr

	if vln <= cap {
		w.Cursor += copy(bf[cr:], s)
	} else {
		copy(bf[cr:], s[:cap])
		w.Storage = append(w.Storage, bf)
		bf = make([]byte, vln-cap+CHUNK_SIZE)
		w.Cursor = copy(bf, s[cap:])
		w.Buffer = bf
	}
}

// Writes a single byte to the buffer without escaping.
func (w *Writer) Byte(b byte) {
	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)

	if ln != cr {
		bf[cr] = b
		w.Cursor++
	} else {
		w.Storage = append(w.Storage, bf)
		bf = make([]byte, CHUNK_SIZE)
		bf[0] = b
		w.Cursor, w.Buffer = 1, bf
	}
}

// Writes "null" to the buffer when nil, otherwise writes a string.
func (w *Writer) Stringp(s *string) {
	if s != nil {
		w.String(*s)
	} else {
		w.Raw("null")
	}
}

// Writes a string to the buffer with quotes. Special characters are escaped.
func (w *Writer) String(s string) {
	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	beg, vln, cap, esc := 0, 0, 0, ""

	if ln != cr {
		bf[cr] = '"'
		cr++
	} else {
		w.Storage = append(w.Storage, bf)
		cr, ln = 1, CHUNK_SIZE
		bf = make([]byte, CHUNK_SIZE)
		bf[0] = '"'
	}

	for i, c := range s {
		if esc = escape[c&0xFF]; esc != "" && c < 128 {
			w.swrite(&bf, &cr, &ln, s[beg:i])
			w.swrite(&bf, &cr, &ln, esc)
			beg = i + 1
		}
	}

	if beg != len(s) {
		vln, cap = len(s)-beg, ln-cr
		if vln <= cap {
			cr += copy(bf[cr:], s[beg:])
		} else {
			copy(bf[cr:], s[beg:beg+cap])
			w.Storage = append(w.Storage, bf)
			ln = vln - cap + CHUNK_SIZE
			bf = make([]byte, ln)
			cr = copy(bf, s[beg+cap:])
		}
	}

	if ln != cr {
		bf[cr] = '"'
		cr++
	} else {
		w.Storage = append(w.Storage, bf)
		cr, ln = 1, CHUNK_SIZE
		bf = make([]byte, CHUNK_SIZE)
		bf[0] = '"'
	}
	w.Cursor, w.Buffer = cr, bf
}

// Writes an array of quoted strings separated by commas and enclosed by
// square brackets to the buffer. When the slice is nil, writes "null".
func (w *Writer) Strings(ss []string) {
	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	cap := ln - cr

	if ss == nil {
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
	} else if len(ss) == 0 {
		if 2 <= cap {
			bf[cr], bf[cr+1] = '[', ']'
			w.Cursor += 2
		} else if cap == 1 {
			bf[cr] = '['
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, CHUNK_SIZE)
			bf[0] = ']'
			w.Cursor, w.Buffer = 1, bf
		} else {
			w.Storage = append(w.Storage, bf)
			bf = make([]byte, CHUNK_SIZE)
			bf[0], bf[1] = '[', ']'
			w.Cursor, w.Buffer = 2, bf
		}
		return
	} else {
		if 2 <= cap {
			bf[cr], bf[cr+1] = '[', '"'
			cr += 2
		} else if cap == 1 {
			bf[cr] = '['
			w.Storage = append(w.Storage, bf)
			cr, ln = 1, CHUNK_SIZE
			bf = make([]byte, CHUNK_SIZE)
			bf[0] = '"'
		} else {
			w.Storage = append(w.Storage, bf)
			cr, ln = 2, CHUNK_SIZE
			bf = make([]byte, CHUNK_SIZE)
			bf[0], bf[1] = '[', '"'
		}

		beg, vln, esc, s := 0, 0, "", ss[0]
		for i, c := range s {
			if esc = escape[c&0xFF]; esc != "" {
				w.swrite(&bf, &cr, &ln, s[beg:i])
				w.swrite(&bf, &cr, &ln, esc)
				beg = i + 1
			}
		}
		if beg != len(s) {
			w.swrite(&bf, &cr, &ln, s[beg:])
		}

		for _, s := range ss[1:] {
			if cap = ln - cr; 3 <= cap {
				cr += copy(bf[cr:], "\",\"")
			} else {
				copy(bf[cr:], "\",\""[:cap])
				w.Storage = append(w.Storage, bf)
				ln = 3 - cap + CHUNK_SIZE
				bf = make([]byte, ln)
				cr = copy(bf, "\",\""[cap:])
			}

			beg = 0
			for i, c := range s {
				if esc = escape[c&0xFF]; esc != "" {
					w.swrite(&bf, &cr, &ln, s[beg:i])
					w.swrite(&bf, &cr, &ln, esc)
					beg = i + 1
				}
			}

			if beg != len(s) {
				vln, cap = len(s)-beg, ln-cr
				if vln <= cap {
					cr += copy(bf[cr:], s[beg:])
				} else {
					copy(bf[cr:], s[beg:beg+cap])
					w.Storage = append(w.Storage, bf)
					ln = vln - cap + CHUNK_SIZE
					bf = make([]byte, ln)
					cr = copy(bf, s[beg+cap:])
				}
			}
		}

		if cap = ln - cr; 2 <= cap {
			bf[cr], bf[cr+1] = '"', ']'
			cr += 2
		} else if cap == 1 {
			bf[cr] = '"'
			w.Storage = append(w.Storage, bf)
			cr, ln = 1, CHUNK_SIZE
			bf = make([]byte, CHUNK_SIZE)
			bf[0] = ']'
		} else {
			w.Storage = append(w.Storage, bf)
			cr, ln = 2, CHUNK_SIZE
			bf = make([]byte, CHUNK_SIZE)
			bf[0], bf[1] = '"', ']'
		}
		w.Cursor, w.Buffer = cr, bf
	}
}
