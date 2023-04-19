package main

import (
	"fmt"
	"strconv"
	"strings"
)

// List of escape sequences for control characters.
var control = []string{
	"u0000", "u0001", "u0002", "u0003", "u0004", "u0005", "u0006", "u0007",
	"u0008", "t", "n", "u000B", "u000C", "r", "u000E", "u000F",
	"u0010", "u0011", "u0012", "u0013", "u0014", "u0015", "u0016", "u0017",
	"u0018", "u0019", "u001A", "u001B", "u001C", "u001D", "u001E", "u001F",
}

// isBasicType checks if the datatype is a basic type supported by the library.
func isBasicType(typename string) (basic bool) {
	switch typename {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "bool", "string", "time.Time":
		return true
	}
	return false
}

// isVolatile checks if the datatype can have volatile space.
func isVolatile(typename string) (volatile bool) {
	switch typename {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "bool", "time.Time":
		return false
	}
	return true
}

// getValueCheck returns the code used to compare values against their
// default values. If the type is unknown, empty string is returned.
func getValueCheck(fi fieldInfo) (zv string) {
	if fi.Array || fi.Pointer {
		return " != nil"
	}
	switch fi.TypeName {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64":
		return " != 0"
	case "bool":
		return " != false"
	case "string":
		return " != \"\""
	case "time.Time":
		return ".IsZero() != true"
	default:
		return ""
	}
}

// getLengthTypeFormat returns the code for finding the byte length of values.
func getLengthTypeFormat(typename string) (tmpl string, unknown bool) {
	switch typename {
	case "int":
		tmpl = "writer.Int%sLen(%s)"
	case "int8":
		tmpl = "writer.Int8%sLen(%s)"
	case "int16":
		tmpl = "writer.Int16%sLen(%s)"
	case "int32":
		tmpl = "writer.Int32%sLen(%s)"
	case "int64":
		tmpl = "writer.Int64%sLen(%s)"
	case "uint":
		tmpl = "writer.UInt%sLen(%s)"
	case "uint8":
		tmpl = "writer.UInt8%sLen(%s)"
	case "uint16":
		tmpl = "writer.UInt16%sLen(%s)"
	case "uint32":
		tmpl = "writer.UInt32%sLen(%s)"
	case "uint64":
		tmpl = "writer.UInt64%sLen(%s)"
	case "float32":
		tmpl = "writer.Float32%sLen(%s)"
	case "float64":
		tmpl = "writer.Float64%sLen(%s)"
	case "bool":
		tmpl = "writer.Bool%sLen(%s)"
	case "string":
		tmpl = "writer.String%sLen(%s)"
	case "time.Time":
		tmpl = "writer.Time%sLen(%s)"
	default:
		unknown = true
	}
	return
}

// getReaderTypeFormat returns the code for decoding basic type values.
func getReaderTypeFormat(typename string) (tmpl string, unknown bool) {
	switch typename {
	case "int":
		tmpl = "r.Int%s()"
	case "int8":
		tmpl = "r.Int8%s()"
	case "int16":
		tmpl = "r.Int16%s()"
	case "int32":
		tmpl = "r.Int32%s()"
	case "int64":
		tmpl = "r.Int64%s()"
	case "uint":
		tmpl = "r.UInt%s()"
	case "uint8":
		tmpl = "r.UInt8%s()"
	case "uint16":
		tmpl = "r.UInt16%s()"
	case "uint32":
		tmpl = "r.UInt32%s()"
	case "uint64":
		tmpl = "r.UInt64%s()"
	case "float32":
		tmpl = "r.Float32%s()"
	case "float64":
		tmpl = "r.Float64%s()"
	case "bool":
		tmpl = "r.Bool%s()"
	case "string":
		tmpl = "r.String%s()"
	case "time.Time":
		tmpl = "r.Time%s()"
	default:
		unknown = true
	}
	return
}

