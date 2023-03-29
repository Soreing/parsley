package parsley

import "github.com/Soreing/parsley/reader"

type ParsleyJSONUnmarshaller interface {
	UnmarshalParsleyJSON(*reader.Reader) error
}

func Unmarshal(src []byte, dst ParsleyJSONUnmarshaller) error {
	r := reader.NewReader(src)
	r.SkipWhiteSpace()
	return dst.UnmarshalParsleyJSON(r)
}
