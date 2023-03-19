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

	r.skipWhiteSpace()
	return sig * (intg + frc), nil
}

func (r *Reader) getInt() (int64, error) {
	sig, intg := int64(1), int64(0)

	// Reading the sign
	if r.pos >= len(r.dat) {
		return 0, NewEndOfFileError()
	} else if r.dat[r.pos] == '-' {
		sig = -1
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
			intg = intg*10 + int64(r.dat[r.pos]-'0')
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
		case '}', ']', ',', ' ', '\t', '\n', '\r':
			// Empty //
		default:
			return 0, NewInvalidCharacterError(r.dat[r.pos], r.pos)
		}
	}

	r.skipWhiteSpace()
	return sig * intg, nil
}

func (r *Reader) getUInt() (uint64, error) {
	intg := uint64(0)

	// Reading the integer part
	if r.pos >= len(r.dat) {
		return 0, NewEndOfFileError()
	} else if r.dat[r.pos] == '0' {
		r.pos++
	} else if r.dat[r.pos] >= '1' && r.dat[r.pos] <= '9' {
		for r.pos < len(r.dat) &&
			r.dat[r.pos] >= '0' &&
			r.dat[r.pos] <= '9' {
			intg = intg*10 + uint64(r.dat[r.pos]-'0')
			r.pos++
		}
	} else {
		return 0, NewInvalidCharacterError(r.dat[r.pos], r.pos)
	}

	// Reading the fraction part
	if r.pos >= len(r.dat) {
		return intg, nil
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
					return 0, NewInvalidCharacterError(r.dat[r.pos], r.pos)
				} else {
					return 0, NewEndOfFileError()
				}
			}
		case '}', ']', ',', ' ', '\t', '\n', '\r':
			// Empty //
		default:
			return 0, NewInvalidCharacterError(r.dat[r.pos], r.pos)
		}
	}

	r.skipWhiteSpace()
	return intg, nil
}

