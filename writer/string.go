package writer

const hexDigits = "0123456789ABCDEF"

var controlLength = [32]int{
	5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 5, 5, 1, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
}

func StringSpace(str string) (ln int) {
	ln = len(str)
	for _, c := range str {
		if c <= 0x1F {
			ln += controlLength[c]
		} else if c == '"' || c == '\\' {
			ln++
		}
	}
	return
}

func StringsSpace(strs []string) (ln int) {
	for _, s := range strs {
		ln += StringSpace(s) + 3
	}
	if ln == 0 {
		return 2
	} else {
		return ln + 1
	}
}

func StringFieldSpace(f string, s string) (ln int) {
	if s == "" {
		return StringSpace(f) + 6
	} else {
		return StringSpace(f) + 6 + StringSpace(s)
	}
}

func StringPtrFieldSpace(f string, s *string) (ln int) {
	if s == nil {
		return StringSpace(f) + 8
	} else {
		return StringSpace(f) + 6 + StringSpace(*s)
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
