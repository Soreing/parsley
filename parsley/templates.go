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

// Code for finding the byte length of values
var lengthTypeFormat = map[string]string{
	"int":       "writer.Int%sLen(%s)",
	"int8":      "writer.Int8%sLen(%s)",
	"int16":     "writer.Int16%sLen(%s)",
	"int32":     "writer.Int32%sLen(%s)",
	"int64":     "writer.Int64%sLen(%s)",
	"uint":      "writer.UInt%sLen(%s)",
	"uint8":     "writer.UInt8%sLen(%s)",
	"uint16":    "writer.UInt16%sLen(%s)",
	"uint32":    "writer.UInt32%sLen(%s)",
	"uint64":    "writer.UInt64%sLen(%s)",
	"float32":   "writer.Float32%sLen(%s)",
	"float64":   "writer.Float64%sLen(%s)",
	"bool":      "writer.Bool%sLen(%s)",
	"string":    "writer.String%sLen(%s)",
	"time.Time": "writer.Time%sLen(%s)",
}

// Code for decoding basic type values.
var readerTypeFormat = map[string]string{
	"int":       "r.Int%s()",
	"int8":      "r.Int8%s()",
	"int16":     "r.Int16%s()",
	"int32":     "r.Int32%s()",
	"int64":     "r.Int64%s()",
	"uint":      "r.UInt%s()",
	"uint8":     "r.UInt8%s()",
	"uint16":    "r.UInt16%s()",
	"uint32":    "r.UInt32%s()",
	"uint64":    "r.UInt64%s()",
	"float32":   "r.Float32%s()",
	"float64":   "r.Float64%s()",
	"bool":      "r.Bool%s()",
	"string":    "r.String%s()",
	"time.Time": "r.Time%s()",
}

// Code for encoding basic type values.
var writerTypeFormat = map[string]string{
	"int":       "w.Int%s(%s)",
	"int8":      "w.Int8%s(%s)",
	"int16":     "w.Int16%s(%s)",
	"int32":     "w.Int32%s(%s)",
	"int64":     "w.Int64%s(%s)",
	"uint":      "w.UInt%s(%s)",
	"uint8":     "w.UInt8%s(%s)",
	"uint16":    "w.UInt16%s(%s)",
	"uint32":    "w.UInt32%s(%s)",
	"uint64":    "w.UInt64%s(%s)",
	"float32":   "w.Float32%s(%s)",
	"float64":   "w.Float64%s(%s)",
	"bool":      "w.Bool%s(%s)",
	"string":    "w.String%s(%s)",
	"time.Time": "w.Time%s(%s)",
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
		if tmpl, ok := lengthTypeFormat[fi.TypeName]; ok {
			if fi.Array {
				lenFn = fmt.Sprintf(tmpl, "s", "o."+fi.Name)
			} else if fi.Pointer {
				lenFn = fmt.Sprintf(tmpl, "p", "o."+fi.Name)
			} else {
				lenFn = fmt.Sprintf(tmpl, "", "o."+fi.Name)
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

		if tmpl, ok := readerTypeFormat[fi.TypeName]; ok {
			subs[i] += "o." + fi.Name + ", err = "
			if fi.Array {
				subs[i] += fmt.Sprintf(tmpl, "s") + "\n"
			} else if fi.Pointer {
				subs[i] += fmt.Sprintf(tmpl, "p") + "\n"
			} else {
				subs[i] += fmt.Sprintf(tmpl, "") + "\n"
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
		if tmpl, ok := writerTypeFormat[fi.TypeName]; ok {
			if fi.Array {
				subs[i] += fmt.Sprintf(tmpl, "s", "o."+fi.Name) + "\n"
			} else if fi.Pointer {
				subs[i] += fmt.Sprintf(tmpl, "p", "o."+fi.Name) + "\n"
			} else {
				subs[i] += fmt.Sprintf(tmpl, "", "o."+fi.Name) + "\n"
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

	if tmpl, ok := lengthTypeFormat[di.TypeName]; ok {
		if di.Array {
			return fmt.Sprintf(tmpl, "s", "*o") + vlt
		} else {
			return fmt.Sprintf(tmpl, "", di.TypeName+"(*o)") + vlt + "\n"
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
	if tmpl, ok := readerTypeFormat[di.TypeName]; ok {
		if di.Array {
			return "*o, err = " + fmt.Sprintf(tmpl, "s") + "\n"
		} else {
			return "*(*" + di.TypeName + ")(o), err = " + fmt.Sprintf(tmpl, "") + "\n"
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
	if tmpl, ok := writerTypeFormat[di.TypeName]; ok {
		if di.Array {
			return fmt.Sprintf(tmpl, "s", "*o") + "\n"
		} else {
			return fmt.Sprintf(tmpl, "", di.TypeName+"(*o)") + "\n"
		}
	} else {
		if di.Array {
			return "(*" + di.TypeName + ")(nil).EncodeSlicePJSON(w, filter, *o)\n"
		} else {
			return "o.EncodeObjectPJSON(filter, dst[ln:])\n"
		}
	}
}
