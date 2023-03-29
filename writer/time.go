package writer

import (
	"time"
)

func WriteTime(dst []byte, t time.Time) (ln int) {
	ln = copy(dst[1:], t.Format(time.RFC3339Nano)) + 1
	dst[0], dst[ln] = '"', '"'
	return ln + 1
}

func WriteTimePtr(dst []byte, t *time.Time) (ln int) {
	if t != nil {
		ln = copy(dst[1:], t.Format(time.RFC3339Nano)) + 1
		dst[0], dst[ln] = '"', '"'
		return ln + 1
	} else {
		dst[0] = 'n'
		dst[1] = 'u'
		dst[2] = 'l'
		dst[3] = 'l'
		return 4
	}
}

func WriteTimes(dst []byte, ts []time.Time) (ln int) {
	if ts != nil {
		dst[0] = '['
		ln++

		if len(ts) > 0 {
			ln += WriteTime(dst[ln:], ts[0])
			for _, t := range ts[1:] {
				dst[ln] = ','
				ln++
				ln += WriteTime(dst[ln:], t)
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
