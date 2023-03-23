package writer

import "strconv"

func WriteInt8(dst []byte, n int8) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		strconv.AppendInt(dst, int64(n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[0] = '0'
		return 1
	}
}

func WriteInt8Ptr(dst []byte, n *int8) (ln int) {
	if n != nil {
		tmp := make([]byte, 0, 32)
		strconv.AppendInt(dst, int64(*n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[ln+0] = 'n'
		dst[ln+1] = 'u'
		dst[ln+2] = 'l'
		dst[ln+3] = 'l'
		return ln + 4
	}
}

func WriteInt16(dst []byte, n int16) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		strconv.AppendInt(dst, int64(n), 10)
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
		strconv.AppendInt(dst, int64(*n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[ln+0] = 'n'
		dst[ln+1] = 'u'
		dst[ln+2] = 'l'
		dst[ln+3] = 'l'
		return ln + 4
	}
}
func WriteInt32(dst []byte, n int32) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		strconv.AppendInt(dst, int64(n), 10)
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
		strconv.AppendInt(dst, int64(*n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[ln+0] = 'n'
		dst[ln+1] = 'u'
		dst[ln+2] = 'l'
		dst[ln+3] = 'l'
		return ln + 4
	}
}
func WriteInt64(dst []byte, n int64) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		strconv.AppendInt(dst, int64(n), 10)
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
		strconv.AppendInt(dst, int64(*n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[ln+0] = 'n'
		dst[ln+1] = 'u'
		dst[ln+2] = 'l'
		dst[ln+3] = 'l'
		return ln + 4
	}
}
func WriteInt(dst []byte, n int) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		strconv.AppendInt(dst, int64(n), 10)
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
		strconv.AppendInt(dst, int64(*n), 10)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[ln+0] = 'n'
		dst[ln+1] = 'u'
		dst[ln+2] = 'l'
		dst[ln+3] = 'l'
		return ln + 4
	}
}
