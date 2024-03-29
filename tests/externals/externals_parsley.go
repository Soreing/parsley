// Code generated by parsley for scanning JSON strings. DO NOT EDIT.
package externals

import (
	parse "github.com/Soreing/parsley"
	reader "github.com/Soreing/parsley/reader"
)

var _ *reader.Reader

func (o *Device) DecodeObjectPJSON(r *reader.Reader, filter []parse.Filter) (err error) {
	c := [2]bool{}
	f := [2][]parse.Filter{}
	if filter == nil {
		for i := range c {
			c[i] = true
		}
	} else {
		for i := range filter {
			k := filter[i].Field
			if k == "name" {
				c[0] = true
			} else if k == "type" {
				c[1], f[1] = true, filter[i].Filter
			}
		}
	}
	var key []byte
	_ = key
	err = r.OpenObject()
	if r.Token() != reader.TerminatorToken {
		for err == nil {
			if key, err = r.Key(); err == nil {
				if r.IsNull() {
					r.SkipNull()
				} else {
					if string(key) == "name" && c[0] {
						o.Name, err = r.String()
					} else if string(key) == "type" && c[1] {
						err = o.Type.DecodeObjectPJSON(r, f[1])
					} else {
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

func (o *Device) sequencePJSON(r *reader.Reader, filter []parse.Filter, idx int) (res []Device, err error) {
	var e Device
	if err = e.DecodeObjectPJSON(r, filter); err == nil {
		if !r.Next() {
			res = make([]Device, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequencePJSON(r, filter, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *Device) DecodeSlicePJSON(r *reader.Reader, filter []parse.Filter) (res []Device, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == reader.TerminatorToken {
			res = []Device{}
			err = r.CloseArray()
		} else if res, err = o.sequencePJSON(r, filter, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *DeviceType) DecodeObjectPJSON(r *reader.Reader, filter []parse.Filter) (err error) {
	*(*int)(o), err = r.Int()
	return
}

func (o *DeviceType) sequencePJSON(r *reader.Reader, filter []parse.Filter, idx int) (res []DeviceType, err error) {
	var e DeviceType
	if err = e.DecodeObjectPJSON(r, filter); err == nil {
		if !r.Next() {
			res = make([]DeviceType, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequencePJSON(r, filter, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *DeviceType) DecodeSlicePJSON(r *reader.Reader, filter []parse.Filter) (res []DeviceType, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == reader.TerminatorToken {
			res = []DeviceType{}
			err = r.CloseArray()
		} else if res, err = o.sequencePJSON(r, filter, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}
