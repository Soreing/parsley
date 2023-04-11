package parsley

import (
	"fmt"

	"github.com/Soreing/parsley/writer"
)

type Marshaller interface {
	EncodeObjectPJSON(w *writer.Writer, filter []Filter)
	ObjectLengthPJSON(filter []Filter) (bytes int, volatile int)
}

func Marshal(v any) ([]byte, error) {
	if v == nil {
		return nil, fmt.Errorf("object is nil")
	} else if val, ok := v.(Marshaller); !ok {
		return nil, fmt.Errorf("object does not implement parsley json interface")
	} else {
		bytes, volatile := val.ObjectLengthPJSON(nil)
		w := writer.NewWriter(bytes + volatile/10)
		val.EncodeObjectPJSON(w, nil)
		return w.Build(), nil
	}
}

func Encode(val Marshaller, cfgs ...config) ([]byte, error) {
	if val == nil {
		return nil, fmt.Errorf("object is nil")
	}

	cfg := MergeConfigs(cfgs...)
	bytes, volatile := val.ObjectLengthPJSON(cfg.filter)

	switch cfg.btype {
	case fixed:
		if cfg.bsize > volatile*4 {
			bytes += volatile * 4
		} else if cfg.bsize > 0 {
			bytes += cfg.bsize
		}
	case relative:
		if cfg.bsize > 400 {
			bytes += volatile * 4
		} else if cfg.bsize > 0 {
			bytes += volatile * cfg.bsize / 100
		}
	}

	w := writer.NewWriter(bytes)
	val.EncodeObjectPJSON(w, cfg.filter)
	return w.Build(), nil
}