// getWriterTypeFormat returns the code for encoding basic type values.
func getWriterTypeFormat(typename string) (tmpl string, unknown bool) {
	switch typename {
	case "int":
		tmpl = "w.Int%s(%s)"
	case "int8":
		tmpl = "w.Int8%s(%s)"
	case "int16":
		tmpl = "w.Int16%s(%s)"
	case "int32":
		tmpl = "w.Int32%s(%s)"
	case "int64":
		tmpl = "w.Int64%s(%s)"
	case "uint":
		tmpl = "w.UInt%s(%s)"
	case "uint8":
		tmpl = "w.UInt8%s(%s)"
	case "uint16":
		tmpl = "w.UInt16%s(%s)"
	case "uint32":
		tmpl = "w.UInt32%s(%s)"
	case "uint64":
		tmpl = "w.UInt64%s(%s)"
	case "float32":
		tmpl = "w.Float32%s(%s)"
	case "float64":
		tmpl = "w.Float64%s(%s)"
	case "bool":
		tmpl = "w.Bool%s(%s)"
	case "string":
		tmpl = "w.String%s(%s)"
	case "time.Time":
		tmpl = "w.Time%s(%s)"
	default:
		unknown = true
	}
	return
}

// createFilterHeader creates a filter header for encoding/decoding functions
func createFilterHeader(fis []fieldInfo) (code string) {
	if len(fis) == 0 {
		return ""
	}
	lns := strconv.Itoa(len(fis))
	subf := 0
	for _, fi := range fis {
		if !isBasicType(fi.TypeName) {
			subf++
		}
	}

	code += "c := [" + lns + "]bool{}\n"
	if subf > 0 {
		code += "f := [" + lns + "][]parse.Filter{}\n"
	}
	code += "if filter == nil {\n"
	code += "    for i := range c {\n"
	code += "        c[i]=true\n"
	code += "    }\n"
	code += "} else {\n"
	code += "    for i := range filter {\n"
	code += "        k := filter[i].Field\n"

	nest := ""
	for i, fi := range fis {
		is := strconv.Itoa(i)
		nest += "} else if k == \"" + fi.AliasEsc + "\" {\n"
		if isBasicType(fi.TypeName) {
			nest += "c[" + is + "] = true\n"
		} else {
			nest += "c[" + is + "], f[" + is + "] = true, filter[i].Filter\n"
		}
	}

	code += strings.TrimPrefix(nest, "} else ")
	code += "        }\n"
	code += "    }\n"
	code += "}\n"
	return
}

func createObjectLengthBody(fis []fieldInfo) (code string) {
	subs := make([]string, len(fis))
	for i, fi := range fis {
		is := strconv.Itoa(i)
		valCheck := getValueCheck(fi)
		volatile := isVolatile(fi.TypeName)
		fieldLen := strconv.Itoa(len(fi.AliasEsc) + len(`"":,`))

		if fi.OmitEmpty && valCheck != "" {
			subs[i] += "if c[" + is + "] && o." + fi.Name + valCheck + "{\n"
		} else {
			subs[i] += "if c[" + is + "] {\n"
		}

		lenFn := ""
		if fn, unknown := getLengthTypeFormat(fi.TypeName); !unknown {
			if fi.Array {
				lenFn = fmt.Sprintf(fn, "s", "o."+fi.Name)
			} else if fi.Pointer {
				lenFn = fmt.Sprintf(fn, "p", "o."+fi.Name)
			} else {
				lenFn = fmt.Sprintf(fn, "", "o."+fi.Name)
			}
		} else {
			if fi.Array {
				lenFn = "(*" + fi.TypeName + ")(nil).SliceLengthPJSON(f[" + is + "], o." + fi.Name + ")"
			} else {
				lenFn = "o." + fi.Name + ".ObjectLengthPJSON(f[" + is + "])"
			}
		}

		if volatile {
			subs[i] += "b, v := " + lenFn + "\n"
			subs[i] += "bytes, volatile = bytes+b+" + fieldLen + ", volatile+v\n"
		} else {
			subs[i] += "bytes += " + lenFn + " + " + fieldLen + "\n"
		}
		subs[i] += "}\n"
	}

	return strings.Join(subs, "")
}

