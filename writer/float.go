package writer

import "strconv"

func Float32Length(n float32) (ln int) {
	return 24
}

func Float32sLength(ns []float32) (ln int) {
	for range ns {
		ln += 24 + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

func WriteFloat32(dst []byte, n float32) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendFloat(tmp, float64(n), 'g', -1, 32)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[0] = '0'
		return 1
	}
}

func WriteFloat32Ptr(dst []byte, n *float32) (ln int) {
	if n != nil {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendFloat(tmp, float64(*n), 'g', -1, 32)
		copy(dst, tmp)
		return len(tmp)
	} else {
		return copy(dst, "null")
	}
}

func WriteFloat32s(dst []byte, ns []float32) (ln int) {
	tmp, res := make([]byte, 0, 32), ([]byte)(nil)
	if len(ns) > 0 {
		ln = 1
		for _, n := range ns {
			res = strconv.AppendFloat(tmp, float64(n), 'g', -1, 32)
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

func Float64Length(n float64) (ln int) {
	return 24
}

func Float64sLength(ns []float64) (ln int) {
	for range ns {
		ln += 24 + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

func WriteFloat64(dst []byte, n float64) (ln int) {
	if n != 0 {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendFloat(tmp, float64(n), 'g', -1, 64)
		copy(dst, tmp)
		return len(tmp)
	} else {
		dst[0] = '0'
		return 1
	}
}

func WriteFloat64Ptr(dst []byte, n *float64) (ln int) {
	if n != nil {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendFloat(tmp, float64(*n), 'g', -1, 64)
		copy(dst, tmp)
		return len(tmp)
	} else {
		return copy(dst, "null")
	}
}

func WriteFloat64s(dst []byte, ns []float64) (ln int) {
	tmp, res := make([]byte, 0, 32), ([]byte)(nil)
	if len(ns) > 0 {
		ln = 1
		for _, n := range ns {
			res = strconv.AppendFloat(tmp, float64(n), 'g', -1, 64)
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
