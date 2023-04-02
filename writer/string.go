package writer

const hexDigits = "0123456789ABCDEF"

func StringLength(str string) (ln int) {
	ln = len(str) + 2
	for _, c := range str {
		if c == '"' || c == '\\' || c == '\t' || c == '\n' || c == '\r' {
			ln++
		} else if c <= 0x1F {
			ln += 5
		}
	}
	return
}

func StringsLength(strs []string) (ln int) {
	for _, s := range strs {
		ln += StringLength(s) + 1
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

func WriteString(dst []byte, s string) (ln int) {
	dst[0] = '"'
	pos := 1
	for i := 0; i < len(s); i++ {
		switch c := s[i]; c {
		case '"':
			dst[pos] = '\\'
			dst[pos+1] = '"'
			pos += 2
		case '\\':
			dst[pos] = '\\'
			dst[pos+1] = '\\'
			pos += 2
		case '\n':
			dst[pos] = '\\'
			dst[pos+1] = 'n'
			pos += 2
		case '\r':
			dst[pos] = '\\'
			dst[pos+1] = 'r'
			pos += 2
		case '\t':
			dst[pos] = '\\'
			dst[pos+1] = 't'
			pos += 2
		default:
			if c <= 0x1F {
				dst[pos+0] = '\\'
				dst[pos+1] = 'u'
				dst[pos+2] = '0'
				dst[pos+3] = '0'
				dst[pos+4] = c>>4 + '0'
				dst[pos+5] = hexDigits[c&0xF]
				pos += 6
			} else {
				dst[pos] = c
				pos++
			}
		}
	}
	dst[pos] = '"'
	return pos + 1
}

func WriteStringPtr(dst []byte, s *string) (ln int) {
	if s != nil {
		return ln + WriteString(dst[ln:], *s)
	} else {
		return copy(dst, "null")
	}
}

func WriteStrings(dst []byte, ss []string) (ln int) {
	if len(ss) > 0 {
		ln = 1
		for _, s := range ss {
			ln += WriteString(dst[ln:], s)
			dst[ln] = ','
			ln++
		}

		dst[0], dst[ln-1] = '[', ']'
		return ln
	} else if ss != nil {
		return copy(dst, "[]")
	} else {
		return copy(dst, "null")
	}
}