func createDecodeObjectBody(fis []fieldInfo) (code string) {
	if len(fis) == 0 {
		return "err = r.Skip()"
	}

	subs := make([]string, len(fis))
	for i, fi := range fis {
		is := strconv.Itoa(i)
		subs[i] = "} else if string(key) == \"" + fi.AliasEsc + "\" && c[" + is + "] {\n"

		if fn, unknown := getReaderTypeFormat(fi.TypeName); !unknown {
			subs[i] += "o." + fi.Name + ", err = "
			if fi.Array {
				subs[i] += fmt.Sprintf(fn, "s") + "\n"
			} else if fi.Pointer {
				subs[i] += fmt.Sprintf(fn, "p") + "\n"
			} else {
				subs[i] += fmt.Sprintf(fn, "") + "\n"
			}
		} else {
			if fi.Array {
				subs[i] += "o." + fi.Name + ", err = (*" + fi.TypeName + ")(nil).DecodeSlicePJSON(r, f[" + is + "])\n"
			} else if fi.Pointer {
				subs[i] += "o." + fi.Name + " = &" + fi.TypeName + "{}\n" +
					"err = o." + fi.Name + ".DecodeObjectPJSON(r, f[" + is + "])\n"
			} else {
				subs[i] += "err = o." + fi.Name + ".DecodeObjectPJSON(r, f[" + is + "])\n"
			}
		}
	}
	subs = append(subs, "} else { \nerr = r.Skip()\n}")
	subs[0] = strings.TrimPrefix(subs[0], "} else ")
	return strings.Join(subs, "")
}

func createEncodeObjectBody(fis []fieldInfo) (code string) {
	if len(fis) == 0 {
		return ""
	}

	subs := make([]string, len(fis))
	for i, fi := range fis {
		is := strconv.Itoa(i)

		valCheck := getValueCheck(fi)
		if fi.OmitEmpty && valCheck != "" {
			subs[i] += "if c[" + is + "] && o." + fi.Name + valCheck + "{\n"
		} else {
			subs[i] += "if c[" + is + "] {\n"
		}

		subs[i] += "w.Raw(\",\\\"" + fi.AliasEscEsc + "\\\":\"[off:])\n"
		if fn, unknown := getWriterTypeFormat(fi.TypeName); !unknown {
			if fi.Array {
				subs[i] += fmt.Sprintf(fn, "s", "o."+fi.Name) + "\n"
			} else if fi.Pointer {
				subs[i] += fmt.Sprintf(fn, "p", "o."+fi.Name) + "\n"
			} else {
				subs[i] += fmt.Sprintf(fn, "", "o."+fi.Name) + "\n"
			}
		} else {
			if fi.Array {
				subs[i] += "(*" + fi.TypeName + ")(nil).EncodeSlicePJSON(w, f[" + is + "], o." + fi.Name + ")\n"
			} else {
				subs[i] += "o." + fi.Name + ".EncodeObjectPJSON(w, f[" + is + "])\n"
			}
		}
		subs[i] += "off = 0\n"
		subs[i] += "}\n"
	}

	return "off := 1\n" + strings.Join(subs, "")
}

func createDefineLengthBody(di defineInfo) (code string) {
	vlt := ""
	if !isVolatile(di.TypeName) {
		vlt = ", 0"
	}

	if fn, unknown := getLengthTypeFormat(di.TypeName); !unknown {
		if di.Array {
			return fmt.Sprintf(fn, "s", "*o") + vlt
		} else {
			return fmt.Sprintf(fn, "", di.TypeName+"(*o)") + vlt + "\n"
		}
	} else {
		if di.Array {
			return "(*" + di.TypeName + ")(nil).SliceLengthPJSON(filter, *o)"
		} else {
			return "o.ObjectLengthPJSON(filter)"
		}
	}
}

func createDecodeDefineBody(di defineInfo) (code string) {
	if fn, unknown := getReaderTypeFormat(di.TypeName); !unknown {
		if di.Array {
			return "*o, err = " + fmt.Sprintf(fn, "s") + "\n"
		} else {
			return "*(*" + di.TypeName + ")(o), err = " + fmt.Sprintf(fn, "") + "\n"
		}
	} else {
		if di.Array {
			return "*o, err = (*" + di.TypeName + ")(nil).DecodeSlicePJSON(r, filter)\n"
		} else {
			return "err = o.DecodeObjectPJSON(r, filter)\n"
		}
	}
}

func createEncodeDefineBody(di defineInfo) (code string) {
	if fn, unknown := getWriterTypeFormat(di.TypeName); !unknown {
		if di.Array {
			return fmt.Sprintf(fn, "s", "*o") + "\n"
		} else {
			return fmt.Sprintf(fn, "", di.TypeName+"(*o)") + "\n"
		}
	} else {
		if di.Array {
			return "(*" + di.TypeName + ")(nil).EncodeSlicePJSON(w, filter, *o)\n"
		} else {
			return "o.EncodeObjectPJSON(filter, dst[ln:])\n"
		}
	}
}
