package reader

import (
	"reflect"
	"testing"
	"time"
)

func Test_ReadTime(t *testing.T) {
	tests := []struct {
		Name      string
		Time, Fmt string
		In        []byte
		Pos       int
		Err       error
	}{
		{
			Name: "Reading RFC3339 Time",
			Time: "2023-04-14T10:00:00.123456-07:30",
			Fmt:  time.RFC3339Nano,
			In:   []byte(`"2023-04-14T10:00:00.123456-07:30"`),
			Pos:  34, Err: nil,
		},
		{
			Name: "Reading RFC822 Time",
			Time: "14 Apr 23 10:00 MST",
			Fmt:  time.RFC822,
			In:   []byte(`"14 Apr 23 10:00 MST"`),
			Pos:  21, Err: nil,
		},
		{
			Name: "Reading RFC822Z Time",
			Time: "14 Apr 23 10:00 -0730",
			Fmt:  time.RFC822Z,
			In:   []byte(`"14 Apr 23 10:00 -0730"`),
			Pos:  23, Err: nil,
		},
		{
			Name: "Reading Kitchen Time",
			Time: "10:00AM",
			Fmt:  time.Kitchen,
			In:   []byte(`"10:00AM"`),
			Pos:  9, Err: nil,
		},
		{
			Name: "Reading ANSIC Time",
			Time: "Fri Apr 14 10:00:00 2023",
			Fmt:  time.ANSIC,
			In:   []byte(`"Fri Apr 14 10:00:00 2023"`),
			Pos:  26, Err: nil,
		},
		{
			Name: "Reading UnixDate Time",
			Time: "Fri Apr 14 10:00:00 MST 2023",
			Fmt:  time.UnixDate,
			In:   []byte(`"Fri Apr 14 10:00:00 MST 2023"`),
			Pos:  30, Err: nil,
		},
		{
			Name: "Reading RubyDate Time",
			Time: "Fri Apr 14 10:00:00 -0730 2023",
			Fmt:  time.RubyDate,
			In:   []byte(`"Fri Apr 14 10:00:00 -0730 2023"`),
			Pos:  32, Err: nil,
		},
		{
			Name: "Reading RFC850 Time",
			Time: "Friday, 14-Apr-23 10:00:00 MST",
			Fmt:  time.RFC850,
			In:   []byte(`"Friday, 14-Apr-23 10:00:00 MST"`),
			Pos:  32, Err: nil,
		},
		{
			Name: "Reading RFC1123 Time",
			Time: "Fri, 14 Apr 2023 10:00:00 MST",
			Fmt:  time.RFC1123,
			In:   []byte(`"Fri, 14 Apr 2023 10:00:00 MST"`),
			Pos:  31, Err: nil,
		},
		{
			Name: "Reading RFC1123Z Time",
			Time: "Fri, 14 Apr 2023 10:00:00 -0730",
			Fmt:  time.RFC1123Z,
			In:   []byte(`"Fri, 14 Apr 2023 10:00:00 -0730"`),
			Pos:  33, Err: nil,
		},
		{
			Name: "Reading Unknown Time Format",
			Time: "Friday the 14th, April 2023, 10:00:00 -07:30",
			Fmt:  time.Kitchen,
			In:   []byte(`"Friday the 14th, April 2023, 10:00:00 -07:30"`),
			Pos:  46, Err: NewUnknownTimeFormatError("Friday the 14th, April 2023, 10:00:00 -07:30", 0),
		},
		{
			Name: "Reading Nothing",
			Time: "",
			Fmt:  time.Kitchen,
			In:   []byte(`""`),
			Pos:  2, Err: NewUnknownTimeFormatError("", 0),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.Time()
			tm, _ := time.Parse(test.Fmt, test.Time)
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if !tm.Equal(res) {
				t.Errorf("got result %v, want %v", res, tm)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_ReadTimes(t *testing.T) {
	t1, _ := time.Parse(time.RFC3339Nano, "2023-04-14T10:00:00Z")
	t2, _ := time.Parse(time.RFC3339Nano, "2023-05-15T15:25:45Z")

	tests := []struct {
		Name string
		In   []byte
		Out  []time.Time
		Pos  int
		Err  error
	}{
		{
			Name: "Empty Slice",
			In:   []byte(`[]`),
			Out:  []time.Time{},
			Pos:  2, Err: nil,
		},
		{
			Name: "Slice With One Element",
			In:   []byte(`["2023-04-14T10:00:00Z"]`),
			Out:  []time.Time{t1},
			Pos:  24, Err: nil,
		},
		{
			Name: "Slice With Multiple Elements",
			In:   []byte(`["2023-04-14T10:00:00Z","2023-05-15T15:25:45Z"]`),
			Out:  []time.Time{t1, t2},
			Pos:  47, Err: nil,
		},
		{
			Name: "Slice With Whitespaces",
			In:   []byte(`[ "2023-04-14T10:00:00Z" , "2023-05-15T15:25:45Z" ]`),
			Out:  []time.Time{t1, t2},
			Pos:  51, Err: nil,
		},
		{
			Name: "Missing Opening Bracket",
			In:   []byte(`"2023-04-14T10:00:00Z","2023-05-15T15:25:45Z"]`),
			Out:  nil,
			Pos:  0, Err: NewInvalidCharacterError('"', 0),
		},
		{
			Name: "Missing Closing Bracket",
			In:   []byte(`["2023-04-14T10:00:00Z","2023-05-15T15:25:45Z"`),
			Out:  []time.Time{t1, t2},
			Pos:  46, Err: NewEndOfFileError(),
		},
		{
			Name: "Incomplete Slice",
			In:   []byte(`["2023-04-14T10:00:00Z","2023-05-15T15:25:45Z",`),
			Out:  nil,
			Pos:  47, Err: NewEndOfFileError(),
		},
		{
			Name: "Missing Comma Between Elements",
			In:   []byte(`["2023-04-14T10:00:00Z""2023-05-15T15:25:45Z"]`),
			Out:  []time.Time{t1},
			Pos:  23, Err: NewInvalidCharacterError('"', 23),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.Times()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if !reflect.DeepEqual(res, test.Out) {
				t.Errorf("got result %v, want %v", res, test.Out)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}
