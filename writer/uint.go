package writer

import (
	"encoding/base64"
	"strconv"
)

func UInt8Length(n uint8) (ln int) {
	return ui8dc(uint8(n))
}

func UInt8sLength(ns []uint8) (ln int) {
	for _, n := range ns {
		ln += ui8dc(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

func WriteUInt8(dst []byte, n uint8) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(n), 10)
		return copy(dst, tmp)
	} else {
		dst[0] = '0'
		return 1
	}
}

func WriteUInt8Ptr(dst []byte, n *uint8) (ln int) {
	if n != nil {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(*n), 10)
		return copy(dst, tmp)
	} else {
		return copy(dst, "null")
	}
}

func WriteUInt8s(dst []byte, ns []uint8) (ln int) {
	if ns != nil {
		if len(ns) > 0 {
			base64.StdEncoding.Encode(dst[1:], ns)
			ln = (len(ns)+2)/3*4 + 2
			dst[0], dst[ln-1] = '"', '"'
			return
		} else {
			dst[0] = '"'
			dst[1] = '"'
			return 2
		}
	} else {
		return copy(dst, "null")
	}
}

func UInt16Length(n uint16) (ln int) {
	return ui16dc(uint16(n))
}

func UInt16sLength(ns []uint16) (ln int) {
	for _, n := range ns {
		ln += ui16dc(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

func WriteUInt16(dst []byte, n uint16) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(n), 10)
		return copy(dst, tmp)
	} else {
		dst[0] = '0'
		return 1
	}
}

func WriteUInt16Ptr(dst []byte, n *uint16) (ln int) {
	if n != nil {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(*n), 10)
		return copy(dst, tmp)
	} else {
		return copy(dst, "null")
	}
}

func WriteUInt16s(dst []byte, ns []uint16) (ln int) {
	tmp, res := make([]byte, 0, 32), ([]byte)(nil)
	if len(ns) > 0 {
		ln = 1
		for _, n := range ns {
			res = strconv.AppendInt(tmp, int64(n), 10)
			ln += copy(dst[ln:], res)
			dst[ln] = ','
			ln++
		}

		dst[0], dst[ln-1] = '[', ']'
		return ln
	} else if ns != nil {
		return copy(dst, "[]")
	} else {
		return copy(dst, "null")
	}
}

func UInt32Length(n uint32) (ln int) {
	return ui32dc(uint32(n))
}

func UInt32sLength(ns []uint32) (ln int) {
	for _, n := range ns {
		ln += ui32dc(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

func WriteUInt32(dst []byte, n uint32) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(n), 10)
		return copy(dst, tmp)
	} else {
		dst[0] = '0'
		return 1
	}
}

func WriteUInt32Ptr(dst []byte, n *uint32) (ln int) {
	if n != nil {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(*n), 10)
		return copy(dst, tmp)
	} else {
		return copy(dst, "null")
	}
}

func WriteUInt32s(dst []byte, ns []uint32) (ln int) {
	tmp, res := make([]byte, 0, 32), ([]byte)(nil)
	if len(ns) > 0 {
		ln = 1
		for _, n := range ns {
			res = strconv.AppendInt(tmp, int64(n), 10)
			ln += copy(dst[ln:], res)
			dst[ln] = ','
			ln++
		}

		dst[0], dst[ln-1] = '[', ']'
		return ln
	} else if ns != nil {
		return copy(dst, "[]")
	} else {
		return copy(dst, "null")
	}
}

func UInt64Length(n uint64) (ln int) {
	return ui64dc(uint64(n))
}

func UInt64sLength(ns []uint64) (ln int) {
	for _, n := range ns {
		ln += ui64dc(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

func WriteUInt64(dst []byte, n uint64) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(n), 10)
		return copy(dst, tmp)
	} else {
		dst[0] = '0'
		return 1
	}
}

func WriteUInt64Ptr(dst []byte, n *uint64) (ln int) {
	if n != nil {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(*n), 10)
		return copy(dst, tmp)
	} else {
		return copy(dst, "null")
	}
}

func WriteUInt64s(dst []byte, ns []uint64) (ln int) {
	tmp, res := make([]byte, 0, 32), ([]byte)(nil)
	if len(ns) > 0 {
		ln = 1
		for _, n := range ns {
			res = strconv.AppendInt(tmp, int64(n), 10)
			ln += copy(dst[ln:], res)
			dst[ln] = ','
			ln++
		}

		dst[0], dst[ln-1] = '[', ']'
		return ln
	} else if ns != nil {
		return copy(dst, "[]")
	} else {
		return copy(dst, "null")
	}
}

func UIntLength(n uint) (ln int) {
	return ui32dc(uint32(n))
}

func UIntsLength(ns []uint) (ln int) {
	for _, n := range ns {
		ln += ui32dc(uint32(n)) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

func WriteUInt(dst []byte, n uint) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(n), 10)
		return copy(dst, tmp)
	} else {
		dst[0] = '0'
		return 1
	}
}

func WriteUIntPtr(dst []byte, n *uint) (ln int) {
	if n != nil {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(*n), 10)
		return copy(dst, tmp)
	} else {
		return copy(dst, "null")
	}
}

func WriteUInts(dst []byte, ns []uint) (ln int) {
	tmp, res := make([]byte, 0, 32), ([]byte)(nil)
	if len(ns) > 0 {
		ln = 1
		for _, n := range ns {
			res = strconv.AppendInt(tmp, int64(n), 10)
			ln += copy(dst[ln:], res)
			dst[ln] = ','
			ln++
		}

		dst[0], dst[ln-1] = '[', ']'
		return ln
	} else if ns != nil {
		return copy(dst, "[]")
	} else {
		return copy(dst, "null")
	}
}
