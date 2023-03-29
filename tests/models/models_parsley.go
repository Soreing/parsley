// Code generated by parsley for scanning JSON strings. DO NOT EDIT.
package models

import (
	reader "github.com/Soreing/parsley/reader"
	externals "github.com/Soreing/parsley/tests/externals"
	writer "github.com/Soreing/parsley/writer"
)

var _ *reader.Reader
var _ *writer.Writer

func (o *Employee) UnmarshalParsleyJSON(r *reader.Reader) (err error) {
	var key []byte
	err = r.OpenObject()
	if r.GetType() != reader.TerminatorToken {
		for err == nil {
			if key, err = r.GetKey(); err == nil {
				if r.IsNull() {
					r.SkipNull()
				} else {
					switch string(key) {
					case "id":
						o.Id, err = r.GetString()
					case "person":
						err = o.Person.UnmarshalParsleyJSON(r)
					case "devices":
						o.Devices, err = (*externals.Device)(nil).UnmarshalParsleyJSONSlice(r)
					case "isActive":
						o.IsActive, err = r.GetBool()
					case "rating":
						o.Rating, err = r.GetFloat64()
					case "lineManager":
						o.LineManager = &Employee{}
						err = o.LineManager.UnmarshalParsleyJSON(r)
					case "tags":
						o.Tags, err = r.GetStrings()
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

func (o *Employee) sequenceParsleyJSON(r *reader.Reader, idx int) (res []Employee, err error) {
	var e Employee
	if err = e.UnmarshalParsleyJSON(r); err == nil {
		if !r.Next() {
			res = make([]Employee, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequenceParsleyJSON(r, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *Employee) UnmarshalParsleyJSONSlice(r *reader.Reader) (res []Employee, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequenceParsleyJSON(r, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *Employee) MarshalParsleyJSON(dst []byte) (ln int) {
	if o == nil {
		return writer.WriteNull(dst)
	}
	off := 1
	dst[0] = '{'
	ln++
	ln += copy(dst[ln:], ",\"id\":"[off:])
	ln += writer.WriteString(dst[ln:], o.Id)
	off = 0
	ln += copy(dst[ln:], ",\"person\":")
	ln += o.Person.MarshalParsleyJSON(dst[ln:])
	ln += copy(dst[ln:], ",\"devices\":")
	ln += (*externals.Device)(nil).MarshalParsleyJSONSlice(dst[ln:], o.Devices)
	ln += copy(dst[ln:], ",\"isActive\":")
	ln += writer.WriteBool(dst[ln:], o.IsActive)
	ln += copy(dst[ln:], ",\"rating\":")
	ln += writer.WriteFloat64(dst[ln:], o.Rating)
	ln += copy(dst[ln:], ",\"lineManager\":")
	ln += o.LineManager.MarshalParsleyJSON(dst[ln:])
	ln += copy(dst[ln:], ",\"tags\":")
	ln += writer.WriteStrings(dst[ln:], o.Tags)
	dst[ln] = '}'
	ln++
	return ln
}

func (o *Employee) MarshalParsleyJSONSlice(dst []byte, slc []Employee) (ln int) {
	if o == nil {
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

func (o *Person) UnmarshalParsleyJSON(r *reader.Reader) (err error) {
	var key []byte
	err = r.OpenObject()
	if r.GetType() != reader.TerminatorToken {
		for err == nil {
			if key, err = r.GetKey(); err == nil {
				if r.IsNull() {
					r.SkipNull()
				} else {
					switch string(key) {
					case "fname":
						o.Fname, err = r.GetString()
					case "lname":
						o.Lname, err = r.GetString()
					case "dob":
						o.DOB, err = r.GetTime()
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

func (o *Person) sequenceParsleyJSON(r *reader.Reader, idx int) (res []Person, err error) {
	var e Person
	if err = e.UnmarshalParsleyJSON(r); err == nil {
		if !r.Next() {
			res = make([]Person, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequenceParsleyJSON(r, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *Person) UnmarshalParsleyJSONSlice(r *reader.Reader) (res []Person, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequenceParsleyJSON(r, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *Person) MarshalParsleyJSON(dst []byte) (ln int) {
	if o == nil {
		return writer.WriteNull(dst)
	}
	off := 1
	dst[0] = '{'
	ln++
	ln += copy(dst[ln:], ",\"fname\":"[off:])
	ln += writer.WriteString(dst[ln:], o.Fname)
	off = 0
	ln += copy(dst[ln:], ",\"lname\":")
	ln += writer.WriteString(dst[ln:], o.Lname)
	ln += copy(dst[ln:], ",\"dob\":")
	ln += writer.WriteTime(dst[ln:], o.DOB)
	dst[ln] = '}'
	ln++
	return ln
}

func (o *Person) MarshalParsleyJSONSlice(dst []byte, slc []Person) (ln int) {
	if o == nil {
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

func (o *EmployeeList) UnmarshalParsleyJSON(r *reader.Reader) (err error) {
	*o, err = (*Employee)(nil).UnmarshalParsleyJSONSlice(r)
	return
}

func (o *EmployeeList) sequenceParsleyJSON(r *reader.Reader, idx int) (res []EmployeeList, err error) {
	var e EmployeeList
	if err = e.UnmarshalParsleyJSON(r); err == nil {
		if !r.Next() {
			res = make([]EmployeeList, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequenceParsleyJSON(r, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *EmployeeList) UnmarshalParsleyJSONSlice(r *reader.Reader) (res []EmployeeList, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequenceParsleyJSON(r, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *EmployeeList) MarshalParsleyJSON(dst []byte) (ln int) {
	return (*Employee)(nil).MarshalParsleyJSONSlice(dst[ln:], *o)

}

func (o *EmployeeList) MarshalParsleyJSONSlice(dst []byte, slc []EmployeeList) (ln int) {
	if slc != nil {
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
	} else {
		return writer.WriteNull(dst)
	}
}
