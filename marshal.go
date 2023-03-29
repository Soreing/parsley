package parsley

type ParsleyJSONMarshaller interface {
	MarshalParsleyJSON(dst []byte) int
}

func Marshal(o ParsleyJSONMarshaller) ([]byte, error) {
	// TODO: Calculate required size
	dst := make([]byte, 2000)
	ln := o.MarshalParsleyJSON(dst)
	return dst[:ln], nil
}

func MarshalInto(dst []byte, o ParsleyJSONMarshaller) error {
	o.MarshalParsleyJSON(dst)
	return nil
}
