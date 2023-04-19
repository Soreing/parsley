package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type import_ struct {
	name string
	path string
}

type type_ struct {
	arr bool
	ptr bool
	pkg string
	typ string
}

type struct_ struct {
	name   string
	fields []field_
}

type field_ struct {
	name string
	tag  string
	typ  type_
}

type define_ struct {
	name string
	typ  type_
}

func newImport(at *ast.ImportSpec) import_ {
	n, p := "", at.Path.Value
	if at.Name != nil {
		n = at.Name.Name
	} else {
		toks := strings.Split(strings.Trim(at.Path.Value, "\""), "/")
		n = toks[len(toks)-1]
	}
	return import_{
		name: n,
		path: p,
	}
}

func getType(e ast.Expr) (t type_, err error) {
	if arr, ok := e.(*ast.ArrayType); ok {
		t.arr = true
		e = arr.Elt
	} else if ptr, ok := e.(*ast.StarExpr); ok {
		t.ptr = true
		e = ptr.X
	}

	switch st := e.(type) {
	case *ast.Ident:
		t.typ = st.Name
	case *ast.SelectorExpr:
		t.typ = st.Sel.Name
		if pkg, ok := st.X.(*ast.Ident); ok {
			t.pkg = pkg.Name
		} else {
			err = errors.New("unsuported type")
		}
	default:
		err = errors.New("unsupported type")
	}
	return
}

func newField(f *ast.Field) (field_, error) {
	if len(f.Names) == 0 || f.Names[0].Name == "" {
		return field_{}, errors.New("field has no name")
	}

	ts, err := getType(f.Type)
	if err != nil {
		return field_{}, err
	}

	tag := ""
	if f.Tag != nil {
		tag = f.Tag.Value
	}
	return field_{
		name: f.Names[0].Name,
		tag:  tag,
		typ:  ts,
	}, nil
}

func newStruct(name string, public bool, st *ast.StructType) struct_ {
	flds := []field_{}
	for _, f := range st.Fields.List {
		if fl, err := newField(f); err == nil {
			if ast.IsExported(fl.name) || public {
				flds = append(flds, fl)
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	return struct_{
		name:   name,
		fields: flds,
	}
}

type parser_ struct {
	PkgDir    string
	PkgName   string
	Structs   []struct_
	Defines   []define_
	Imports   []import_
	AllTypes  bool
	AllPublic bool
}

type visitor_ struct {
	*parser_

	skip bool
	json bool
	publ bool
}

const (
	optionTag = "parsley:"
	jsonTag   = "json"
	skipTag   = "skip"
	publicTag = "public"
)

func (v *visitor_) handleComment(comments *ast.CommentGroup) {
	if comments == nil {
		return
	}

	for _, c := range comments.List {
		comment := c.Text
		if len(comment) < 3 {
			return
		}

		switch comment[1] {
		case '/':
			comment = comment[2:]
		case '*':
			comment = comment[2 : len(comment)-2]
		}

		for _, comment := range strings.Split(comment, "\n") {
			comment = strings.TrimSpace(comment)
			if strings.HasPrefix(comment, optionTag) {
				opts := strings.TrimPrefix(comment, optionTag)
				toks := strings.Split(opts, ",")
				for _, e := range toks {
					switch e {
					case jsonTag:
						v.json = true
					case skipTag:
						v.skip = true
					case publicTag:
						v.publ = true
					}
				}
			}
		}
	}
}

func (v *visitor_) Visit(n ast.Node) (w ast.Visitor) {
	switch n := n.(type) {
	case *ast.Package:
		return v

	case *ast.ImportSpec:
		imp := newImport(n)
		for _, e := range v.Imports {
			if e.name == imp.name {
				if e.path != imp.path {
					panic("conflicting import names")
				} else {
					return v
				}
			}
		}
		v.Imports = append(v.Imports, newImport(n))

	case *ast.GenDecl:
		return v

	case *ast.File:
		v.PkgName = n.Name.String()
		return v

	case *ast.CommentGroup:
		v.handleComment(n)
		return v

	case *ast.TypeSpec:
		if !v.skip && (v.json || v.AllTypes) {
			name := n.Name.String()
			switch t := n.Type.(type) {
			case *ast.StructType:
				st := newStruct(name, v.publ || v.AllPublic, t)
				v.Structs = append(v.Structs, st)
			case *ast.Ident, *ast.SelectorExpr, *ast.ArrayType:
				if ts, err := getType(n.Type); err == nil {
					v.Defines = append(v.Defines, define_{
						name: name,
						typ:  ts,
					})
				}
			}
		}

		v.skip, v.json, v.publ = false, false, false
	}

	return nil
}

func (p *parser_) Parse(fname string, isDir bool) (err error) {
	info, err := os.Stat(fname)
	if err != nil {
		log.Fatal(err)
	}
	if info.IsDir() {
		p.PkgDir = fname
	} else {
		p.PkgDir = filepath.Dir(fname)
	}

	fset := token.NewFileSet()
	if isDir {
		packages, err := parser.ParseDir(
			fset,
			fname,
			excludeTestFiles,
			parser.ParseComments,
		)
		if err != nil {
			return err
		}
		for _, pckg := range packages {
			ast.Walk(&visitor_{parser_: p}, pckg)
		}
	} else {
		f, err := parser.ParseFile(
			fset,
			fname,
			nil,
			parser.ParseComments,
		)
		if err != nil {
			return err
		}
		ast.Walk(&visitor_{parser_: p}, f)
	}

	// Sort structs for consistent code generation
	sort.Slice(p.Structs, func(i, j int) bool {
		return p.Structs[i].name < p.Structs[j].name
	})
	sort.Slice(p.Defines, func(i, j int) bool {
		return p.Defines[i].name < p.Defines[j].name
	})
	sort.Slice(p.Imports, func(i, j int) bool {
		return p.Imports[i].name < p.Imports[j].name
	})

	return nil
}

func excludeTestFiles(fi os.FileInfo) bool {
	return !strings.HasSuffix(fi.Name(), "_test.go")
}
