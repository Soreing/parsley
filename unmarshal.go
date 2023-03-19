package fastjson

import "github.com/Soreing/fastjson/reader"

type FastJSONUnmarshaller interface {
	UnmarshalFastJSON(*reader.Reader) error
}

func Unmarshal(src []byte, dst FastJSONUnmarshaller) error {
	return dst.UnmarshalFastJSON(reader.NewReader(src))
}
