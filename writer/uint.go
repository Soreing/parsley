package writer

import (
	"encoding/base64"
	"strconv"
)

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
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
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
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
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
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
	}
}

func WriteUInt16s(dst []byte, ns []uint16) (ln int) {
	tmp, res := make([]byte, 0, 32), ([]byte)(nil)
	if ns != nil {
		dst[0] = '['
		ln++

		if len(ns) > 0 {
			res = strconv.AppendInt(tmp, int64(ns[0]), 10)
			ln += copy(dst[1:], res)
			for _, n := range ns[1:] {
				dst[ln] = ','
				ln++

				res = strconv.AppendInt(tmp, int64(n), 10)
				ln += copy(dst[ln:], res)
			}
		}

		dst[ln] = ']'
		return ln + 1
	} else {
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
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
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
	}
}

func WriteUInt32s(dst []byte, ns []uint32) (ln int) {
	tmp, res := make([]byte, 0, 32), ([]byte)(nil)
	if ns != nil {
		dst[0] = '['
		ln++

		if len(ns) > 0 {
			res = strconv.AppendInt(tmp, int64(ns[0]), 10)
			ln += copy(dst[1:], res)
			for _, n := range ns[1:] {
				dst[ln] = ','
				ln++

				res = strconv.AppendInt(tmp, int64(n), 10)
				ln += copy(dst[ln:], res)
			}
		}

		dst[ln] = ']'
		return ln + 1
	} else {
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
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
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
	}
}

func WriteUInt64s(dst []byte, ns []uint64) (ln int) {
	tmp, res := make([]byte, 0, 32), ([]byte)(nil)
	if ns != nil {
		dst[0] = '['
		ln++

		if len(ns) > 0 {
			res = strconv.AppendInt(tmp, int64(ns[0]), 10)
			ln += copy(dst[1:], res)
			for _, n := range ns[1:] {
				dst[ln] = ','
				ln++

				res = strconv.AppendInt(tmp, int64(n), 10)
				ln += copy(dst[ln:], res)
			}
		}

		dst[ln] = ']'
		return ln + 1
	} else {
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
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
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
	}
}

func WriteUInts(dst []byte, ns []uint) (ln int) {
	tmp, res := make([]byte, 0, 32), ([]byte)(nil)
	if ns != nil {
		dst[0] = '['
		ln++

		if len(ns) > 0 {
			res = strconv.AppendInt(tmp, int64(ns[0]), 10)
			ln += copy(dst[1:], res)
			for _, n := range ns[1:] {
				dst[ln] = ','
				ln++

				res = strconv.AppendInt(tmp, int64(n), 10)
				ln += copy(dst[ln:], res)
			}
		}

		dst[ln] = ']'
		return ln + 1
	} else {
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
	}
}