func (r *Reader) intSeq(idx int) (res []int, err error) {
	var n int
	if n, err = r.GetInt(); err == nil {
		if r.Next() {
			res, err = r.intSeq(idx + 1)
		} else {
			res = make([]int, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

func (r *Reader) GetInts() (res []int, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.intSeq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetInt() (int, error) {
	if n, err := r.getInt(); err != nil {
		return 0, err
	} else {
		return int(n), nil
	}
}

func (r *Reader) GetIntPtr() (res *int, err error) {
	if v, err := r.GetInt(); err == nil {
		res = &v
	}
	return
}

func (r *Reader) int8Seq(idx int) (res []int8, err error) {
	var n int8
	if n, err = r.GetInt8(); err == nil {
		if r.Next() {
			res, err = r.int8Seq(idx + 1)
		} else {
			res = make([]int8, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

func (r *Reader) GetInt8s() (res []int8, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.int8Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetInt8() (int8, error) {
	if n, err := r.getInt(); err != nil {
		return 0, err
	} else {
		return int8(n), nil
	}
}

func (r *Reader) GetInt8Ptr() (res *int8, err error) {
	if v, err := r.GetInt8(); err == nil {
		res = &v
	}
	return
}

func (r *Reader) int16Seq(idx int) (res []int16, err error) {
	var n int16
	if n, err = r.GetInt16(); err == nil {
		if r.Next() {
			res, err = r.int16Seq(idx + 1)
		} else {
			res = make([]int16, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

func (r *Reader) GetInt16s() (res []int16, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.int16Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetInt16() (int16, error) {
	if n, err := r.getInt(); err != nil {
		return 0, err
	} else {
		return int16(n), nil
	}
}

func (r *Reader) GetInt16Ptr() (res *int16, err error) {
	if v, err := r.GetInt16(); err == nil {
		res = &v
	}
	return
}

func (r *Reader) int32Seq(idx int) (res []int32, err error) {
	var n int32
	if n, err = r.GetInt32(); err == nil {
		if r.Next() {
			res, err = r.int32Seq(idx + 1)
		} else {
			res = make([]int32, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

func (r *Reader) GetInt32s() (res []int32, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.int32Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetInt32() (int32, error) {
	if n, err := r.getInt(); err != nil {
		return 0, err
	} else {
		return int32(n), nil
	}
}

func (r *Reader) GetInt32Ptr() (res *int32, err error) {
	if v, err := r.GetInt32(); err == nil {
		res = &v
	}
	return
}

func (r *Reader) int64Seq(idx int) (res []int64, err error) {
	var n int64
	if n, err = r.GetInt64(); err == nil {
		if r.Next() {
			res, err = r.int64Seq(idx + 1)
		} else {
			res = make([]int64, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

func (r *Reader) GetInt64s() (res []int64, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.int64Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetInt64() (int64, error) {
	return r.getInt()
}

func (r *Reader) GetInt64Ptr() (res *int64, err error) {
	if v, err := r.GetInt64(); err == nil {
		res = &v
	}
	return
}

func (r *Reader) uintSeq(idx int) (res []uint, err error) {
	var n uint
	if n, err = r.GetUInt(); err == nil {
		if r.Next() {
			res, err = r.uintSeq(idx + 1)
		} else {
			res = make([]uint, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

func (r *Reader) GetUInts() (res []uint, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.uintSeq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetUInt() (uint, error) {
	if n, err := r.getUInt(); err != nil {
		return 0, err
	} else {
		return uint(n), nil
	}
}

func (r *Reader) GetUIntPtr() (res *uint, err error) {
	if v, err := r.GetUInt(); err == nil {
		res = &v
	}
	return
}

func (r *Reader) uint8Seq(idx int) (res []uint8, err error) {
	var n uint8
	if n, err = r.GetUInt8(); err == nil {
		if r.Next() {
			res, err = r.uint8Seq(idx + 1)
		} else {
			res = make([]uint8, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

func (r *Reader) GetUInt8s() (res []uint8, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.uint8Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetUInt8() (uint8, error) {
	if n, err := r.getUInt(); err != nil {
		return 0, err
	} else {
		return uint8(n), nil
	}
}

func (r *Reader) GetUInt8Ptr() (res *uint8, err error) {
	if v, err := r.GetUInt8(); err == nil {
		res = &v
	}
	return
}

func (r *Reader) uint16Seq(idx int) (res []uint16, err error) {
	var n uint16
	if n, err = r.GetUInt16(); err == nil {
		if r.Next() {
			res, err = r.uint16Seq(idx + 1)
		} else {
			res = make([]uint16, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

func (r *Reader) GetUInt16s() (res []uint16, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.uint16Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetUInt16() (uint16, error) {
	if n, err := r.getUInt(); err != nil {
		return 0, err
	} else {
		return uint16(n), nil
	}
}

func (r *Reader) GetUInt16Ptr() (res *uint16, err error) {
	if v, err := r.GetUInt16(); err == nil {
		res = &v
	}
	return
}

func (r *Reader) uint32Seq(idx int) (res []uint32, err error) {
	var n uint32
	if n, err = r.GetUInt32(); err == nil {
		if r.Next() {
			res, err = r.uint32Seq(idx + 1)
		} else {
			res = make([]uint32, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

func (r *Reader) GetUInt32s() (res []uint32, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.uint32Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetUInt32() (uint32, error) {
	if n, err := r.getUInt(); err != nil {
		return 0, err
	} else {
		return uint32(n), nil
	}
}

func (r *Reader) GetUInt32Ptr() (res *uint32, err error) {
	if v, err := r.GetUInt32(); err == nil {
		res = &v
	}
	return
}

func (r *Reader) uint64Seq(idx int) (res []uint64, err error) {
	var n uint64
	if n, err = r.GetUInt64(); err == nil {
		if r.Next() {
			res, err = r.uint64Seq(idx + 1)
		} else {
			res = make([]uint64, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

func (r *Reader) GetUInt64s() (res []uint64, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.uint64Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetUInt64() (uint64, error) {
	return r.getUInt()
}

func (r *Reader) GetUInt64Ptr() (res *uint64, err error) {
	if v, err := r.GetUInt64(); err == nil {
		res = &v
	}
	return
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
