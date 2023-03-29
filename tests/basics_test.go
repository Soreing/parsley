package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Soreing/parsley"
	"github.com/Soreing/parsley/tests/basics"
)

const IntegersJSON = `{
	"i8dat": -4,
	"i8slc": [5, -85],
	"i8ptr": 2,
	"i16dat": -5,
	"i16slc": [6, -86],
	"i16ptr": 3,
	"i32dat": -6,
	"i32slc": [7, -87],
	"i32ptr": 4,
	"i64dat": -7,
	"i64slc": [8, -88],
	"i64ptr": 5,
	"idat": -8,
	"islc": [9, -89],
	"iptr": 6
}`

var I8Val, I16Val, I32Val, I64Val, IVal = int8(2), int16(3), int32(4), int64(5), int(6)
var IntegersObject = basics.IntegersColl{
	I8Dat:  -4,
	I8Slc:  []int8{5, -85},
	I8Ptr:  &I8Val,
	I16Dat: -5,
	I16Slc: []int16{6, -86},
	I16Ptr: &I16Val,
	I32Dat: -6,
	I32Slc: []int32{7, -87},
	I32Ptr: &I32Val,
	I64Dat: -7,
	I64Slc: []int64{8, -88},
	I64Ptr: &I64Val,
	IDat:   -8,
	ISlc:   []int{9, -89},
	IPtr:   &IVal,
}

func Test_UnmarshalIntegers(t *testing.T) {
	dat := []byte(IntegersJSON)
	ints := basics.IntegersColl{}

	if err := parsley.Unmarshal(dat, &ints); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if ints.I8Dat != -4 {
			t.Error("i8dat property value mismatch")
		}
		if len(ints.I8Slc) != 2 {
			t.Error("i8slc property length mismatch")
		} else {
			if ints.I8Slc[0] != 5 {
				t.Error("ints.i8slc[0] element value mismatch")
			}
			if ints.I8Slc[1] != -85 {
				t.Error("ints.i8slc[1] element value mismatch")
			}
		}
		if ints.I8Ptr == nil || *ints.I8Ptr != 2 {
			t.Error("i8ptr property value mismatch")
		}

		if ints.I16Dat != -5 {
			t.Error("i16dat property value mismatch")
		}
		if len(ints.I16Slc) != 2 {
			t.Error("i16slc property length mismatch")
		} else {
			if ints.I16Slc[0] != 6 {
				t.Error("ints.i16slc[0] element value mismatch")
			}
			if ints.I16Slc[1] != -86 {
				t.Error("ints.i16slc[1] element value mismatch")
			}
		}
		if ints.I16Ptr == nil || *ints.I16Ptr != 3 {
			t.Error("i16ptr property value mismatch")
		}

		if ints.I32Dat != -6 {
			t.Error("i32dat property value mismatch")
		}
		if len(ints.I32Slc) != 2 {
			t.Error("i32slc property length mismatch")
		} else {
			if ints.I32Slc[0] != 7 {
				t.Error("ints.i32slc[0] element value mismatch")
			}
			if ints.I32Slc[1] != -87 {
				t.Error("ints.i32slc[1] element value mismatch")
			}
		}
		if ints.I32Ptr == nil || *ints.I32Ptr != 4 {
			t.Error("i32ptr property value mismatch")
		}

		if ints.I64Dat != -7 {
			t.Error("i64dat property value mismatch")
		}
		if len(ints.I8Slc) != 2 {
			t.Error("i64slc property length mismatch")
		} else {
			if ints.I64Slc[0] != 8 {
				t.Error("ints.i64slc[0] element value mismatch")
			}
			if ints.I64Slc[1] != -88 {
				t.Error("ints.i64slc[1] element value mismatch")
			}
		}
		if ints.I64Ptr == nil || *ints.I64Ptr != 5 {
			t.Error("i64ptr property value mismatch")
		}

		if ints.IDat != -8 {
			t.Error("idat property value mismatch")
		}
		if len(ints.ISlc) != 2 {
			t.Error("islc property length mismatch")
		} else {
			if ints.ISlc[0] != 9 {
				t.Error("ints.islc[0] element value mismatch")
			}
			if ints.ISlc[1] != -89 {
				t.Error("ints.islc[1] element value mismatch")
			}
		}
		if ints.IPtr == nil || *ints.IPtr != 6 {
			t.Error("iptr property value mismatch")
		}
	}
}

func Test_MarshalIntegers(t *testing.T) {
	if buf, err := parsley.Marshal(&IntegersObject); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if jbuf, err := json.Marshal(IntegersObject); err != nil {
			t.Error("standard library unmarshal failed", err)
		} else if string(buf) != string(jbuf) {
			t.Errorf(
				"marshal result mismatch \n\tHave: %s\n\tWant: %s",
				string(buf), string(jbuf),
			)
		}
	}
}

