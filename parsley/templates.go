package main

import (
	"fmt"
	"strconv"
	"strings"
)

var control = []string{
	"u0000", "u0001", "u0002", "u0003", "u0004", "u0005", "u0006", "u0007",
	"u0008", "t", "n", "u000B", "u000C", "r", "u000E", "u000F",
	"u0010", "u0011", "u0012", "u0013", "u0014", "u0015", "u0016", "u0017",
	"u0018", "u0019", "u001A", "u001B", "u001C", "u001D", "u001E", "u001F",
}

// Returns the code/value used to compare types against their default values.
// If the type is unknown, empty string is returned.
func getValueCheck(fi FieldInfo) (zv string) {
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

// Returns the byte length of the default values for each known type.
// Unknown types return zero
func getDefaultValueByteLength(fi FieldInfo) (ln int) {
	if fi.Array || fi.Pointer {
		return 4
	}
	switch fi.TypeName {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64":
		return 1
	case "bool":
		return 5
	case "string":
		return 2
	case "time.Time":
		return 22
	default:
		return 0
	}
}

// Returns the code template for reading defined datatypes.
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

// Returns the code template for reading defined datatypes.
func getReaderTypeFormat(typename string) (tmpl string, unknown bool) {
	switch typename {
	case "int":
		tmpl = "r.GetInt%s()"
	case "int8":
		tmpl = "r.GetInt8%s()"
	case "int16":
		tmpl = "r.GetInt16%s()"
	case "int32":
		tmpl = "r.GetInt32%s()"
	case "int64":
		tmpl = "r.GetInt64%s()"
	case "uint":
		tmpl = "r.GetUInt%s()"
	case "uint8":
		tmpl = "r.GetUInt8%s()"
	case "uint16":
		tmpl = "r.GetUInt16%s()"
	case "uint32":
		tmpl = "r.GetUInt32%s()"
	case "uint64":
		tmpl = "r.GetUInt64%s()"
	case "float32":
		tmpl = "r.GetFloat32%s()"
	case "float64":
		tmpl = "r.GetFloat64%s()"
	case "bool":
		tmpl = "r.GetBool%s()"
	case "string":
		tmpl = "r.GetString%s()"
	case "time.Time":
		tmpl = "r.GetTime%s()"
	default:
		unknown = true
	}
	return
}

// Returns the code template for writing defined datatypes.
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

func calculateDefaultLength(fis []FieldInfo) (ln int) {
	for _, fi := range fis {
		ln += getDefaultValueByteLength(fi)
		if !fi.OmitEmpty {
			ln += len(fi.AliasEsc) + len(`"":,`)
		}
	}
	return
}

func createStructLengthBody(fis []FieldInfo) (code string) {
	subs := make([]string, len(fis))
	for i, fi := range fis {
		valCheck := getValueCheck(fi)
		modBytes := -getDefaultValueByteLength(fi)
		format, value, mod := "ln+=%s\n", "", ""

		if fn, unknown := getLengthTypeFormat(fi.TypeName); !unknown {
			if fi.Array {
				value = fmt.Sprintf(fn, "s", "o."+fi.Name)
			} else if fi.Pointer {
				value = fmt.Sprintf(fn, "", "*o."+fi.Name)
			} else {
				value = fmt.Sprintf(fn, "", "o."+fi.Name)
			}
		} else {
			if fi.Array {
				value = "(*" + fi.TypeName + ")(nil).LengthParsleyJSONSlice(o." + fi.Name + ")"
			} else {
				value = "o." + fi.Name + ".LengthParsleyJSON()"
			}
		}

		if valCheck != "" {
			format = "if o." + fi.Name + valCheck + "{\nln+=%s\n}\n"
		}
		if fi.OmitEmpty {
			modBytes += len(fi.AliasEsc) + len(`"":,`)
		}
		if modBytes > 0 {
			mod = "+" + strconv.Itoa(modBytes)
		} else if modBytes < 0 {
			mod = "-" + strconv.Itoa(-modBytes)
		}

		subs[i] = fmt.Sprintf(format, value+mod)
	}

	return strings.Join(subs, "")
}

func createUnmarshalStructBody(fis []FieldInfo) (code string) {
	subs := make([]string, len(fis))
	for i, fi := range fis {
		subs[i] = "case \"" + fi.AliasEsc + "\":\n"

		if fn, unknown := getReaderTypeFormat(fi.TypeName); !unknown {
			subs[i] += "o." + fi.Name + ", err = "
			if fi.Array {
				subs[i] += fmt.Sprintf(fn, "s") + "\n"
			} else if fi.Pointer {
				subs[i] += fmt.Sprintf(fn, "Ptr") + "\n"
			} else {
				subs[i] += fmt.Sprintf(fn, "") + "\n"
			}
		} else {
			if fi.Array {
				subs[i] += "o." + fi.Name + ", err = (*" + fi.TypeName + ")(nil).UnmarshalParsleyJSONSlice(r)\n"
			} else if fi.Pointer {
				subs[i] += "o." + fi.Name + " = &" + fi.TypeName + "{}\n" +
					"err = o." + fi.Name + ".UnmarshalParsleyJSON(r)\n"
			} else {
				subs[i] += "err = o." + fi.Name + ".UnmarshalParsleyJSON(r)\n"
			}
		}
	}

	return strings.Join(subs, "")
}

func createMarshalStructBody(fis []FieldInfo) (code string) {
	subs := make([]string, len(fis))
	skipComma, resetOffset, offsetSuffix := false, "off = 0\n", "[off:]"

	if len(fis) == 0 {
		return ""
	}

	for i, fi := range fis {
		value := "w.Raw(\",\\\"" + fi.AliasEscEsc + "\\\":\"" + offsetSuffix + ")\n"
		if fn, unknown := getWriterTypeFormat(fi.TypeName); !unknown {
			if fi.Array {
				value += fmt.Sprintf(fn, "s", "o."+fi.Name) + "\n"
			} else if fi.Pointer {
				value += fmt.Sprintf(fn, "p", "o."+fi.Name) + "\n"
			} else {
				value += fmt.Sprintf(fn, "", "o."+fi.Name) + "\n"
			}
		} else {
			if fi.Array {
				value += "(*" + fi.TypeName + ")(nil).MarshalParsleyJSONSlice(w, o." + fi.Name + ")\n"
			} else {
				value += "o." + fi.Name + ".MarshalParsleyJSON(w)\n"
			}
		}
		if !skipComma {
			value += resetOffset
		}

		format := "%s"
		valCheck := getValueCheck(fi)
		if fi.OmitEmpty && valCheck != "" {
			format = "if o." + fi.Name + valCheck + "{\n%s}\n"
		} else {
			skipComma = true
			offsetSuffix = ""
		}

		subs[i] = fmt.Sprintf(format, value)
	}

	return "off := 1\n" + strings.Join(subs, "")
}

func createDefineLengthBody(di DefineInfo) (code string) {
	if fn, unknown := getLengthTypeFormat(di.TypeName); !unknown {
		if di.Array {
			return fmt.Sprintf(fn, "s", "*o")
		} else {
			return fmt.Sprintf(fn, "", di.TypeName+"(*o)") + "\n"
		}
	} else {
		if di.Array {
			return "(*" + di.TypeName + ")(nil).LengthParsleyJSONSlice(*o)"
		} else {
			return "o.LengthParsleyJSON()"
		}
	}
}

func createUnmarshalDefineBody(di DefineInfo) (code string) {
	if fn, unknown := getReaderTypeFormat(di.TypeName); !unknown {
		if di.Array {
			return "*o, err = " + fmt.Sprintf(fn, "s") + "\n"
		} else {
			return "*(*" + di.TypeName + ")(o), err = " + fmt.Sprintf(fn, "") + "\n"
		}
	} else {
		if di.Array {
			return "*o, err = (*" + di.TypeName + ")(nil).UnmarshalParsleyJSONSlice(r)\n"
		} else {
			return "err = o.UnmarshalParsleyJSON(r)\n"
		}
	}
}

func createMarshalDefineBody(di DefineInfo) (code string) {
	if fn, unknown := getWriterTypeFormat(di.TypeName); !unknown {
		if di.Array {
			return fmt.Sprintf(fn, "s", "*o") + "\n"
		} else {
			return fmt.Sprintf(fn, "", di.TypeName+"(*o)") + "\n"
		}
	} else {
		if di.Array {
			return "(*" + di.TypeName + ")(nil).MarshalParsleyJSONSlice(w, *o)\n"
		} else {
			return "o.MarshalParsleyJSON(dst[ln:])\n"
		}
	}
}
