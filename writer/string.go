package writer

func StringSpace(str string) (ln int) {
	ln = len(str)
	for _, e := range str {
		switch e {
		case '"', '\\', '/', '\b', '\f', '\n', '\r', '\t':
			ln++
		}
	}
	return
}

func StringsSpace(strs []string) (ln int) {
	ln = 1
	for _, s := range strs {
		ln += StringSpace(s) + 3
	}
	if ln == 1 {
		return 2
	} else {
		return ln
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
	for i := range s {
		switch s[i] {
		case '"':
			dst[pos] = '\\'
			dst[pos+1] = '"'
			pos++
		case '\\':
			dst[pos] = '\\'
			dst[pos+1] = '\\'
			pos++
		case '/':
			dst[pos] = '\\'
			dst[pos+1] = '/'
			pos++
		case '\b':
			dst[pos] = '\\'
			dst[pos+1] = 'b'
			pos++
		case '\f':
			dst[pos] = '\\'
			dst[pos+1] = 'f'
			pos++
		case '\n':
			dst[pos] = '\\'
			dst[pos+1] = 'n'
			pos++
		case '\r':
			dst[pos] = '\\'
			dst[pos+1] = 'r'
			pos++
		case '\t':
			dst[pos] = '\\'
			dst[pos+1] = 't'
			pos++
		default:
			dst[pos] = s[i]
		}
		pos++
	}
	dst[pos] = '"'
	return pos + 1
}

func WriteStringPtr(dst []byte, s *string) (ln int) {
	if s != nil {
		return ln + WriteString(dst[ln:], *s)
	} else {
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
	}
}

func WriteStrings(dst []byte, ss []string) (ln int) {
	if ss != nil {
		dst[0] = '['
		ln++

		if len(ss) > 0 {
			ln += WriteString(dst[1:], ss[0])
			for _, s := range ss[1:] {
				dst[ln] = ','
				ln++
				ln += WriteString(dst[ln:], s)
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
