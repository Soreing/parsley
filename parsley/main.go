package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var specifiedName = flag.String("output_filename", "", "specify the filename of the output")
var allTypes = flag.Bool("all", false, "generate scanners for all types in a file")
var lowerCase = flag.Bool("lower_case", false, "use lower case names by default")
var camelCase = flag.Bool("camel_case", false, "use camel case names by default")
var kebabCase = flag.Bool("kebab_case", false, "use kebab case names by default")
var snakeCase = flag.Bool("snake_case", false, "use snake case names by default")
var pascalCase = flag.Bool("pascal_case", false, "use pascal case names by default")
var allPublic = flag.Bool("public", false, "include private fields in encoding/decoding")

func main() {
	flag.Parse()
	files := flag.Args()

	if len(files) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for _, file := range files {
		if err := generate(file); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func generate(fname string) error {
	finf, err := os.Stat(fname)
	if err != nil {
		return err
	}

	p := Parser{
		AllTypes:  *allTypes,
		AllPublic: *allPublic,
	}
	if err := p.Parse(fname, finf.IsDir()); err != nil {
		return fmt.Errorf("error parsing %v: %v", fname, err)
	}

	var outName string
	if *specifiedName != "" {
		outName = *specifiedName
	} else if finf.IsDir() {
		outName = filepath.Join(fname, p.PkgName+"_parsley.go")
	} else {
		outName = strings.TrimSuffix(fname, ".go") + "_parsley.go"
	}

	g := NewGenerator()

	g.SetPackage(p.PkgName)

	if *lowerCase {
		g.SetDefaultCase(LOWER_CASE)
	} else if *camelCase {
		g.SetDefaultCase(CAMEL_CASE)
	} else if *kebabCase {
		g.SetDefaultCase(KEBAB_CASE)
	} else if *snakeCase {
		g.SetDefaultCase(SNAKE_CASE)
	} else if *pascalCase {
		g.SetDefaultCase(PASCAL_CASE)
	} else {
		g.SetDefaultCase(LOWER_CASE)
	}

	g.WriteHeader()

	pkgs := g.GetRequiredPackages(p.Imports, p.Defines, p.Structs)
	g.WriteImports(pkgs)

	for _, st := range p.Structs {
		g.WriteStruct(st)
	}

	for _, df := range p.Defines {
		g.WriteDefine(df)
	}

	fmtd, err := g.Format(g.ReadAll())
	if err != nil {
		log.Fatalf("formating output: %s", err)
	}

	err = ioutil.WriteFile(outName, fmtd, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}

	return nil
}