func Test_MarshalEmptyIntegers(t *testing.T) {
	obj := basics.IntegersColl{}
	if buf, err := parsley.Marshal(&obj); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if jbuf, err := json.Marshal(obj); err != nil {
			t.Error("standard library unmarshal failed", err)
		} else if string(buf) != string(jbuf) {
			t.Errorf(
				"marshal result mismatch \n\tHave: %s\n\tWant: %s",
				string(buf), string(jbuf),
			)
		}
	}
}

const UnsignedIntegersJSON = `{
	"ui8dat": 4,
	"ui8slc": [5, 85],
	"ui8ptr": 2,
	"ui16dat": 5,
	"ui16slc": [6, 86],
	"ui16ptr": 3,
	"ui32dat": 6,
	"ui32slc": [7, 87],
	"ui32ptr": 4,
	"ui64dat": 7,
	"ui64slc": [8, 88],
	"ui64ptr": 5,
	"uidat": 8,
	"uislc": [9, 89],
	"uiptr": 6
}`

var UI8Val, UI16Val, UI32Val, UI64Val, UIVal = uint8(2), uint16(3), uint32(4), uint64(5), uint(6)
var UnsignedIntegersObject = basics.UnsignedIntegersColl{
	UI8Dat:  4,
	UI8Slc:  []uint8{5, 85},
	UI8Ptr:  &UI8Val,
	UI16Dat: 5,
	UI16Slc: []uint16{6, 86},
	UI16Ptr: &UI16Val,
	UI32Dat: 6,
	UI32Slc: []uint32{7, 87},
	UI32Ptr: &UI32Val,
	UI64Dat: 7,
	UI64Slc: []uint64{8, 88},
	UI64Ptr: &UI64Val,
	UIDat:   8,
	UISlc:   []uint{9, 89},
	UIPtr:   &UIVal,
}

func Test_UnmarshalUnsignedIntegers(t *testing.T) {
	dat := []byte(UnsignedIntegersJSON)
	uints := basics.UnsignedIntegersColl{}

	if err := parsley.Unmarshal(dat, &uints); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if uints.UI8Dat != 4 {
			t.Error("ui8dat property value mismatch")
		}
		if len(uints.UI8Slc) != 2 {
			t.Error("ui8slc property length mismatch")
		} else {
			if uints.UI8Slc[0] != 5 {
				t.Error("uints.ui8slc[0] element value mismatch")
			}
			if uints.UI8Slc[1] != 85 {
				t.Error("uints.ui8slc[1] element value mismatch")
			}
		}
		if uints.UI8Ptr == nil || *uints.UI8Ptr != 2 {
			t.Error("ui8ptr property value mismatch")
		}

		if uints.UI16Dat != 5 {
			t.Error("ui16dat property value mismatch")
		}
		if len(uints.UI16Slc) != 2 {
			t.Error("ui16slc property length mismatch")
		} else {
			if uints.UI16Slc[0] != 6 {
				t.Error("uints.ui16slc[0] element value mismatch")
			}
			if uints.UI16Slc[1] != 86 {
				t.Error("uints.ui16slc[1] element value mismatch")
			}
		}
		if uints.UI16Ptr == nil || *uints.UI16Ptr != 3 {
			t.Error("ui16ptr property value mismatch")
		}

		if uints.UI32Dat != 6 {
			t.Error("ui32dat property value mismatch")
		}
		if len(uints.UI32Slc) != 2 {
			t.Error("ui32slc property length mismatch")
		} else {
			if uints.UI32Slc[0] != 7 {
				t.Error("uints.ui32slc[0] element value mismatch")
			}
			if uints.UI32Slc[1] != 87 {
				t.Error("uints.ui32slc[1] element value mismatch")
			}
		}
		if uints.UI32Ptr == nil || *uints.UI32Ptr != 4 {
			t.Error("ui32ptr property value mismatch")
		}

		if uints.UI64Dat != 7 {
			t.Error("ui64dat property value mismatch")
		}
		if len(uints.UI8Slc) != 2 {
			t.Error("ui64slc property length mismatch")
		} else {
			if uints.UI64Slc[0] != 8 {
				t.Error("uints.ui64slc[0] element value mismatch")
			}
			if uints.UI64Slc[1] != 88 {
				t.Error("uints.ui64slc[1] element value mismatch")
			}
		}
		if uints.UI64Ptr == nil || *uints.UI64Ptr != 5 {
			t.Error("ui64ptr property value mismatch")
		}

		if uints.UIDat != 8 {
			t.Error("uidat property value mismatch")
		}
		if len(uints.UISlc) != 2 {
			t.Error("uislc property length mismatch")
		} else {
			if uints.UISlc[0] != 9 {
				t.Error("uints.uislc[0] element value mismatch")
			}
			if uints.UISlc[1] != 89 {
				t.Error("uints.uislc[1] element value mismatch")
			}
		}
		if uints.UIPtr == nil || *uints.UIPtr != 6 {
			t.Error("uiptr property value mismatch")
		}
	}
}

