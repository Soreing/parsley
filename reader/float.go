package reader

import (
	"math"

	"github.com/Soreing/parsley/reader/floatconv"
)

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

func (r *Reader) Float32() (flt float32, err error) {
	dat, ok, done := r.dat[r.pos:], true, false

	m, d, e, n, t, dp, sp, i, ok := readFloat(dat)
	if !ok {
		if i == len(dat) {
			return 0, NewEndOfFileError()
		} else {
			return 0, NewInvalidCharacterError(dat[i], r.pos+i)
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
			return flt, NewNumberOutOfRangeError(dat[sp:i], r.pos)
		}
	}

	r.pos += i
	r.SkipWhiteSpace()
	return
}

func (r *Reader) Float32p() (res *float32, err error) {
	if v, err := r.Float32(); err == nil {
		res = &v
	}
	return
}

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

func (r *Reader) Float64() (flt float64, err error) {
	dat, ok, done := r.dat[r.pos:], true, false

	m, d, e, n, t, dp, sp, i, ok := readFloat(dat)
	if !ok {
		if i == len(dat) {
			return 0, NewEndOfFileError()
		} else {
			return 0, NewInvalidCharacterError(dat[i], r.pos+i)
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
			return flt, NewNumberOutOfRangeError(dat[sp:i], r.pos)
		}
	}

	r.pos += i
	r.SkipWhiteSpace()
	return
}

func (r *Reader) Float64p() (res *float64, err error) {
	if v, err := r.Float64(); err == nil {
		res = &v
	}
	return
}
