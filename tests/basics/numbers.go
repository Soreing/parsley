package basics

type IntegersColl struct {
	I8Dat  int8    `json:"i8dat"`
	I8Slc  []int8  `json:"i8slc"`
	I8Ptr  *int8   `json:"i8ptr"`
	I16Dat int16   `json:"i16dat"`
	I16Slc []int16 `json:"i16slc"`
	I16Ptr *int16  `json:"i16ptr"`
	I32Dat int32   `json:"i32dat"`
	I32Slc []int32 `json:"i32slc"`
	I32Ptr *int32  `json:"i32ptr"`
	I64Dat int64   `json:"i64dat"`
	I64Slc []int64 `json:"i64slc"`
	I64Ptr *int64  `json:"i64ptr"`
	IDat   int     `json:"idat"`
	ISlc   []int   `json:"islc"`
	IPtr   *int    `json:"iptr"`
}

type UnsignedIntegersColl struct {
	UI8Dat  uint8    `json:"ui8dat"`
	UI8Slc  []uint8  `json:"ui8slc"`
	UI8Ptr  *uint8   `json:"ui8ptr"`
	UI16Dat uint16   `json:"ui16dat"`
	UI16Slc []uint16 `json:"ui16slc"`
	UI16Ptr *uint16  `json:"ui16ptr"`
	UI32Dat uint32   `json:"ui32dat"`
	UI32Slc []uint32 `json:"ui32slc"`
	UI32Ptr *uint32  `json:"ui32ptr"`
	UI64Dat uint64   `json:"ui64dat"`
	UI64Slc []uint64 `json:"ui64slc"`
	UI64Ptr *uint64  `json:"ui64ptr"`
	UIDat   uint     `json:"uidat"`
	UISlc   []uint   `json:"uislc"`
	UIPtr   *uint    `json:"uiptr"`
}

type FloatingPointColl struct {
	F32Dat float32   `json:"f32dat"`
	F32Slc []float32 `json:"f32slc"`
	F32Ptr *float32  `json:"f32ptr"`
	F64Dat float64   `json:"f64dat"`
	F64Slc []float64 `json:"f64slc"`
	F64Ptr *float64  `json:"f64ptr"`
}
