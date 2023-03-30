// Code generated by parsley for scanning JSON strings. DO NOT EDIT.
package basics

import (
	reader "github.com/Soreing/parsley/reader"
	writer "github.com/Soreing/parsley/writer"
)

var _ *reader.Reader
var _ *writer.Writer

func (o *BooleansColl) UnmarshalParsleyJSON(r *reader.Reader) (err error) {
	var key []byte
	err = r.OpenObject()
	if r.GetType() != reader.TerminatorToken {
		for err == nil {
			if key, err = r.GetKey(); err == nil {
				if r.IsNull() {
					r.SkipNull()
				} else {
					switch string(key) {
					case "bdat":
						o.BDat, err = r.GetBool()
					case "bslc":
						o.BSlc, err = r.GetBools()
					case "bptr":
						o.BPtr, err = r.GetBoolPtr()
					default:
						err = r.Skip()
					}
				}
				if err == nil && !r.Next() {
					break
				}
			}
		}
	}
	if err == nil {
		err = r.CloseObject()
	}
	return
}

func (o *BooleansColl) sequenceParsleyJSON(r *reader.Reader, idx int) (res []BooleansColl, err error) {
	var e BooleansColl
	if err = e.UnmarshalParsleyJSON(r); err == nil {
		if !r.Next() {
			res = make([]BooleansColl, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequenceParsleyJSON(r, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *BooleansColl) UnmarshalParsleyJSONSlice(r *reader.Reader) (res []BooleansColl, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequenceParsleyJSON(r, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *BooleansColl) MarshalParsleyJSON(dst []byte) (ln int) {
	if o == nil {
		return writer.WriteNull(dst)
	}
	off := 1
	_ = off
	dst[0] = '{'
	ln++
	ln += copy(dst[ln:], ",\"bdat\":"[off:])
	ln += writer.WriteBool(dst[ln:], o.BDat)
	off = 0
	ln += copy(dst[ln:], ",\"bslc\":")
	ln += writer.WriteBools(dst[ln:], o.BSlc)
	ln += copy(dst[ln:], ",\"bptr\":")
	ln += writer.WriteBoolPtr(dst[ln:], o.BPtr)
	dst[ln] = '}'
	ln++
	return ln
}

func (o *BooleansColl) MarshalParsleyJSONSlice(dst []byte, slc []BooleansColl) (ln int) {
	if slc == nil {
		return writer.WriteNull(dst)
	}
	dst[0] = '['
	ln++
	if len(slc) > 0 {
		ln += slc[0].MarshalParsleyJSON(dst[1:])
		for _, o := range slc[1:] {
			dst[ln] = ','
			ln++
			ln += o.MarshalParsleyJSON(dst[ln:])
		}
	}
	dst[ln] = ']'
	return ln + 1
}

func (o *FloatingPointColl) UnmarshalParsleyJSON(r *reader.Reader) (err error) {
	var key []byte
	err = r.OpenObject()
	if r.GetType() != reader.TerminatorToken {
		for err == nil {
			if key, err = r.GetKey(); err == nil {
				if r.IsNull() {
					r.SkipNull()
				} else {
					switch string(key) {
					case "f32dat":
						o.F32Dat, err = r.GetFloat32()
					case "f32slc":
						o.F32Slc, err = r.GetFloat32s()
					case "f32ptr":
						o.F32Ptr, err = r.GetFloat32Ptr()
					case "f64dat":
						o.F64Dat, err = r.GetFloat64()
					case "f64slc":
						o.F64Slc, err = r.GetFloat64s()
					case "f64ptr":
						o.F64Ptr, err = r.GetFloat64Ptr()
					default:
						err = r.Skip()
					}
				}
				if err == nil && !r.Next() {
					break
				}
			}
		}
	}
	if err == nil {
		err = r.CloseObject()
	}
	return
}

func (o *FloatingPointColl) sequenceParsleyJSON(r *reader.Reader, idx int) (res []FloatingPointColl, err error) {
	var e FloatingPointColl
	if err = e.UnmarshalParsleyJSON(r); err == nil {
		if !r.Next() {
			res = make([]FloatingPointColl, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequenceParsleyJSON(r, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *FloatingPointColl) UnmarshalParsleyJSONSlice(r *reader.Reader) (res []FloatingPointColl, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequenceParsleyJSON(r, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *FloatingPointColl) MarshalParsleyJSON(dst []byte) (ln int) {
	if o == nil {
		return writer.WriteNull(dst)
	}
	off := 1
	_ = off
	dst[0] = '{'
	ln++
	ln += copy(dst[ln:], ",\"f32dat\":"[off:])
	ln += writer.WriteFloat32(dst[ln:], o.F32Dat)
	off = 0
	ln += copy(dst[ln:], ",\"f32slc\":")
	ln += writer.WriteFloat32s(dst[ln:], o.F32Slc)
	ln += copy(dst[ln:], ",\"f32ptr\":")
	ln += writer.WriteFloat32Ptr(dst[ln:], o.F32Ptr)
	ln += copy(dst[ln:], ",\"f64dat\":")
	ln += writer.WriteFloat64(dst[ln:], o.F64Dat)
	ln += copy(dst[ln:], ",\"f64slc\":")
	ln += writer.WriteFloat64s(dst[ln:], o.F64Slc)
	ln += copy(dst[ln:], ",\"f64ptr\":")
	ln += writer.WriteFloat64Ptr(dst[ln:], o.F64Ptr)
	dst[ln] = '}'
	ln++
	return ln
}

func (o *FloatingPointColl) MarshalParsleyJSONSlice(dst []byte, slc []FloatingPointColl) (ln int) {
	if slc == nil {
		return writer.WriteNull(dst)
	}
	dst[0] = '['
	ln++
	if len(slc) > 0 {
		ln += slc[0].MarshalParsleyJSON(dst[1:])
		for _, o := range slc[1:] {
			dst[ln] = ','
			ln++
			ln += o.MarshalParsleyJSON(dst[ln:])
		}
	}
	dst[ln] = ']'
	return ln + 1
}

func (o *IntegersColl) UnmarshalParsleyJSON(r *reader.Reader) (err error) {
	var key []byte
	err = r.OpenObject()
	if r.GetType() != reader.TerminatorToken {
		for err == nil {
			if key, err = r.GetKey(); err == nil {
				if r.IsNull() {
					r.SkipNull()
				} else {
					switch string(key) {
					case "i8dat":
						o.I8Dat, err = r.GetInt8()
					case "i8slc":
						o.I8Slc, err = r.GetInt8s()
					case "i8ptr":
						o.I8Ptr, err = r.GetInt8Ptr()
					case "i16dat":
						o.I16Dat, err = r.GetInt16()
					case "i16slc":
						o.I16Slc, err = r.GetInt16s()
					case "i16ptr":
						o.I16Ptr, err = r.GetInt16Ptr()
					case "i32dat":
						o.I32Dat, err = r.GetInt32()
					case "i32slc":
						o.I32Slc, err = r.GetInt32s()
					case "i32ptr":
						o.I32Ptr, err = r.GetInt32Ptr()
					case "i64dat":
						o.I64Dat, err = r.GetInt64()
					case "i64slc":
						o.I64Slc, err = r.GetInt64s()
					case "i64ptr":
						o.I64Ptr, err = r.GetInt64Ptr()
					case "idat":
						o.IDat, err = r.GetInt()
					case "islc":
						o.ISlc, err = r.GetInts()
					case "iptr":
						o.IPtr, err = r.GetIntPtr()
					default:
						err = r.Skip()
					}
				}
				if err == nil && !r.Next() {
					break
				}
			}
		}
	}
	if err == nil {
		err = r.CloseObject()
	}
	return
}

func (o *IntegersColl) sequenceParsleyJSON(r *reader.Reader, idx int) (res []IntegersColl, err error) {
	var e IntegersColl
	if err = e.UnmarshalParsleyJSON(r); err == nil {
		if !r.Next() {
			res = make([]IntegersColl, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequenceParsleyJSON(r, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *IntegersColl) UnmarshalParsleyJSONSlice(r *reader.Reader) (res []IntegersColl, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequenceParsleyJSON(r, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *IntegersColl) MarshalParsleyJSON(dst []byte) (ln int) {
	if o == nil {
		return writer.WriteNull(dst)
	}
	off := 1
	_ = off
	dst[0] = '{'
	ln++
	ln += copy(dst[ln:], ",\"i8dat\":"[off:])
	ln += writer.WriteInt8(dst[ln:], o.I8Dat)
	off = 0
	ln += copy(dst[ln:], ",\"i8slc\":")
	ln += writer.WriteInt8s(dst[ln:], o.I8Slc)
	ln += copy(dst[ln:], ",\"i8ptr\":")
	ln += writer.WriteInt8Ptr(dst[ln:], o.I8Ptr)
	ln += copy(dst[ln:], ",\"i16dat\":")
	ln += writer.WriteInt16(dst[ln:], o.I16Dat)
	ln += copy(dst[ln:], ",\"i16slc\":")
	ln += writer.WriteInt16s(dst[ln:], o.I16Slc)
	ln += copy(dst[ln:], ",\"i16ptr\":")
	ln += writer.WriteInt16Ptr(dst[ln:], o.I16Ptr)
	ln += copy(dst[ln:], ",\"i32dat\":")
	ln += writer.WriteInt32(dst[ln:], o.I32Dat)
	ln += copy(dst[ln:], ",\"i32slc\":")
	ln += writer.WriteInt32s(dst[ln:], o.I32Slc)
	ln += copy(dst[ln:], ",\"i32ptr\":")
	ln += writer.WriteInt32Ptr(dst[ln:], o.I32Ptr)
	ln += copy(dst[ln:], ",\"i64dat\":")
	ln += writer.WriteInt64(dst[ln:], o.I64Dat)
	ln += copy(dst[ln:], ",\"i64slc\":")
	ln += writer.WriteInt64s(dst[ln:], o.I64Slc)
	ln += copy(dst[ln:], ",\"i64ptr\":")
	ln += writer.WriteInt64Ptr(dst[ln:], o.I64Ptr)
	ln += copy(dst[ln:], ",\"idat\":")
	ln += writer.WriteInt(dst[ln:], o.IDat)
	ln += copy(dst[ln:], ",\"islc\":")
	ln += writer.WriteInts(dst[ln:], o.ISlc)
	ln += copy(dst[ln:], ",\"iptr\":")
	ln += writer.WriteIntPtr(dst[ln:], o.IPtr)
	dst[ln] = '}'
	ln++
	return ln
}

func (o *IntegersColl) MarshalParsleyJSONSlice(dst []byte, slc []IntegersColl) (ln int) {
	if slc == nil {
		return writer.WriteNull(dst)
	}
	dst[0] = '['
	ln++
	if len(slc) > 0 {
		ln += slc[0].MarshalParsleyJSON(dst[1:])
		for _, o := range slc[1:] {
			dst[ln] = ','
			ln++
			ln += o.MarshalParsleyJSON(dst[ln:])
		}
	}
	dst[ln] = ']'
	return ln + 1
}

func (o *StringsColl) UnmarshalParsleyJSON(r *reader.Reader) (err error) {
	var key []byte
	err = r.OpenObject()
	if r.GetType() != reader.TerminatorToken {
		for err == nil {
			if key, err = r.GetKey(); err == nil {
				if r.IsNull() {
					r.SkipNull()
				} else {
					switch string(key) {
					case "sdat":
						o.SDat, err = r.GetString()
					case "sslc":
						o.SSlc, err = r.GetStrings()
					case "sptr":
						o.SPtr, err = r.GetStringPtr()
					case "tdat":
						o.TDat, err = r.GetTime()
					case "tslc":
						o.TSlc, err = r.GetTimes()
					case "tptr":
						o.TPtr, err = r.GetTimePtr()
					default:
						err = r.Skip()
					}
				}
				if err == nil && !r.Next() {
					break
				}
			}
		}
	}
	if err == nil {
		err = r.CloseObject()
	}
	return
}

func (o *StringsColl) sequenceParsleyJSON(r *reader.Reader, idx int) (res []StringsColl, err error) {
	var e StringsColl
	if err = e.UnmarshalParsleyJSON(r); err == nil {
		if !r.Next() {
			res = make([]StringsColl, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequenceParsleyJSON(r, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *StringsColl) UnmarshalParsleyJSONSlice(r *reader.Reader) (res []StringsColl, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequenceParsleyJSON(r, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *StringsColl) MarshalParsleyJSON(dst []byte) (ln int) {
	if o == nil {
		return writer.WriteNull(dst)
	}
	off := 1
	_ = off
	dst[0] = '{'
	ln++
	ln += copy(dst[ln:], ",\"sdat\":"[off:])
	ln += writer.WriteString(dst[ln:], o.SDat)
	off = 0
	ln += copy(dst[ln:], ",\"sslc\":")
	ln += writer.WriteStrings(dst[ln:], o.SSlc)
	ln += copy(dst[ln:], ",\"sptr\":")
	ln += writer.WriteStringPtr(dst[ln:], o.SPtr)
	ln += copy(dst[ln:], ",\"tdat\":")
	ln += writer.WriteTime(dst[ln:], o.TDat)
	ln += copy(dst[ln:], ",\"tslc\":")
	ln += writer.WriteTimes(dst[ln:], o.TSlc)
	ln += copy(dst[ln:], ",\"tptr\":")
	ln += writer.WriteTimePtr(dst[ln:], o.TPtr)
	dst[ln] = '}'
	ln++
	return ln
}

func (o *StringsColl) MarshalParsleyJSONSlice(dst []byte, slc []StringsColl) (ln int) {
	if slc == nil {
		return writer.WriteNull(dst)
	}
	dst[0] = '['
	ln++
	if len(slc) > 0 {
		ln += slc[0].MarshalParsleyJSON(dst[1:])
		for _, o := range slc[1:] {
			dst[ln] = ','
			ln++
			ln += o.MarshalParsleyJSON(dst[ln:])
		}
	}
	dst[ln] = ']'
	return ln + 1
}

func (o *UnsignedIntegersColl) UnmarshalParsleyJSON(r *reader.Reader) (err error) {
	var key []byte
	err = r.OpenObject()
	if r.GetType() != reader.TerminatorToken {
		for err == nil {
			if key, err = r.GetKey(); err == nil {
				if r.IsNull() {
					r.SkipNull()
				} else {
					switch string(key) {
					case "ui8dat":
						o.UI8Dat, err = r.GetUInt8()
					case "ui8slc":
						o.UI8Slc, err = r.GetUInt8s()
					case "ui8ptr":
						o.UI8Ptr, err = r.GetUInt8Ptr()
					case "ui16dat":
						o.UI16Dat, err = r.GetUInt16()
					case "ui16slc":
						o.UI16Slc, err = r.GetUInt16s()
					case "ui16ptr":
						o.UI16Ptr, err = r.GetUInt16Ptr()
					case "ui32dat":
						o.UI32Dat, err = r.GetUInt32()
					case "ui32slc":
						o.UI32Slc, err = r.GetUInt32s()
					case "ui32ptr":
						o.UI32Ptr, err = r.GetUInt32Ptr()
					case "ui64dat":
						o.UI64Dat, err = r.GetUInt64()
					case "ui64slc":
						o.UI64Slc, err = r.GetUInt64s()
					case "ui64ptr":
						o.UI64Ptr, err = r.GetUInt64Ptr()
					case "uidat":
						o.UIDat, err = r.GetUInt()
					case "uislc":
						o.UISlc, err = r.GetUInts()
					case "uiptr":
						o.UIPtr, err = r.GetUIntPtr()
					default:
						err = r.Skip()
					}
				}
				if err == nil && !r.Next() {
					break
				}
			}
		}
	}
	if err == nil {
		err = r.CloseObject()
	}
	return
}

func (o *UnsignedIntegersColl) sequenceParsleyJSON(r *reader.Reader, idx int) (res []UnsignedIntegersColl, err error) {
	var e UnsignedIntegersColl
	if err = e.UnmarshalParsleyJSON(r); err == nil {
		if !r.Next() {
			res = make([]UnsignedIntegersColl, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequenceParsleyJSON(r, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *UnsignedIntegersColl) UnmarshalParsleyJSONSlice(r *reader.Reader) (res []UnsignedIntegersColl, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequenceParsleyJSON(r, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *UnsignedIntegersColl) MarshalParsleyJSON(dst []byte) (ln int) {
	if o == nil {
		return writer.WriteNull(dst)
	}
	off := 1
	_ = off
	dst[0] = '{'
	ln++
	ln += copy(dst[ln:], ",\"ui8dat\":"[off:])
	ln += writer.WriteUInt8(dst[ln:], o.UI8Dat)
	off = 0
	ln += copy(dst[ln:], ",\"ui8slc\":")
	ln += writer.WriteUInt8s(dst[ln:], o.UI8Slc)
	ln += copy(dst[ln:], ",\"ui8ptr\":")
	ln += writer.WriteUInt8Ptr(dst[ln:], o.UI8Ptr)
	ln += copy(dst[ln:], ",\"ui16dat\":")
	ln += writer.WriteUInt16(dst[ln:], o.UI16Dat)
	ln += copy(dst[ln:], ",\"ui16slc\":")
	ln += writer.WriteUInt16s(dst[ln:], o.UI16Slc)
	ln += copy(dst[ln:], ",\"ui16ptr\":")
	ln += writer.WriteUInt16Ptr(dst[ln:], o.UI16Ptr)
	ln += copy(dst[ln:], ",\"ui32dat\":")
	ln += writer.WriteUInt32(dst[ln:], o.UI32Dat)
	ln += copy(dst[ln:], ",\"ui32slc\":")
	ln += writer.WriteUInt32s(dst[ln:], o.UI32Slc)
	ln += copy(dst[ln:], ",\"ui32ptr\":")
	ln += writer.WriteUInt32Ptr(dst[ln:], o.UI32Ptr)
	ln += copy(dst[ln:], ",\"ui64dat\":")
	ln += writer.WriteUInt64(dst[ln:], o.UI64Dat)
	ln += copy(dst[ln:], ",\"ui64slc\":")
	ln += writer.WriteUInt64s(dst[ln:], o.UI64Slc)
	ln += copy(dst[ln:], ",\"ui64ptr\":")
	ln += writer.WriteUInt64Ptr(dst[ln:], o.UI64Ptr)
	ln += copy(dst[ln:], ",\"uidat\":")
	ln += writer.WriteUInt(dst[ln:], o.UIDat)
	ln += copy(dst[ln:], ",\"uislc\":")
	ln += writer.WriteUInts(dst[ln:], o.UISlc)
	ln += copy(dst[ln:], ",\"uiptr\":")
	ln += writer.WriteUIntPtr(dst[ln:], o.UIPtr)
	dst[ln] = '}'
	ln++
	return ln
}

func (o *UnsignedIntegersColl) MarshalParsleyJSONSlice(dst []byte, slc []UnsignedIntegersColl) (ln int) {
	if slc == nil {
		return writer.WriteNull(dst)
	}
	dst[0] = '['
	ln++
	if len(slc) > 0 {
		ln += slc[0].MarshalParsleyJSON(dst[1:])
		for _, o := range slc[1:] {
			dst[ln] = ','
			ln++
			ln += o.MarshalParsleyJSON(dst[ln:])
		}
	}
	dst[ln] = ']'
	return ln + 1
}
