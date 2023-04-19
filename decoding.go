package parsley

import (
	"fmt"

	"github.com/Soreing/parsley/reader"
)

// Unmarshaller describes methods for an object to unmarshal a JSON string
// as a byte array into an object, mapping the values to the fields
type Unmarshaller interface {
	DecodeObjectPJSON(r *reader.Reader, filter []Filter) error
}

// Unmarshal takes a JSON string in a byte array and reads the content of the JSON
// into an object in the parameter list. The object must implement the Unmarshaller
// interface, otherwise it returns an error. This function matches the signature
// of the standard library's Unmarshal function. To enforce objects implementing
// the Unmarshaller interface at compile time, use the Decode function instead.
func Unmarshal(data []byte, v any) error {
	if v == nil {
		return fmt.Errorf("object is nil")
	} else if val, ok := v.(Unmarshaller); !ok {
		return fmt.Errorf("object does not implement unmarshaller")
	} else {
		r := reader.NewReader(data)
		r.SkipWhiteSpace()
		return val.DecodeObjectPJSON(r, nil)
	}
}

// Decode takes a JSON string in a byte array and reads the content of the JSON
// into an object in the parameter list that implements the Unmarshaller interface.
// An optional list of config options can be provided to modify the decoder.
func Decode(data []byte, val Unmarshaller, cfgs ...Config) error {
	if val == nil {
		return fmt.Errorf("object is nil")
	}

	cfg := MergeConfigs(cfgs...)
	r := reader.NewReader(data)
	r.SkipWhiteSpace()
	return val.DecodeObjectPJSON(r, cfg.filter)
}
