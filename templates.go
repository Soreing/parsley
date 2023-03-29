package parsley

import (
	"fmt"
	"strings"
)

// Returns the code/value used to compare types against their default values.
// If the type is unknown, empty string is returned.
func getZeroValue(fi FieldInfo) (zv string) {
	if fi.Array || fi.Pointer {
		return "nil"
	}
	switch fi.TypeName {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64":
		return "0"
	case "bool":
		return "false"
	case "string":
		return "\"\""
	case "time.Time":
		return "time.Time{}"
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

// Returns the code for reading the next value into a field. For unknown types,
// it is expected that the type implements the interface for unmarshalling
// func getReadFieldFunction(fi FieldInfo) (fn string, unknown bool) {
// 	fn, unknown = getReadTypeFormat(fi.TypeName)

// 	if !unknown {
// 		if fi.Array {
// 			fn = fmt.Sprintf(fn, "s")
// 		} else if fi.Pointer {
// 			fn = fmt.Sprintf(fn, "Ptr")
// 		} else {
// 			fn = fmt.Sprintf(fn, "")
// 		}
// 	} else {
// 		if fi.Array {
// 			fn = "(*" + fi.TypeName + ")(nil).UnmarshalParsleyJSONSlice(r)"
// 		} else {
// 			fn = "o." + fi.Name + ".UnmarshalParsleyJSON(r)"
// 		}
// 	}

// 	return
// }

// Returns the code template for writing defined datatypes.
func getWriterTypeFormat(typename string) (tmpl string, unknown bool) {
	switch typename {
	case "int":
		tmpl = "writer.WriteInt%s(dst[ln:], %s)"
	case "int8":
		tmpl = "writer.WriteInt8%s(dst[ln:], %s)"
	case "int16":
		tmpl = "writer.WriteInt16%s(dst[ln:], %s)"
	case "int32":
		tmpl = "writer.WriteInt32%s(dst[ln:], %s)"
	case "int64":
		tmpl = "writer.WriteInt64%s(dst[ln:], %s)"
	case "uint":
		tmpl = "writer.WriteUInt%s(dst[ln:], %s)"
	case "uint8":
		tmpl = "writer.WriteUInt8%s(dst[ln:], %s)"
	case "uint16":
		tmpl = "writer.WriteUInt16%s(dst[ln:], %s)"
	case "uint32":
		tmpl = "writer.WriteUInt32%s(dst[ln:], %s)"
	case "uint64":
		tmpl = "writer.WriteUInt64%s(dst[ln:], %s)"
	case "float32":
		tmpl = "writer.WriteFloat32%s(dst[ln:], %s)"
	case "float64":
		tmpl = "writer.WriteFloat64%s(dst[ln:], %s)"
	case "bool":
		tmpl = "writer.WriteBool%s(dst[ln:], %s)"
	case "string":
		tmpl = "writer.WriteString%s(dst[ln:], %s)"
	case "time.Time":
		tmpl = "writer.WriteTime%s(dst[ln:], %s)"
	default:
		unknown = true
	}
	return
}

// Returns the code for writing a field to a byte array. For unknown types,
// it is expected that the type implements the interface for marshalling
func getWriteFieldFunction(fi FieldInfo) (fn string) {
	fn, unknown := getWriterTypeFormat(fi.TypeName)

	if !unknown {
		if fi.Array {
			fn = fmt.Sprintf(fn, "s", "o."+fi.Name)
		} else if fi.Pointer {
			fn = fmt.Sprintf(fn, "Ptr", "o."+fi.Name)
		} else {
			fn = fmt.Sprintf(fn, "", "o."+fi.Name)
		}
	} else {
		if fi.Array {
			fn = "(*" + fi.TypeName + ")(nil).MarshalParsleyJSONSlice(dst[ln:], o." + fi.Name + ")"
		} else {
			fn = "o." + fi.Name + ".MarshalParsleyJSON(dst[ln:])"
		}
	}
	return
}

// NEW FUNCTIONS

func createUnmarshalStructBody(fis []FieldInfo) (code string) {
	subs := make([]string, len(fis))
	for i, fi := range fis {
		subs[i] = "case \"" + fi.Alias + "\":\n"

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

	for i, fi := range fis {
		value := "ln += copy(dst[ln:], \",\\\"" + fi.AliasEsc + "\\\":\"" + offsetSuffix + ")\n"
		if fn, unknown := getWriterTypeFormat(fi.TypeName); !unknown {
			if fi.Array {
				value += "ln += " + fmt.Sprintf(fn, "s", "o."+fi.Name) + "\n"
			} else if fi.Pointer {
				value += "ln += " + fmt.Sprintf(fn, "Ptr", "o."+fi.Name) + "\n"
			} else {
				value += "ln += " + fmt.Sprintf(fn, "", "o."+fi.Name) + "\n"
			}
		} else {
			if fi.Array {
				value += "ln += (*" + fi.TypeName + ")(nil).MarshalParsleyJSONSlice(dst[ln:], o." + fi.Name + ")\n"
			} else {
				value += "ln += o." + fi.Name + ".MarshalParsleyJSON(dst[ln:])\n"
			}
		}
		if !skipComma {
			value += resetOffset
		}

		format := "%s"
		zeroVal := getZeroValue(fi)
		if fi.OmitEmpty && zeroVal != "" {
			format = "if o." + fi.Name + " != " + zeroVal + "{\n%s}\n"
		} else {
			skipComma = true
			offsetSuffix = ""
		}

		subs[i] = fmt.Sprintf(format, value)
	}

	return strings.Join(subs, "")
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
			return "(*" + di.TypeName + ")(nil).MarshalParsleyJSONSlice(dst[ln:], *o)\n"
		} else {
			return "o.MarshalParsleyJSON(dst[ln:])\n"
		}
	}
}
