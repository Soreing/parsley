package writer

import "strconv"

func WriteInt8(dst []byte, n int8) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(n), 10)
		return copy(dst, tmp)
	} else {
		dst[0] = '0'
		return 1
	}
}

func WriteInt8Ptr(dst []byte, n *int8) (ln int) {
	if n != nil {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(*n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
	}
}

func WriteInt8s(dst []byte, ns []int8) (ln int) {
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

func WriteInt16(dst []byte, n int16) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[0] = '0'
		return 1
	}
}

func WriteInt16Ptr(dst []byte, n *int16) (ln int) {
	if n != nil {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(*n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
	}
}

func WriteInt16s(dst []byte, ns []int16) (ln int) {
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

func WriteInt32(dst []byte, n int32) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[0] = '0'
		return 1
	}
}

func WriteInt32Ptr(dst []byte, n *int32) (ln int) {
	if n != nil {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(*n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
	}
}

func WriteInt32s(dst []byte, ns []int32) (ln int) {
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

func WriteInt64(dst []byte, n int64) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[0] = '0'
		return 1
	}
}

func WriteInt64Ptr(dst []byte, n *int64) (ln int) {
	if n != nil {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(*n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
	}
}

func WriteInt64s(dst []byte, ns []int64) (ln int) {
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

func WriteInt(dst []byte, n int) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[0] = '0'
		return 1
	}
}

func WriteIntPtr(dst []byte, n *int) (ln int) {
	if n != nil {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendInt(tmp, int64(*n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
	}
}

func WriteInts(dst []byte, ns []int) (ln int) {
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
