package controls

//parsley:json,public
type PublicField struct {
	field string
}

func (o *PublicField) GetFieldValue() string {
	return o.field
}

func (o *PublicField) SetFieldValue(str string) {
	o.field = str
}

//parsley:json
type PrivateField struct {
	field string
}

func (o *PrivateField) GetFieldValue() string {
	return o.field
}

func (o *PrivateField) SetFieldValue(str string) {
	o.field = str
}
