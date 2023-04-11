package main

import (
	"bytes"
	"fmt"
	"go/format"
	"reflect"
	"regexp"
	"strings"
)

type Case int

const (
	LOWER_CASE Case = iota
	CAMEL_CASE
	PASCAL_CASE
	SNAKE_CASE
	KEBAB_CASE
)

type Generator struct {
	buf         bytes.Buffer
	pkgName     string
	defaultCase Case
}

// Creates a new Generator
func NewGenerator() *Generator {
	return &Generator{
		buf:         bytes.Buffer{},
		pkgName:     "",
		defaultCase: LOWER_CASE,
	}
}

// Sets target package name
func (g *Generator) SetPackage(pkg string) {
	g.pkgName = pkg
}

// Sets default casing of the column names
func (g *Generator) SetDefaultCase(cs Case) {
	g.defaultCase = cs
}

// Writes into the generator's buffer
func (g *Generator) Printf(f string, v ...any) {
	g.buf.WriteString(fmt.Sprintf(f, v...))
}

// Gets the content of the generator's buffer
func (g *Generator) ReadAll() []byte {
	return g.buf.Bytes()
}

// Formats a given source code
func (g *Generator) Format(src []byte) ([]byte, error) {
	return format.Source(src)
}

func (g *Generator) GetRequiredPackages(
	i []Import,
	d []Define,
	s []Struct,
) (pkgs map[string]string) {
	pkgs = map[string]string{
		"parse":  "\"github.com/Soreing/parsley\"",
		"reader": "\"github.com/Soreing/parsley/reader\"",
		"writer": "\"github.com/Soreing/parsley/writer\"",
	}

	for _, e := range d {
		_, ok := pkgs[e.typ.pkg]
		if !ok && e.typ.pkg != "" && e.typ.arr {
			pkgs[e.typ.pkg] = ""
		}
	}
	for _, e := range s {
		for _, f := range e.fields {
			_, ok := pkgs[f.typ.pkg]
			if !ok && f.typ.pkg != "" && (f.typ.arr || f.typ.ptr) {
				pkgs[f.typ.pkg] = ""
			}
		}
	}
	for _, e := range i {
		if path, ok := pkgs[e.name]; ok && path == "" {
			pkgs[e.name] = e.path
		}
	}
	if path, ok := pkgs["time"]; ok && path == "\"time\"" {
		delete(pkgs, "time")
	}

	return
}

type FieldInfo struct {
	TypeName    string
	OmitEmpty   bool
	Name        string
	Alias       string
	AliasEsc    string
	AliasEscEsc string
	Array       bool
	Pointer     bool
}

func NewFieldInfo(f Field, dcs Case) (fi FieldInfo) {
	fi.Alias, fi.OmitEmpty = parseTag(strings.Trim(f.tag, "`"))
	fi.Array, fi.Pointer = f.typ.arr, f.typ.ptr
	if fi.Alias == "" {
		fi.Alias = caseString(f.name, dcs)
	}
	fi.AliasEsc = escapeJSONString(fi.Alias, "\\")
	fi.AliasEscEsc = escapeJSONString(fi.Alias, "\\\\\\")
	fi.Name, fi.TypeName = f.name, f.typ.typ
	if f.typ.pkg != "" {
		fi.TypeName = f.typ.pkg + "." + fi.TypeName
	}
	return
}

type DefineInfo struct {
	TypeName string
	Array    bool
	Pointer  bool
}

func NewDefineInfo(f Define) (di DefineInfo) {
	di.Array, di.Pointer, di.TypeName = f.typ.arr, f.typ.ptr, f.typ.typ
	if f.typ.pkg != "" {
		di.TypeName = f.typ.pkg + "." + di.TypeName
	}
	return
}

// Changes a string's casing to the given format
// The string must be in camel or pascal case
func caseString(str string, cs Case) string {
	if str == "" {
		return str
	}

	r, _ := regexp.Compile("[A-Z]?([A-Z]+|[a-z]+)")
	words := r.FindAllString(str, -1)
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}

	switch cs {
	case LOWER_CASE:
		return strings.Join(words, "")
	case CAMEL_CASE:
		for i := 1; i < len(words); i++ {
			words[i] = strings.ToUpper(words[i][:1]) + words[i][1:]
		}
		return strings.Join(words, "")

	case PASCAL_CASE:
		for i := 0; i < len(words); i++ {
			words[i] = strings.ToUpper(words[i][:1]) + words[i][1:]
		}
		return strings.Join(words, "")
	case SNAKE_CASE:
		return strings.Join(words, "_")
	case KEBAB_CASE:
		return strings.Join(words, "-")
	}

	return str
}

// Parses the struct tag for alias and omitempty
func parseTag(str string) (alias string, omit bool) {
	tag := reflect.StructTag(str).Get("json")
	if tag == "" {
		return "", false
	}

	tokens := strings.Split(tag, ",")
	alias = tokens[0]

	for _, t := range tokens {
		if t == "omitempty" {
			omit = true
		}
	}

	return
}

// Escapes a string to be a valid JSON string
// Unescapes unicode escape sequence to runes
func escapeJSONString(s string, escape string) (res string) {
	esc, utf, hex, acc := false, false, 0, 0
	for _, c := range s {
		if utf {
			// Add the next hex digit to the accumulator
			if c >= 'A' && c <= 'F' {
				acc = acc<<4 | int(c-'A')
			} else if c >= 'a' && c <= 'a' {
				acc = acc<<4 | int(c-'a')
			} else if c >= '0' && c <= '9' {
				acc = acc<<4 | int(c-'0')
			} else {
				panic("invalid UTF-8 hexadecimal digit")
			}

			// At the end of the 4 digit sequence, add the rune
			if hex == 3 {
				res += escape + string(rune(acc))
				esc, utf, hex, acc = false, false, 0, 0
			} else {
				hex++
			}
		} else if esc {
			// Check if escape is a unicode escape sequence
			if c == 'u' {
				utf = true
			} else {
				esc = false
				res += escape + "\\"
			}
		} else if c == '\\' {
			// Enter and escape sequence on "\"
			esc = true
		} else if c < 0x1F {
			// Escape control characters
			res += escape + control[c]
		} else if c == '"' {
			res += escape + "\""
		} else {
			res += string(c)
		}
	}
	return
}
