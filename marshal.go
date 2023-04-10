package parsley

import (
	"fmt"

	"github.com/Soreing/parsley/writer"
)

type ParsleyJSONMarshaller interface {
	MarshalParsleyJSON(w *writer.Writer)
	LengthParsleyJSON() int
}

func Marshal(ptr ParsleyJSONMarshaller) ([]byte, error) {
	if ptr == nil {
		return nil, fmt.Errorf("object is nil")
	} else {
		w := writer.NewWriter(ptr.LengthParsleyJSON())
		ptr.MarshalParsleyJSON(w)
		return w.Build(), nil
	}
}
