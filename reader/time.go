package reader

import "time"

// timeSeq extracts time values recursively until the closing bracket
// is found, then assigns the elements to the allocated slice.
func (r *Reader) timeSeq(idx int) (res []time.Time, err error) {
	var bs []byte
	if bs, err = r.Bytes(); err == nil {
		if r.Next() {
			res, err = r.timeSeq(idx + 1)
		} else {
			res = make([]time.Time, idx+1)
		}
		if err == nil {
			res[idx], err = time.Parse(time.RFC3339, string(bs))
		}
	}
	return
}

// Times extracts an array of time values from the data and skips all
// whitespace after it. The values must be enclosed in square brackets "[...]"
// and the values must be separated by commas.
func (r *Reader) Times() (res []time.Time, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []time.Time{}
			err = r.CloseArray()
		} else if res, err = r.timeSeq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

// Time extracts the next time value from the data and skips all
// whitespace after it. All standard time formats are processed, but
// RFC3339Nano is prioritized.
func (r *Reader) Time() (time.Time, error) {
	rpos := r.pos
	if bs, err := r.Bytes(); err != nil {
		return time.Time{}, err
	} else if len(bs) == 0 {
		return time.Time{}, newUnknownTimeFormatError(string(bs), rpos)
	} else if bs[0]-'0' < 10 {
		if tm, err := time.Parse(time.RFC3339Nano, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.RFC822, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.RFC822Z, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.Kitchen, string(bs)); err == nil {
			return tm, nil
		} else {
			return tm, newUnknownTimeFormatError(string(bs), rpos)
		}
	} else {
		if tm, err := time.Parse(time.ANSIC, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.UnixDate, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.RubyDate, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.RFC850, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.RFC1123, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.RFC1123Z, string(bs)); err == nil {
			return tm, nil
		} else {
			return tm, newUnknownTimeFormatError(string(bs), rpos)
		}
	}
}

// Timep extracts the next time value and returns a pointer variable.
func (r *Reader) Timep() (res *time.Time, err error) {
	if v, err := r.Time(); err == nil {
		res = &v
	}
	return
}
