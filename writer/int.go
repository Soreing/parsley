package writer

import "strconv"

func Int8Length(n int8) (ln int) {
	if n < 0 {
		return ui8dc(uint8(-n)) + 1
	} else {
		return ui8dc(uint8(n))
	}
}

func Int8sLength(ns []int8) (ln int) {
	for _, n := range ns {
		ln += Int8Length(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

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
		return copy(dst, "null")
	}
}

func WriteInt8s(dst []byte, ns []int8) (ln int) {
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

func Int16Length(n int16) (ln int) {
	if n < 0 {
		return ui16dc(uint16(-n)) + 1
	} else {
		return ui16dc(uint16(n))
	}
}

func Int16sLength(ns []int16) (ln int) {
	for _, n := range ns {
		ln += Int16Length(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
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
		return copy(dst, "null")
	}
}

func WriteInt16s(dst []byte, ns []int16) (ln int) {
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

func Int32Length(n int32) (ln int) {
	if n < 0 {
		return ui32dc(uint32(-n)) + 1
	} else {
		return ui32dc(uint32(n))
	}
}

func Int32sLength(ns []int32) (ln int) {
	for _, n := range ns {
		ln += Int32Length(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
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
		return copy(dst, "null")
	}
}

func WriteInt32s(dst []byte, ns []int32) (ln int) {
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

func Int64Length(n int64) (ln int) {
	if n < 0 {
		return ui64dc(-uint64(n)) + 1
	} else {
		return ui64dc(-uint64(n))
	}
}

func Int64sLength(ns []int64) (ln int) {
	for _, n := range ns {
		ln += Int64Length(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
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
		return copy(dst, "null")
	}
}

func WriteInt64s(dst []byte, ns []int64) (ln int) {
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

func IntLength(n int) (ln int) {
	if n < 0 {
		return ui32dc(uint32(-n)) + 1
	} else {
		return ui32dc(uint32(n))
	}
}

func IntsLength(ns []int) (ln int) {
	for _, n := range ns {
		ln += IntLength(n) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
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
		return copy(dst, "null")
	}
}

func WriteInts(dst []byte, ns []int) (ln int) {
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
