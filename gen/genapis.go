package main

import (
	"embed"
	"fmt"
	"go/types"
	"log"
	"path"
	"strings"

	"golang.org/x/tools/go/packages"
)

//go:generate go run .

//go:embed funcs/*
var funcs embed.FS

const ebitenURI = "github.com/hajimehoshi/ebiten/v2"

func main() {
	if err := generate(); err != nil {
		log.Fatal(err)
	}
}

func generate() error {
	const dirname = "funcs"

	entries, err := funcs.ReadDir(dirname)
	if err != nil {
		return err
	}

	for _, e := range entries {
		contetns, err := funcs.ReadFile(path.Join(dirname, e.Name()))
		if err != nil {
			return err
		}

		lines := strings.Split(string(contetns), "\n")
		for _, l := range lines {
			if err := printSignature(l); err != nil {
				return err
			}
		}
	}

	return nil
}

func printSignature(line string) error {
	fqn := FQN(strings.TrimSpace(line))
	if len(fqn) == 0 {
		return nil
	}

	uri, err := getURI(fqn)
	if err != nil {
		return err
	}

	pkgs, err := getPackages(uri)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		if len(pkg.Errors) > 0 {
			fmt.Println(pkg.Errors)
		}

		names := pkg.Types.Scope().Names()
		for _, name := range names {
			if name == fqn.Function() {
				fn := pkg.Types.Scope().Lookup(name)
				if sig, ok := fn.Type().(*types.Signature); ok {
					fmt.Println(name, sig.String())
				}
			}
		}
	}

	return nil
}

var packageCache map[string][]*packages.Package

func init() {
	packageCache = make(map[string][]*packages.Package)
}

func getPackages(uri string) ([]*packages.Package, error) {
	if pkgs, ok := packageCache[uri]; ok {
		return pkgs, nil
	}

	conf := packages.Config{
		Mode: packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
	}

	pkgs, err := packages.Load(&conf, uri)
	if err != nil {
		return nil, err
	}
	packageCache[uri] = pkgs

	return pkgs, nil
}

func getURI(fqn FQN) (string, error) {
	if len(fqn.Package()) == 0 {
		return "", fmt.Errorf("func name must be 'package.Function' format, but %s is not", fqn)
	}

	if fqn.Package() == "ebiten" {
		return ebitenURI, nil
	} else {
		return fmt.Sprintf("%s/%s", ebitenURI, fqn.Package()), nil
	}
}

type FQN string

func (n FQN) Split() (pkgName, funcName string) {
	items := strings.Split(string(n), ".")
	if len(items) < 2 {
		return "", string(n)
	}

	return items[0], items[1]
}

func (n FQN) Package() string {
	p, _ := n.Split()
	return p
}

func (n FQN) Function() string {
	_, f := n.Split()
	return f
}
