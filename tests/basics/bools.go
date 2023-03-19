package basics

type BooleansColl struct {
	BDat bool   `json:"bdat"`
	BSlc []bool `json:"bslc"`
	BPtr *bool  `json:"bptr"`
}
