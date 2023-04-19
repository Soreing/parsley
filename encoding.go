package parsley

import (
	"fmt"

	"github.com/Soreing/parsley/writer"
)

// Marshaller describes methods for an object to marshal into a JSON byte array.
type Marshaller interface {
	EncodeObjectPJSON(w *writer.Writer, filter []Filter)
	ObjectLengthPJSON(filter []Filter) (bytes int, volatile int)
}

// Marshal takes any object as a parameter and encodes it into a JSON byte array.
// The object must implement the Marshaller interface, otherwise it returns an
// error. This function matches the signature of the standard library's Marshal
// function. To enforce objects implementing the Marshaller interface at compile
// time, use the Encode function instead.
//
// Marshal does not allow for configuration. Filtering is disabled and a 10%
// relative buffer size is used by default.
func Marshal(v any) ([]byte, error) {
	if v == nil {
		return nil, fmt.Errorf("object is nil")
	} else if val, ok := v.(Marshaller); !ok {
		return nil, fmt.Errorf("object does not implement marshaller")
	} else {
		bytes, volatile := val.ObjectLengthPJSON(nil)
		w := writer.NewWriter(bytes + volatile/10)
		val.EncodeObjectPJSON(w, nil)
		return w.Build(), nil
	}
}

// Encode takes an object that implements the Marshaller interface and encodes
// the object into a JSON byte array. An optional list of config options can
// be provided to modify the encoder.
func Encode(val Marshaller, cfgs ...Config) ([]byte, error) {
	if val == nil {
		return nil, fmt.Errorf("object is nil")
	}

	cfg := MergeConfigs(cfgs...)
	bytes, volatile := val.ObjectLengthPJSON(cfg.filter)

	switch cfg.btype {
	case fixed:
		if cfg.bsize > volatile*5 {
			bytes += volatile * 5
		} else if cfg.bsize > 0 {
			bytes += cfg.bsize
		}
	case relative:
		if cfg.bsize > 500 {
			bytes += volatile * 5
		} else if cfg.bsize > 0 {
			bytes += volatile * cfg.bsize / 100
		}
	}

	w := writer.NewWriter(bytes)
	val.EncodeObjectPJSON(w, cfg.filter)
	return w.Build(), nil
}
