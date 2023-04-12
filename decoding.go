package parsley

import (
	"fmt"

	"github.com/Soreing/parsley/reader"
)

type Unmarshaller interface {
	DecodeObjectPJSON(r *reader.Reader, filter []Filter) error
}

func Unmarshal(data []byte, v any) error {
	if v == nil {
		return fmt.Errorf("object is nil")
	} else if val, ok := v.(Unmarshaller); !ok {
		return fmt.Errorf("object does not implement parsley json interface")
	} else {
		r := reader.NewReader(data)
		r.SkipWhiteSpace()
		return val.DecodeObjectPJSON(r, nil)
	}
}

func Decode(data []byte, val Unmarshaller, cfgs ...Config) error {
	if val == nil {
		return fmt.Errorf("object is nil")
	}

	cfg := MergeConfigs(cfgs...)
	r := reader.NewReader(data)
	r.SkipWhiteSpace()
	return val.DecodeObjectPJSON(r, cfg.filter)
}
