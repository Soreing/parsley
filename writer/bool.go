package writer

func BoolLength(b bool) (ln int) {
	if b {
		return 4
	} else {
		return 5
	}
}

func BoolsLength(bs []bool) (ln int) {
	ln = 6 * len(bs)
	for _, b := range bs {
		if b {
			ln--
		}
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

func WriteBool(dst []byte, b bool) (ln int) {
	if b {
		return copy(dst, "true")
	} else {
		return copy(dst, "false")
	}
}

func WriteBoolPtr(dst []byte, b *bool) (ln int) {
	if b == nil {
		return copy(dst, "null")
	} else if *b {
		return copy(dst, "true")
	} else {
		return copy(dst, "false")
	}
}

func WriteBools(dst []byte, bs []bool) (ln int) {
	if len(bs) > 0 {
		ldst := dst[1:]
		for _, b := range bs {
			if b {
				ln += copy(ldst, "true,")
				ldst = ldst[5:]
			} else {
				ln += copy(ldst, "false,")
				ldst = ldst[6:]
			}
		}
		dst[0], dst[ln] = '[', ']'
		return ln + 1
	} else if bs != nil {
		return copy(dst, "[]")
	} else {
		return copy(dst, "null")
	}
}
