package writer

// Added for import only
type Writer struct{}

func WriteNull(dst []byte) (ln int) {
	dst[0] = 'n'
	dst[1] = 'u'
	dst[2] = 'l'
	dst[3] = 'l'
	return 4
}
