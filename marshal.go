package parsley

type ParsleyJSONMarshaller interface {
	MarshalParsleyJSON(dst []byte) int
	LengthParsleyJSON() int
}

func Marshal(o ParsleyJSONMarshaller) ([]byte, error) {
	dst := make([]byte, o.LengthParsleyJSON())
	ln := o.MarshalParsleyJSON(dst)
	return dst[:ln], nil
}

func MarshalInto(dst []byte, o ParsleyJSONMarshaller) error {
	o.MarshalParsleyJSON(dst)
	return nil
}
