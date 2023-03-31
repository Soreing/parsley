package writer

import (
	"time"
)

func FastRFC3339NanoLength(t time.Time) (ln int) {
	nano, nanol := t.Nanosecond(), 10

	if nano > 0 {
		for nano%10 == 0 && nanol > 0 {
			nanol--
			nano /= 10
		}
	} else {
		nanol = 0
	}

	_, zone := t.Zone()
	if zone == 0 {
		return 20 + nanol
	} else {
		return 25 + nanol
	}
}

func WriteFastRFC3339Nano(dst []byte, t time.Time) (ln int) {
	year, month, day := t.Date()
	hour, minute, second := t.Clock()
	nano, nanol := t.Nanosecond(), 10
	zonel, zsig := 1, byte('Z')

	_, zone := t.Zone()
	if zone > 0 {
		zonel, zsig = 6, '+'
	} else if zone < 0 {
		zone, zonel, zsig = -zone, 6, '-'
	}

	if nano > 0 {
		for nano%10 == 0 && nanol > 0 {
			nanol--
			nano /= 10
		}
	} else {
		nanol = 0
	}

	// Year
	dst[3] = byte(year%10) + '0'
	year /= 10
	dst[2] = byte(year%10) + '0'
	year /= 10
	dst[1] = byte(year%10) + '0'
	year /= 10
	dst[0] = byte(year) + '0'
	// Month
	dst[6] = byte(month%10) + '0'
	month /= 10
	dst[5] = byte(month%10) + '0'
	// Day
	dst[9] = byte(day%10) + '0'
	day /= 10
	dst[8] = byte(day%10) + '0'
	// Hour
	dst[12] = byte(hour%10) + '0'
	hour /= 10
	dst[11] = byte(hour%10) + '0'
	// Minute
	dst[15] = byte(minute%10) + '0'
	minute /= 10
	dst[14] = byte(minute%10) + '0'
	// Second
	dst[18] = byte(second%10) + '0'
	second /= 10
	dst[17] = byte(second%10) + '0'

	// Nanoseconds
	if nano > 0 {
		dst[19] = '.'
		for i := nanol + 18; i >= 20; i-- {
			dst[i] = byte(nano%10) + '0'
			nano /= 10
		}
	}

	// Timezone
	if zone != 0 {
		hr, mn := zone/3600, zone%3600
		dst[nanol+24] = byte(mn%600) + '0'
		mn /= 600
		dst[nanol+23] = byte(mn) + '0'
		dst[nanol+22] = ':'
		dst[nanol+21] = byte(hr%10) + '0'
		hr /= 10
		dst[nanol+20] = byte(hr) + '0'
	}

	// Dressing
	dst[4], dst[7], dst[10], dst[13], dst[16] = '-', '-', 'T', ':', ':'
	dst[19+nanol] = zsig
	return 19 + nanol + zonel
}

func WriteTime(dst []byte, t time.Time) (ln int) {
	ln = WriteFastRFC3339Nano(dst, t) + 1
	dst[0], dst[ln] = '"', '"'
	return ln + 1
}

func WriteTimePtr(dst []byte, t *time.Time) (ln int) {
	if t != nil {
		ln = WriteFastRFC3339Nano(dst, *t) + 1
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
