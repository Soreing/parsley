package writer

import (
	"time"
)

// Gets the encoded byte length of an RFC3339 time string with quotes.
func TimeLen(t time.Time) (bytes int) {
	nano, nanol := t.Nanosecond(), 10

	if nano > 0 {
		for nano%10 == 0 && nanol > 0 {
			nanol--
			nano /= 10
		}
	} else {
		nanol = 0
	}

	_, zone := t.Zone()
	if zone == 0 {
		return 22 + nanol
	} else {
		return 27 + nanol
	}
}

// Gets the encoded byte length of an RFC3339 time string pointer with quotes.
func TimepLen(t *time.Time) (bytes int) {
	if t == nil {
		return 4
	} else {
		return TimeLen(*t)
	}
}

// Gets the encoded byte length of an RFC3339 time string slice with
// brackets, commas and quotes.
func TimesLen(ts []time.Time) (bytes int) {
	if ts == nil {
		return 4
	} else if len(ts) == 0 {
		return 2
	} else {
		bytes++
		for _, t := range ts {
			bytes += TimeLen(t) + 1
		}
		return
	}
}

// Writes "null" to the buffer when nil, otherwise writes an RFC3339 time string.
func (w *Writer) Timep(t *time.Time) {
	if t != nil {
		w.Time(*t)
	} else {
		w.Raw("null")
	}
}

//go:noinline
func (w *Writer) twrite(bfp *[]byte, crp *int, lnp *int, t time.Time) {
	bf, cr, ln := *bfp, *crp, *lnp
	year, month, day := t.Date()
	hour, minute, second := t.Clock()
	nano, nanol := t.Nanosecond(), 10
	zonel, zsig := 1, byte('Z')
	dst := ([]byte)(nil)

	_, zone := t.Zone()
	if zone > 0 {
		zonel, zsig = 6, '+'
	} else if zone < 0 {
		zone, zonel, zsig = -zone, 6, '-'
	}

	if nano > 0 {
		for nano%10 == 0 && nanol > 0 {
			nanol--
			nano /= 10
		}
	} else {
		nanol = 0
	}

	vln := 19 + nanol + zonel
	if vln <= ln-cr {
		dst = bf[cr:]
		cr += vln
	} else {
		w.Storage = append(w.Storage, bf[:cr])
		dst = make([]byte, vln+CHUNK_SIZE)
		cr = vln
		bf = dst
	}

	// Year
	dst[3] = byte(year%10) + '0'
	year /= 10
	dst[2] = byte(year%10) + '0'
	year /= 10
	dst[1] = byte(year%10) + '0'
	year /= 10
	dst[0] = byte(year) + '0'
	// Month
	dst[6] = byte(month%10) + '0'
	month /= 10
	dst[5] = byte(month%10) + '0'
	// Day
	dst[9] = byte(day%10) + '0'
	day /= 10
	dst[8] = byte(day%10) + '0'
	// Hour
	dst[12] = byte(hour%10) + '0'
	hour /= 10
	dst[11] = byte(hour%10) + '0'
	// Minute
	dst[15] = byte(minute%10) + '0'
	minute /= 10
	dst[14] = byte(minute%10) + '0'
	// Second
	dst[18] = byte(second%10) + '0'
	second /= 10
	dst[17] = byte(second%10) + '0'

	// Nanoseconds
	if nano > 0 {
		dst[19] = '.'
		for i := nanol + 18; i >= 20; i-- {
			dst[i] = byte(nano%10) + '0'
			nano /= 10
		}
	}

	// Timezone
	if zone != 0 {
		hr, mn := zone/3600, zone%3600
		dst[nanol+24] = byte(mn%600) + '0'
		mn /= 600
		dst[nanol+23] = byte(mn) + '0'
		dst[nanol+22] = ':'
		dst[nanol+21] = byte(hr%10) + '0'
		hr /= 10
		dst[nanol+20] = byte(hr) + '0'
	}

	// Dressing
	dst[4], dst[7], dst[10], dst[13], dst[16] = '-', '-', 'T', ':', ':'
	dst[19+nanol] = zsig
	*bfp, *crp, *lnp = bf, cr, ln
}

