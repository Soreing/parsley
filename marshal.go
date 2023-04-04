package parsley

import "fmt"

type ParsleyJSONMarshaller interface {
	MarshalParsleyJSON(dst []byte) int
	LengthParsleyJSON() int
}

func Marshal(obj ParsleyJSONMarshaller) ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("object is nil")
	}

	dst := make([]byte, obj.LengthParsleyJSON())
	ln := obj.MarshalParsleyJSON(dst)
	return dst[:ln], nil
}
