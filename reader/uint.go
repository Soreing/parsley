package reader

import (
	"encoding/base64"
	"math"
)

func (r *Reader) uintSeq(idx int) (res []uint, err error) {
	var n uint
	if n, err = r.UInt(); err == nil {
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

func (r *Reader) UInts() (res []uint, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []uint{}
			err = r.CloseArray()
		} else if res, err = r.uintSeq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) UInt() (n uint, err error) {
	dat := r.dat[r.pos:]
	num, neg, pos, ok := readInteger(dat)
	if !ok {
		if num == 0 && neg {
			return 0, NewNumberOutOfRangeError(dat[:pos], r.pos)
		} else if pos == len(dat) {
			return 0, NewEndOfFileError()
		} else {
			return 0, NewInvalidCharacterError(dat[pos], r.pos+pos)
		}
	} else if !neg && num <= math.MaxUint32 {
		n = uint(num)
	} else {
		return 0, NewNumberOutOfRangeError(dat[:pos], r.pos)
	}

	r.pos += pos
	r.SkipWhiteSpace()
	return
}

func (r *Reader) UIntp() (res *uint, err error) {
	if v, err := r.UInt(); err == nil {
		res = &v
	}
	return
}

func (r *Reader) UInt8s() (res []uint8, err error) {
	dat := r.dat[r.pos:]

	if len(dat) < 2 {
		return nil, NewEndOfFileError()
	} else if dat[0] != '"' {
		return nil, NewInvalidCharacterError(dat[0], r.pos)
	}

	beg, end, c := 1, 1, dat[1]
	for dat[end] != '"' {
		if end == len(dat)-1 {
			return nil, NewEndOfFileError()
		}
		if c-'A' < 26 || c-'a' < 26 || c-'0' < 10 ||
			c == '+' || c == '/' || c == '=' {
			end++
			c = dat[end]
		} else {
			return nil, NewInvalidCharacterError(c, r.pos+end)
		}
	}

	if (end-beg)%4 != 0 {
		return nil, NewBase64PaddingError(r.pos + end)
	}

	bytes := (end - beg) / 4 * 3
	if bytes > 2 && dat[end-1] == '=' {
		bytes--
	}
	if bytes > 1 && dat[end-2] == '=' {
		bytes--
	}

	dst := make([]byte, bytes)
	if _, err := base64.StdEncoding.Decode(dst, dat[beg:end]); err != nil {
		return nil, err
	}

	r.pos += end + 1
	r.SkipWhiteSpace()
	return dst, nil
}

func (r *Reader) UInt8() (n uint8, err error) {
	dat := r.dat[r.pos:]
	num, neg, pos, ok := readInteger(dat)
	if !ok {
		if num == 0 && neg {
			return 0, NewNumberOutOfRangeError(dat[:pos], r.pos)
		} else if pos == len(dat) {
			return 0, NewEndOfFileError()
		} else {
			return 0, NewInvalidCharacterError(dat[pos], r.pos+pos)
		}
	} else if !neg && num <= math.MaxUint8 {
		n = uint8(num)
	} else {
		return 0, NewNumberOutOfRangeError(dat[:pos], r.pos)
	}

	r.pos += pos
	r.SkipWhiteSpace()
	return
}

func (r *Reader) UInt8p() (res *uint8, err error) {
	if v, err := r.UInt8(); err == nil {
		res = &v
	}
	return
}

func (r *Reader) uint16Seq(idx int) (res []uint16, err error) {
	var n uint16
	if n, err = r.UInt16(); err == nil {
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

func (r *Reader) UInt16s() (res []uint16, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []uint16{}
			err = r.CloseArray()
		} else if res, err = r.uint16Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) UInt16() (n uint16, err error) {
	dat := r.dat[r.pos:]
	num, neg, pos, ok := readInteger(dat)
	if !ok {
		if num == 0 && neg {
			return 0, NewNumberOutOfRangeError(dat[:pos], r.pos)
		} else if pos == len(dat) {
			return 0, NewEndOfFileError()
		} else {
			return 0, NewInvalidCharacterError(dat[pos], r.pos+pos)
		}
	} else if !neg && num <= math.MaxUint16 {
		n = uint16(num)
	} else {
		return 0, NewNumberOutOfRangeError(dat[:pos], r.pos)
	}

	r.pos += pos
	r.SkipWhiteSpace()
	return
}

func (r *Reader) UInt16p() (res *uint16, err error) {
	if v, err := r.UInt16(); err == nil {
		res = &v
	}
	return
}

func (r *Reader) uint32Seq(idx int) (res []uint32, err error) {
	var n uint32
	if n, err = r.UInt32(); err == nil {
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

func (r *Reader) UInt32s() (res []uint32, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []uint32{}
			err = r.CloseArray()
		} else if res, err = r.uint32Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) UInt32() (n uint32, err error) {
	dat := r.dat[r.pos:]
	num, neg, pos, ok := readInteger(dat)
	if !ok {
		if num == 0 && neg {
			return 0, NewNumberOutOfRangeError(dat[:pos], r.pos)
		} else if pos == len(dat) {
			return 0, NewEndOfFileError()
		} else {
			return 0, NewInvalidCharacterError(dat[pos], r.pos+pos)
		}
	} else if !neg && num <= math.MaxUint32 {
		n = uint32(num)
	} else {
		return 0, NewNumberOutOfRangeError(dat[:pos], r.pos)
	}

	r.pos += pos
	r.SkipWhiteSpace()
	return
}

func (r *Reader) UInt32p() (res *uint32, err error) {
	if v, err := r.UInt32(); err == nil {
		res = &v
	}
	return
}

func (r *Reader) uint64Seq(idx int) (res []uint64, err error) {
	var n uint64
	if n, err = r.UInt64(); err == nil {
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

func (r *Reader) UInt64s() (res []uint64, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []uint64{}
			err = r.CloseArray()
		} else if res, err = r.uint64Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) UInt64() (n uint64, err error) {
	dat := r.dat[r.pos:]
	num, neg, pos, ok := readInteger(dat)
	if !ok {
		if num == 0 && neg {
			return 0, NewNumberOutOfRangeError(dat[:pos], r.pos)
		} else if pos == len(dat) {
			return 0, NewEndOfFileError()
		} else {
			return 0, NewInvalidCharacterError(dat[pos], r.pos+pos)
		}
	} else if !neg {
		n = uint64(num)
	} else {
		return 0, NewNumberOutOfRangeError(dat[:pos], r.pos)
	}

	r.pos += pos
	r.SkipWhiteSpace()
	return
}

func (r *Reader) UInt64p() (res *uint64, err error) {
	if v, err := r.UInt64(); err == nil {
		res = &v
	}
	return
}