func Test_MarshalUnsignedIntegers(t *testing.T) {
	if buf, err := parsley.Marshal(&UnsignedIntegersObject); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if jbuf, err := json.Marshal(UnsignedIntegersObject); err != nil {
			t.Error("standard library unmarshal failed", err)
		} else if string(buf) != string(jbuf) {
			t.Errorf(
				"marshal result mismatch \n\tHave: %s\n\tWant: %s",
				string(buf), string(jbuf),
			)
		}
	}
}

func Test_MarshalEmptyUnsignedIntegers(t *testing.T) {
	obj := basics.UnsignedIntegersColl{}
	if buf, err := parsley.Marshal(&obj); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if jbuf, err := json.Marshal(obj); err != nil {
			t.Error("standard library unmarshal failed", err)
		} else if string(buf) != string(jbuf) {
			t.Errorf(
				"marshal result mismatch \n\tHave: %s\n\tWant: %s",
				string(buf), string(jbuf),
			)
		}
	}
}

const FloatingPointJSON = `{
	"f32dat": 2.56,
	"f32slc": [0.4, 1.87],
	"f32ptr": -1.05,
	"f64dat": 5.555,
	"f64slc": [53.7, -5.7],
	"f64ptr": -5
}`

var F32Val, F64Val = float32(-1.05), float64(-5)
var FloatingPointsObject = basics.FloatingPointColl{
	F32Dat: 2.56,
	F32Slc: []float32{0.4, 1.87},
	F32Ptr: &F32Val,
	F64Dat: 7,
	F64Slc: []float64{53.7, -5.7},
	F64Ptr: &F64Val,
}

func Test_UnmarshalFloatingPoints(t *testing.T) {
	dat := []byte(FloatingPointJSON)
	floats := basics.FloatingPointColl{}

	if err := parsley.Unmarshal(dat, &floats); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if floats.F32Dat != 2.56 {
			t.Error("f32dat property value mismatch")
		}
		if len(floats.F32Slc) != 2 {
			t.Error("f32slc property length mismatch")
		} else {
			if floats.F32Slc[0] != 0.4 {
				t.Error("floats.f32slc[0] element value mismatch")
			}
			if floats.F32Slc[1] != 1.87 {
				t.Error("floats.f32slc[1] element value mismatch")
			}
		}
		if floats.F32Ptr == nil || *floats.F32Ptr != -1.05 {
			t.Error("f32ptr property value mismatch")
		}

		if floats.F64Dat != 5.555 {
			t.Error("f64dat property value mismatch")
		}
		if len(floats.F32Slc) != 2 {
			t.Error("f64slc property length mismatch")
		} else {
			if floats.F64Slc[0] != 53.7 {
				t.Error("floats.f64slc[0] element value mismatch")
			}
			if floats.F64Slc[1] != -5.7 {
				t.Error("floats.f64slc[1] element value mismatch")
			}
		}
		if floats.F64Ptr == nil || *floats.F64Ptr != -5 {
			t.Error("f64ptr property value mismatch")
		}
	}
}

func Test_MarshalFloatingPoints(t *testing.T) {
	if buf, err := parsley.Marshal(&FloatingPointsObject); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if jbuf, err := json.Marshal(FloatingPointsObject); err != nil {
			t.Error("standard library unmarshal failed", err)
		} else if string(buf) != string(jbuf) {
			t.Errorf(
				"marshal result mismatch \n\tHave: %s\n\tWant: %s",
				string(buf), string(jbuf),
			)
		}
	}
}

func Test_MarshalEmptyFloatingPoints(t *testing.T) {
	obj := basics.FloatingPointColl{}
	if buf, err := parsley.Marshal(&obj); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if jbuf, err := json.Marshal(obj); err != nil {
			t.Error("standard library unmarshal failed", err)
		} else if string(buf) != string(jbuf) {
			t.Errorf(
				"marshal result mismatch \n\tHave: %s\n\tWant: %s",
				string(buf), string(jbuf),
			)
		}
	}
}

const BoooleansJSON = `{
	"bdat": true,
	"bslc": [true, false],
	"bptr": false
}`

var BVal = false
var BooleansObject = basics.BooleansColl{
	BDat: true,
	BSlc: []bool{true, false},
	BPtr: &BVal,
}

