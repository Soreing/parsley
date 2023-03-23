package writer

func WriteBool(dst []byte, b bool) (ln int) {
	if b {
		dst[0] = 't'
		dst[1] = 'r'
		dst[2] = 'u'
		dst[3] = 'e'
		return 4
	} else {
		dst[0] = 'f'
		dst[1] = 'a'
		dst[2] = 'l'
		dst[3] = 's'
		dst[4] = 'e'
		return 5
	}
}

func WriteBoolPtr(dst []byte, b *bool) (ln int) {
	if b == nil {
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
	} else if *b {
		dst[0] = 't'
		dst[1] = 'r'
		dst[2] = 'u'
		dst[3] = 'e'
		return 4
	} else {
		dst[0] = 'f'
		dst[1] = 'a'
		dst[2] = 'l'
		dst[3] = 's'
		dst[4] = 'e'
		return 5
	}
}

func WriteBools(dst []byte, bs []bool) (ln int) {
	if bs != nil {
		dst[0] = '['
		ln++

		if len(bs) > 0 {
			if bs[0] {
				dst[0] = 't'
				dst[1] = 'r'
				dst[2] = 'u'
				dst[3] = 'e'
				dst = dst[4:]
				ln += 4
			} else {
				dst[0] = 'f'
				dst[1] = 'a'
				dst[2] = 'l'
				dst[3] = 's'
				dst[4] = 'e'
				dst = dst[5:]
				ln += 5
			}
			for _, b := range bs[1:] {
				if b {
					dst[0] = ','
					dst[1] = 't'
					dst[2] = 'r'
					dst[3] = 'u'
					dst[4] = 'e'
					dst = dst[5:]
					ln += 5
				} else {
					dst[0] = ','
					dst[1] = 'f'
					dst[2] = 'a'
					dst[3] = 'l'
					dst[4] = 's'
					dst[5] = 'e'
					dst = dst[6:]
					ln += 6
				}
			}
		}

		dst[0] = ']'
		return ln + 1
	} else {
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
	}
}
