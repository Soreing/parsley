// Code generated by fastjson for scanning JSON strings. DO NOT EDIT.
package models

import (
    reader "github.com/Soreing/fastjson/reader"
    externals "github.com/Soreing/fastjson/tests/externals"
)

func (o *Employee)UnmarshalFastJSON(r *reader.Reader) (err error) {
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
                        err = o.Person.UnmarshalFastJSON(r)
                    case "devices":
                        o.Devices, err = (*externals.Device)(nil).UnmarshalFastJSONSlice(r)
                    case "isActive":
                        o.IsActive, err = r.GetBoolean()
                    case "rating":
                        o.Rating, err = r.GetFloat64()
                    case "lineManager":
                        o.LineManager = &Employee{}
                        err = o.LineManager.UnmarshalFastJSON(r)
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

func (o *Employee)sequenceFastJSON(r *reader.Reader, idx int) (res []Employee, err error) {
    var e Employee
    if err = e.UnmarshalFastJSON(r); err == nil {
        if !r.Next() {
            res = make([]Employee, idx+1)
            res[idx] = e
            return
        } else if res, err = o.sequenceFastJSON(r, idx + 1); err == nil {
            res[idx] = e
        }
    }
    return
}

func (o *Employee) UnmarshalFastJSONSlice(r *reader.Reader) (res []Employee, err error) {
    if err = r.OpenArray(); err == nil {
        if res, err = o.sequenceFastJSON(r, 0); err == nil {
            err = r.CloseArray()
        }
    }
    return
}

func (o *Person)UnmarshalFastJSON(r *reader.Reader) (err error) {
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

func (o *Person)sequenceFastJSON(r *reader.Reader, idx int) (res []Person, err error) {
    var e Person
    if err = e.UnmarshalFastJSON(r); err == nil {
        if !r.Next() {
            res = make([]Person, idx+1)
            res[idx] = e
            return
        } else if res, err = o.sequenceFastJSON(r, idx + 1); err == nil {
            res[idx] = e
        }
    }
    return
}

func (o *Person) UnmarshalFastJSONSlice(r *reader.Reader) (res []Person, err error) {
    if err = r.OpenArray(); err == nil {
        if res, err = o.sequenceFastJSON(r, 0); err == nil {
            err = r.CloseArray()
        }
    }
    return
}

func (o *EmployeeList)UnmarshalFastJSON(r *reader.Reader) (err error) {
    *o, err = (*Employee)(nil).UnmarshalFastJSONSlice(r)
    return
}

func (o *EmployeeList)sequenceFastJSON(r *reader.Reader, idx int) (res []EmployeeList, err error) {
    var e EmployeeList
    if err = e.UnmarshalFastJSON(r); err == nil {
        if !r.Next() {
            res = make([]EmployeeList, idx+1)
            res[idx] = e
            return
        } else if res, err = o.sequenceFastJSON(r, idx + 1); err == nil {
            res[idx] = e
        }
    }
    return
}

func (o *EmployeeList) UnmarshalFastJSONSlice(r *reader.Reader) (res []EmployeeList, err error) {
    if err = r.OpenArray(); err == nil {
        if res, err = o.sequenceFastJSON(r, 0); err == nil {
            err = r.CloseArray()
        }
    }
    return
}