// Writes an RFC3339 time string to the buffer with quotes.
func (w *Writer) Time(t time.Time) {
	year, month, day := t.Date()
	hour, minute, second := t.Clock()
	nano, nanol := t.Nanosecond(), 10
	zonel, zsig := 1, byte('Z')
	dst := ([]byte)(nil)

	_, zone := t.Zone()
	if zone > 0 {
		zonel, zsig = 6, '+'
	} else if zone < 0 {
		zone, zonel, zsig = -zone, 6, '-'
	}

	if nano > 0 {
		for nano%10 == 0 && nanol > 0 {
			nanol--
			nano /= 10
		}
	} else {
		nanol = 0
	}

	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	vln := 19 + nanol + zonel + 2
	if vln <= ln-cr {
		dst = bf[cr:]
		w.Cursor += vln
	} else {
		w.Storage = append(w.Storage, bf[:cr])
		dst = make([]byte, vln+CHUNK_SIZE)
		w.Cursor = vln
		w.Buffer = dst
	}

	// Year
	dst[4] = byte(year%10) + '0'
	year /= 10
	dst[3] = byte(year%10) + '0'
	year /= 10
	dst[2] = byte(year%10) + '0'
	year /= 10
	dst[1] = byte(year) + '0'
	// Month
	dst[7] = byte(month%10) + '0'
	month /= 10
	dst[6] = byte(month%10) + '0'
	// Day
	dst[10] = byte(day%10) + '0'
	day /= 10
	dst[9] = byte(day%10) + '0'
	// Hour
	dst[13] = byte(hour%10) + '0'
	hour /= 10
	dst[12] = byte(hour%10) + '0'
	// Minute
	dst[16] = byte(minute%10) + '0'
	minute /= 10
	dst[15] = byte(minute%10) + '0'
	// Second
	dst[19] = byte(second%10) + '0'
	second /= 10
	dst[18] = byte(second%10) + '0'

	// Nanoseconds
	if nano > 0 {
		dst[20] = '.'
		for i := nanol + 19; i >= 21; i-- {
			dst[i] = byte(nano%10) + '0'
			nano /= 10
		}
	}

	// Timezone
	if zone != 0 {
		hr, mn := zone/3600, zone%3600
		dst[nanol+25] = byte(mn%600) + '0'
		mn /= 600
		dst[nanol+24] = byte(mn) + '0'
		dst[nanol+23] = ':'
		dst[nanol+22] = byte(hr%10) + '0'
		hr /= 10
		dst[nanol+21] = byte(hr) + '0'
	}

	// Dressing
	dst[0], dst[20+nanol+zonel] = '"', '"'
	dst[5], dst[8], dst[11], dst[14], dst[17] = '-', '-', 'T', ':', ':'
	dst[20+nanol] = zsig
}

// Writes an array of RFC3339 time string values separated by commas and enclosed
// by square brackets to the buffer. When the slice is nil, writes "null".
func (w *Writer) Times(ts []time.Time) {
	bf := w.Buffer
	cr, ln := w.Cursor, len(bf)
	cap := ln - cr

	if ts == nil {
		w.Raw("null")
		return
	} else if len(ts) == 0 {
		w.Raw("[]")
		return
	} else if 1+len(ts)*38 <= ln-cr {
		bf[cr], bf[cr+1] = '[', '"'
		cr += 2
		w.twrite(&bf, &cr, &ln, ts[0])
		for _, t := range ts[1:] {
			cr += copy(bf[cr:], "\",\"")
			w.twrite(&bf, &cr, &ln, t)
		}

		bf[cr], bf[cr+1] = '"', ']'
		w.Cursor = cr + 2
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

		w.twrite(&bf, &cr, &ln, ts[0])
		for _, t := range ts[1:] {
			if cap = ln - cr; 3 <= cap {
				cr += copy(bf[cr:], "\",\"")
			} else {
				copy(bf[cr:], "\",\""[:cap])
				w.Storage = append(w.Storage, bf)
				ln = 3 - cap + CHUNK_SIZE
				bf = make([]byte, ln)
				cr = copy(bf, "\",\""[cap:])
			}
			w.twrite(&bf, &cr, &ln, t)
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
		return
	}
}
