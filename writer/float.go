package writer

import "strconv"

func WriteFloat32(dst []byte, n float32) (ln int) {
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

func WriteFloat32Ptr(dst []byte, n *float32) (ln int) {
	if n != nil {
		tmp := make([]byte, 0, 32)
		tmp = strconv.AppendFloat(tmp, float64(*n), 'g', -1, 64)
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

func WriteFloat32s(dst []byte, ns []float32) (ln int) {
	tmp, res := make([]byte, 0, 32), ([]byte)(nil)
	if ns != nil {
		dst[0] = '['
		ln++

		if len(ns) > 0 {
			res = strconv.AppendFloat(tmp, float64(ns[0]), 'g', -1, 64)
			ln += copy(dst[1:], res)
			for _, n := range ns[1:] {
				dst[ln] = ','
				ln++

				res = strconv.AppendFloat(tmp, float64(n), 'g', -1, 64)
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
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
	}
}

func WriteFloat64s(dst []byte, ns []float64) (ln int) {
	tmp, res := make([]byte, 0, 32), ([]byte)(nil)
	if ns != nil {
		dst[0] = '['
		ln++

		if len(ns) > 0 {
			res = strconv.AppendFloat(tmp, float64(ns[0]), 'g', -1, 64)
			ln += copy(dst[1:], res)
			for _, n := range ns[1:] {
				dst[ln] = ','
				ln++

				res = strconv.AppendFloat(tmp, float64(n), 'g', -1, 64)
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
