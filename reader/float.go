package reader

import (
	"math"

	"github.com/Soreing/parsley/reader/floatconv"
)

// float32Seq extracts float32 values recursively until the closing bracket
// is found, then assigns the elements to the allocated slice.
func (r *Reader) float32Seq(idx int) (res []float32, err error) {
	var n float32
	if n, err = r.Float32(); err == nil {
		if r.Next() {
			res, err = r.float32Seq(idx + 1)
		} else {
			res = make([]float32, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

// Float32s extracts an array of float32 values from the data and skips all
// whitespace after it. The values must be enclosed in square brackets "[...]"
// and the values must be separated by commas.
func (r *Reader) Float32s() (res []float32, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []float32{}
			err = r.CloseArray()
		} else if res, err = r.float32Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

// Float32 extracts the next float32 value from the data and skips all
// whitespace after it.
func (r *Reader) Float32() (flt float32, err error) {
	dat, done := r.dat[r.pos:], false

	m, d, e, n, t, dp, sp, i, ok := readFloat(dat)
	if !ok {
		if i == len(dat) {
			return 0, newEndOfFileError()
		} else {
			return 0, newInvalidCharacterError(dat[i], r.pos+i)
		}
	}

	// adjusting exponent
	ae := e + dotExp(d, dp, sp, t)

	if d == 0 {
		flt, done = 0, true
	} else if !t {
		if f, ok := floatconv.Atof32exact(m, ae, n); ok {
			flt, done = f, true
		}
	}

	if !done {
		f, ok := floatconv.EiselLemire32(m, ae, n)
		if ok {
			if !t {
				flt, done = f, true
			}
			fu, ok := floatconv.EiselLemire32(m, ae, n)
			if ok && f == fu {
				flt, done = f, true
			}
		}
	}

	if !done {
		var dec floatconv.Decimal
		dec.Set(dat[sp:i], e, n, t, dp)

		b, ovf := dec.FloatBits(&floatconv.Float32info)
		flt = math.Float32frombits(uint32(b))
		if ovf {
			return flt, newNumberOutOfRangeError(dat[sp:i], r.pos)
		}
	}

	r.pos += i
	r.SkipWhiteSpace()
	return
}

// Float32p extracts the next float32 value and returns a pointer variable.
func (r *Reader) Float32p() (res *float32, err error) {
	if v, err := r.Float32(); err == nil {
		res = &v
	}
	return
}

// float64Seq extracts float64 values recursively until the closing bracket
// is found, then assigns the elements to the allocated slice.
func (r *Reader) float64Seq(idx int) (res []float64, err error) {
	var n float64
	if n, err = r.Float64(); err == nil {
		if r.Next() {
			res, err = r.float64Seq(idx + 1)
		} else {
			res = make([]float64, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

// Float64s extracts an array of float64 values from the data and skips all
// whitespace after it. The values must be enclosed in square brackets "[...]"
// and the values must be separated by commas.
func (r *Reader) Float64s() (res []float64, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []float64{}
			err = r.CloseArray()
		} else if res, err = r.float64Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

// Float64 extracts the next float64 value from the data and skips all
// whitespace after it.
func (r *Reader) Float64() (flt float64, err error) {
	dat, done := r.dat[r.pos:], false

	m, d, e, n, t, dp, sp, i, ok := readFloat(dat)
	if !ok {
		if i == len(dat) {
			return 0, newEndOfFileError()
		} else {
			return 0, newInvalidCharacterError(dat[i], r.pos+i)
		}
	}

	// adjusting exponent
	ae := e + dotExp(d, dp, sp, t)

	if d == 0 {
		flt, done = 0, true
	} else if !t {
		if f, ok := floatconv.Atof64exact(m, ae, n); ok {
			flt, done = f, true
		}
	}

	if !done {
		f, ok := floatconv.EiselLemire64(m, ae, n)
		if ok {
			if !t {
				flt, done = f, true
			}
			fu, ok := floatconv.EiselLemire64(m, ae, n)
			if ok && f == fu {
				flt, done = f, true
			}
		}
	}

	if !done {
		var dec floatconv.Decimal
		dec.Set(dat[sp:i], e, n, t, dp)

		b, ovf := dec.FloatBits(&floatconv.Float64info)
		flt = math.Float64frombits(b)
		if ovf {
			return flt, newNumberOutOfRangeError(dat[sp:i], r.pos)
		}
	}

	r.pos += i
	r.SkipWhiteSpace()
	return
}

// Float64p extracts the next float64 value and returns a pointer variable.
func (r *Reader) Float64p() (res *float64, err error) {
	if v, err := r.Float64(); err == nil {
		res = &v
	}
	return
}
