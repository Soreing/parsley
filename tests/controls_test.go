package main

import (
	"testing"

	"github.com/Soreing/parsley"
	"github.com/Soreing/parsley/tests/controls"
)

const EscapedFieldJSON = `{
	"soɯ\u0259 \"value\"": "1\"2\\3\/4\b5\f6\n7\r8\t9\u02e00ɯ😃"
}`

// It seems like encoding/json limits what can be in a field alias
var EscapedFieldResult = `{"soɯə \"value\"":"1\"2\\3/4\u00085\u000C6\n7\r8\t90ˠɯ😃"}`
var EscapedFieldObject = controls.EscapedField{Value: "1\"2\\3/4\b5\f6\n7\r8\t90ˠɯ😃"}

func Test_UnmarshalEscapedField(t *testing.T) {
	dat := []byte(EscapedFieldJSON)
	obj := controls.EscapedField{}
	res := "1\"2\\3/4\b5\f6\n7\r8\t9ˠ0ɯ😃"

	if err := parsley.Unmarshal(dat, &obj); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if obj.Value != res {
			t.Errorf(
				"value property value mismatch \n\tHave: %s\n\tWant: %s",
				obj.Value, res,
			)
		}
	}
}

func Test_MarshalEscapedField(t *testing.T) {
	if buf, err := parsley.Marshal(&EscapedFieldObject); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if string(buf) != EscapedFieldResult {
			t.Errorf(
				"marshal result mismatch \n\tHave: %s\n\tWant: %s",
				string(buf), EscapedFieldResult,
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

const PublicFieldJSON = `{"field":"value"}`
const PrivateFieldJSON = `{}`

func Test_FieldVisibilityDecoding(t *testing.T) {
	dat := []byte(PublicFieldJSON)
	pub := controls.PublicField{}
	prv := controls.PrivateField{}

	if err := parsley.Unmarshal(dat, &pub); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if pub.GetFieldValue() != "value" {
			t.Errorf("value property value mismatch")
		}
	}
	if err := parsley.Unmarshal(dat, &prv); err != nil {
		t.Error("unmarshal failed", err)
	} else {
		if pub.GetFieldValue() != "value" {
			t.Errorf("value property value mismatch")
		}
	}
}

func Test_FieldVisibilityEncoding(t *testing.T) {
	pub := controls.PublicField{}
	pub.SetFieldValue("value")
	prv := controls.PrivateField{}
	prv.SetFieldValue("value")

	if res, err := parsley.Marshal(&pub); err != nil {
		t.Error("marshal failed", err)
	} else {
		if string(res) != PublicFieldJSON {
			t.Errorf("public result mismatch")
		}
	}
	if res, err := parsley.Marshal(&prv); err != nil {
		t.Error("marshal failed", err)
	} else {
		if string(res) != PrivateFieldJSON {
			t.Errorf("private result mismatch")
		}
	}
}
