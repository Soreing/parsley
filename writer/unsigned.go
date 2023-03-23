package writer

import "strconv"

func WriteUInt8(dst []byte, n uint8) (ln int) {
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

func WriteUInt8Ptr(dst []byte, n *uint8) (ln int) {
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

func WriteUInt16(dst []byte, n uint16) (ln int) {
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

func WriteUInt16Ptr(dst []byte, n *uint16) (ln int) {
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
func WriteUInt32(dst []byte, n uint32) (ln int) {
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

func WriteUInt32Ptr(dst []byte, n *uint32) (ln int) {
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
func WriteUInt64(dst []byte, n uint64) (ln int) {
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

func WriteUInt64Ptr(dst []byte, n *uint64) (ln int) {
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
func WriteUInt(dst []byte, n uint) (ln int) {
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

func WriteUIntPtr(dst []byte, n *uint) (ln int) {
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
