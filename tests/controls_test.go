package main

import (
	"encoding/json"
	"testing"

	"github.com/Soreing/parsley"
	"github.com/Soreing/parsley/tests/controls"
)

const EscapedFieldJSON = `{
	"soɯ\u0259 \"value\"": "1\"2\\3\/4\b5\f6\n7\r8\t9\ufefa0ɯ😃"
}`

var EscapedFieldObject = controls.EscapedField{
	Value: "1\"2\\3/4\b5\f6\n7\r8\t9ﻺ0ɯ😃",
}

func Test_UnmarshalEscapedField(t *testing.T) {
	dat := []byte(EscapedFieldJSON)
	obj := controls.EscapedField{}

	if err := parsley.Unmarshal(dat, &obj); err != nil {
		t.Error("unmarshal failed", err)
	} else {
	}
}

func Test_MarshalEscapedField(t *testing.T) {
	if buf, err := parsley.Marshal(&EscapedFieldObject); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if jbuf, err := json.Marshal(EscapedFieldObject); err != nil {
			t.Error("standard library unmarshal failed", err)
		} else if string(buf) != string(jbuf) {
			t.Errorf(
				"marshal result mismatch \n\tHave: %s\n\tWant: %s",
				string(buf), string(jbuf),
			)
		}
	}
}

const WhitespaceJSON = ` {
"key1":"value1"	,
	"key2"	: 
"value2",
	"slice": 
	[ "12"
	, 1,	true ,
null]

} `

func Test_Whitespaces(t *testing.T) {
	dat := []byte(WhitespaceJSON)
	obj := controls.EmptyObject{}

	if err := parsley.Unmarshal(dat, &obj); err != nil {
		t.Error("unmarshal failed", err)
	}
}
