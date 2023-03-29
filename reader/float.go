package reader

// TODO: Exponents are not implemented and making float64 could do with work
func (r *Reader) skipNumber() error {
	// Reading the sign
	if r.pos >= len(r.dat) {
		return NewEndOfFileError()
	} else if r.dat[r.pos] == '-' {
		r.pos++
	}

	// Reading the integer part
	if r.pos >= len(r.dat) {
		return NewEndOfFileError()
	} else if r.dat[r.pos] == '0' {
		r.pos++
	} else if r.dat[r.pos] >= '1' && r.dat[r.pos] <= '9' {
		for r.pos < len(r.dat) &&
			r.dat[r.pos] >= '0' &&
			r.dat[r.pos] <= '9' {
			r.pos++
		}
	} else {
		return NewInvalidCharacterError(r.dat[r.pos], r.pos)
	}

	// Reading the fraction part
	if r.pos >= len(r.dat) {
		return nil
	} else {
		switch r.dat[r.pos] {
		case '.':
			r.pos++
			dgt := 0
			for r.pos < len(r.dat) &&
				r.dat[r.pos] >= '0' &&
				r.dat[r.pos] <= '9' {
				r.pos++
				dgt++
			}
			if dgt == 0 {
				if r.pos < len(r.dat) {
					return NewInvalidCharacterError(r.dat[r.pos], r.pos)
				} else {
					return NewEndOfFileError()
				}
			}
		case '}', ']', ',', ' ', '\t', '\n', '\r':
			// Empty //
		default:
			return NewInvalidCharacterError(r.dat[r.pos], r.pos)
		}
	}

	return nil
}

func (r *Reader) getFloat() (float64, error) {
	sig, intg, frc := 1.0, 0.0, 0.0

	// Reading the sign
	if r.pos >= len(r.dat) {
		return 0, NewEndOfFileError()
	} else if r.dat[r.pos] == '-' {
		sig = -1.0
		r.pos++
	}

	// Reading the integer part
	if r.pos >= len(r.dat) {
		return 0, NewEndOfFileError()
	} else if r.dat[r.pos] == '0' {
		r.pos++
	} else if r.dat[r.pos] >= '1' && r.dat[r.pos] <= '9' {
		for r.pos < len(r.dat) &&
			r.dat[r.pos] >= '0' &&
			r.dat[r.pos] <= '9' {
			intg = intg*10 + float64(r.dat[r.pos]-'0')
			r.pos++
		}
	} else {
		return 0, NewInvalidCharacterError(r.dat[r.pos], r.pos)
	}

	// Reading the fraction part
	if r.pos >= len(r.dat) {
		return sig * intg, nil
	} else {
		switch r.dat[r.pos] {
		case '.':
			r.pos++
			dgt := 0
			for r.pos < len(r.dat) &&
				r.dat[r.pos] >= '0' &&
				r.dat[r.pos] <= '9' {
				frc = frc*10 + float64(r.dat[r.pos]-'0')
				r.pos++
				dgt++
			}
			if dgt == 0 {
				if r.pos < len(r.dat) {
					return 0, NewInvalidCharacterError(r.dat[r.pos], r.pos)
				} else {
					return 0, NewEndOfFileError()
				}
			}
			for dgt > 0 {
				frc = frc / 10
				dgt--
			}
		case '}', ']', ',', ' ', '\t', '\n', '\r':
			// Empty //
		default:
			return 0, NewInvalidCharacterError(r.dat[r.pos], r.pos)
		}
	}

	r.SkipWhiteSpace()
	return sig * (intg + frc), nil
}

func (r *Reader) float32Seq(idx int) (res []float32, err error) {
	var n float32
	if n, err = r.GetFloat32(); err == nil {
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

func (r *Reader) GetFloat32s() (res []float32, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.float32Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetFloat32() (float32, error) {
	if n, err := r.getFloat(); err != nil {
		return 0, err
	} else {
		return float32(n), nil
	}
}

func (r *Reader) GetFloat32Ptr() (res *float32, err error) {
	if v, err := r.GetFloat32(); err == nil {
		res = &v
	}
	return
}

func (r *Reader) float64Seq(idx int) (res []float64, err error) {
	var n float64
	if n, err = r.GetFloat64(); err == nil {
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

func (r *Reader) GetFloat64s() (res []float64, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.float64Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetFloat64() (float64, error) {
	return r.getFloat()
}

func (r *Reader) GetFloat64Ptr() (res *float64, err error) {
	if v, err := r.GetFloat64(); err == nil {
		res = &v
	}
	return
}
