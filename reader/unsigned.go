package reader

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
