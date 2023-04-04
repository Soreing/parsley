package parsley

import (
	"fmt"

	"github.com/Soreing/parsley/reader"
)

type ParsleyJSONUnmarshaller interface {
	UnmarshalParsleyJSON(*reader.Reader) error
}

func Unmarshal(src []byte, dst ParsleyJSONUnmarshaller) error {
	if dst == nil {
		return fmt.Errorf("object is nil")
	}

	r := reader.NewReader(src)
	r.SkipWhiteSpace()
	return dst.UnmarshalParsleyJSON(r)
}