func Test_UnmarshalBooleans(t *testing.T) {
	dat := []byte(BoooleansJSON)
	bools := basics.BooleansColl{}

	if err := parsley.Unmarshal(dat, &bools); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if bools.BDat != true {
			t.Error("bdat property value mismatch")
		}
		if len(bools.BSlc) != 2 {
			t.Error("bslc property length mismatch")
		} else {
			if bools.BSlc[0] != true {
				t.Error("bools.bslc[0] element value mismatch")
			}
			if bools.BSlc[1] != false {
				t.Error("bools.bslc[1] element value mismatch")
			}
		}
		if bools.BPtr == nil || *bools.BPtr != false {
			t.Error("bptr property value mismatch")
		}
	}
}

func Test_MarshalBooleans(t *testing.T) {
	if buf, err := parsley.Marshal(&BooleansObject); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if jbuf, err := json.Marshal(BooleansObject); err != nil {
			t.Error("standard library unmarshal failed", err)
		} else if string(buf) != string(jbuf) {
			t.Errorf(
				"marshal result mismatch \n\tHave: %s\n\tWant: %s",
				string(buf), string(jbuf),
			)
		}
	}
}

func Test_MarshalEmptyBooleans(t *testing.T) {
	obj := basics.BooleansColl{}
	if buf, err := parsley.Marshal(&obj); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if jbuf, err := json.Marshal(obj); err != nil {
			t.Error("standard library unmarshal failed", err)
		} else if string(buf) != string(jbuf) {
			t.Errorf(
				"marshal result mismatch \n\tHave: %s\n\tWant: %s",
				string(buf), string(jbuf),
			)
		}
	}
}

const StringsJSON = `{
	"sdat": "John Smith",
	"sslc": ["Hello", "World"],
	"sptr": "Test",
	"tdat": "2000-01-01T00:00:00Z",
	"tslc": ["1998-04-13T10:25:00Z", "2023-03-12T20:00:50Z"],
	"tptr": "1999-06-27T20:30:30Z"
}`

var SVal, TVal = "Test", time.Now().Add(time.Second * 2564325)
var StringsObject = basics.StringsColl{
	SDat: "John Smith",
	SSlc: []string{"Hello", "World"},
	SPtr: &SVal,
	TDat: time.Now().Add(time.Second * 8137490),
	TSlc: []time.Time{
		time.Now().Add(time.Second * 20348623),
		time.Now().Add(time.Second * 9175179),
	},
	TPtr: &TVal,
}

func Test_UnmarshalStrings(t *testing.T) {
	dat := []byte(StringsJSON)
	strings := basics.StringsColl{}

	if err := parsley.Unmarshal(dat, &strings); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if strings.SDat != "John Smith" {
			t.Error("sdat property value mismatch")
		}
		if len(strings.SSlc) != 2 {
			t.Error("sslc property length mismatch")
		} else {
			if strings.SSlc[0] != "Hello" {
				t.Error("strings.sslc[0] element value mismatch")
			}
			if strings.SSlc[1] != "World" {
				t.Error("strings.sslc[1] element value mismatch")
			}
		}
		if strings.SPtr == nil || *strings.SPtr != "Test" {
			t.Error("sptr property value mismatch")
		}

		if strings.TDat.Format(time.RFC3339) != "2000-01-01T00:00:00Z" {
			t.Error("tdat property value mismatch")
		}
		if len(strings.TSlc) != 2 {
			t.Error("tslc property length mismatch")
		} else {
			if strings.TSlc[0].Format(time.RFC3339) != "1998-04-13T10:25:00Z" {
				t.Error("strings.tslc[0] element value mismatch")
			}
			if strings.TSlc[1].Format(time.RFC3339) != "2023-03-12T20:00:50Z" {
				t.Error("strings.tslc[1] element value mismatch")
			}
		}
		if strings.TPtr == nil || (*strings.TPtr).Format(time.RFC3339) != "1999-06-27T20:30:30Z" {
			t.Error("tptr property value mismatch")
		}
	}
}

func Test_MarshalStrings(t *testing.T) {
	if buf, err := parsley.Marshal(&StringsObject); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if jbuf, err := json.Marshal(StringsObject); err != nil {
			t.Error("standard library unmarshal failed", err)
		} else if string(buf) != string(jbuf) {
			t.Errorf(
				"marshal result mismatch \n\tHave: %s\n\tWant: %s",
				string(buf), string(jbuf),
			)
		}
	}
}

func Test_MarshalEmptyStrings(t *testing.T) {
	obj := basics.StringsColl{}
	if buf, err := parsley.Marshal(&obj); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if jbuf, err := json.Marshal(obj); err != nil {
			t.Error("standard library unmarshal failed", err)
		} else if string(buf) != string(jbuf) {
			t.Errorf(
				"marshal result mismatch \n\tHave: %s\n\tWant: %s",
				string(buf), string(jbuf),
			)
		}
	}
}
