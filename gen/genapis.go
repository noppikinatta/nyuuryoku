package main

import (
	"embed"
	"fmt"
	"go/types"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"golang.org/x/tools/go/packages"
)

//go:generate go run .
//go:generate go fmt ../...

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

	// Map of strings to remove from function names for each file
	strsToRemoveMap := map[string][]string{
		"mouse.txt":    {"MouseButton"},
		"gamepad.txt":  {"Gamepad"},
		"keyboard.txt": {"Keys", "Key"},
	}

	entries, err := funcs.ReadDir(dirname)
	if err != nil {
		return err
	}

	for _, e := range entries {
		// Generate TypeName from filename
		name := strings.TrimSuffix(e.Name(), ".txt")
		typeName := strings.ToUpper(name[0:1]) + name[1:]
		t := &Type{TypeName: typeName}

		contents, err := funcs.ReadFile(path.Join(dirname, e.Name()))
		if err != nil {
			return err
		}

		lines := strings.Split(string(contents), "\n")
		for _, l := range lines {
			if err := generateAPI(l, t, strsToRemoveMap[e.Name()]...); err != nil {
				return err
			}
		}

		// Generate file
		if err := generateFile(t); err != nil {
			return err
		}
	}

	return nil
}

func generateAPI(line string, t *Type, strsToRemove ...string) error {
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
			return fmt.Errorf("package errors: %v", pkg.Errors)
		}

		names := pkg.Types.Scope().Names()
		for _, name := range names {
			if name == fqn.Function() {
				fn := pkg.Types.Scope().Lookup(name)
				if sig, ok := fn.Type().(*types.Signature); ok {
					api := NewAPI(t, fqn, sig, strsToRemove...)
					t.APIs = append(t.APIs, *api)
				}
			}
		}
	}

	return nil
}

func generateFile(t *Type) error {
	tmpl, err := template.New("source").Parse(srcTemplate)
	if err != nil {
		return err
	}

	filename := filepath.Join("../", t.LowerCaseTypeName()+".go")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, t)
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

type Type struct {
	TypeName string
	APIs     []API
}

func (t *Type) Receiver() string {
	return t.LowerCaseTypeName()[0:1]
}

func (t *Type) LowerCaseTypeName() string {
	return strings.ToLower(t.TypeName)
}

type API struct {
	Type
	FieldName        string
	OriginalFuncName string
	ShortenFuncName  string
	Args             []Arg
	ReturnType       string
}

func (a *API) ArgsString() string {
	strs := make([]string, len(a.Args))
	for i, arg := range a.Args {
		strs[i] = arg.String()
	}
	return strings.Join(strs, ", ")
}

func (a *API) ArgNames() string {
	names := make([]string, len(a.Args))
	for i, arg := range a.Args {
		names[i] = arg.Name
	}
	return strings.Join(names, ", ")
}

type Arg struct {
	Name string
	Type string
}

func (a *Arg) String() string {
	return fmt.Sprintf("%s %s", a.Name, a.Type)
}

func NewAPI(t *Type, fqn FQN, sig *types.Signature, strsToRemove ...string) *API {
	// Sort strings by length in descending order
	sort.Slice(strsToRemove, func(i, j int) bool {
		return len(strsToRemove[i]) > len(strsToRemove[j])
	})

	// Get original function name with package
	originalFuncName := string(fqn)

	// Generate field name (first letter to lowercase + "Fn" suffix)
	fieldName := strings.ToLower(fqn.Function()[0:1]) + fqn.Function()[1:] + "Fn"

	// Generate shortened function name
	shortenFuncName := fqn.Function()
	for _, str := range strsToRemove {
		shortenFuncName = strings.Replace(shortenFuncName, str, "", -1)
	}

	converter := NewTypeStringConverter()

	// Generate argument strings
	var args []Arg
	for i := 0; i < sig.Params().Len(); i++ {
		param := sig.Params().At(i)
		paramType := converter.Convert(param.Type())
		args = append(args, Arg{Name: param.Name(), Type: paramType})
	}

	// Generate return type strings
	var returns []string
	for i := 0; i < sig.Results().Len(); i++ {
		result := sig.Results().At(i)
		resultType := converter.Convert(result.Type())
		returns = append(returns, resultType)
	}
	returnType := strings.Join(returns, ", ")
	if sig.Results().Len() > 1 {
		returnType = "(" + returnType + ")"
	}

	return &API{
		Type:             *t,
		FieldName:        fieldName,
		OriginalFuncName: originalFuncName,
		ShortenFuncName:  shortenFuncName,
		Args:             args,
		ReturnType:       returnType,
	}
}

type TypeStringConverter struct {
	Prefixes  map[string]string
	Internals map[string]string
}

func NewTypeStringConverter() *TypeStringConverter {
	return &TypeStringConverter{
		Prefixes: map[string]string{
			"github.com/hajimehoshi/ebiten/v2":           "ebiten",
			"github.com/hajimehoshi/ebiten/v2/inpututil": "inpututil",
		},
		Internals: map[string]string{
			"/internal/gamepad.":           ".Gamepad",
			"/internal/gamepaddb.Standard": ".StandardGamepad",
		},
	}
}

func (c *TypeStringConverter) Convert(t types.Type) string {
	tName := t.String()
	for prefix, shortPkg := range c.Prefixes {
		tName = strings.Replace(tName, prefix, shortPkg, 1)
	}
	for internal, shortInternal := range c.Internals {
		tName = strings.Replace(tName, internal, shortInternal, 1)
	}
	return tName
}

const srcTemplate = `// CODE GENERATED BY genapis.go. DO NOT EDIT.

package nyuuryoku

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type {{.TypeName}} struct {
{{range .APIs -}}
	{{.FieldName}} func({{.ArgsString}}) {{.ReturnType}}
{{end}}
}

func New{{.TypeName}}() *{{.TypeName}} {
	{{.Receiver}} := &{{.TypeName}}{}
	s := New{{.TypeName}}Setter({{.Receiver}})
	s.SetDefault()

	return {{.Receiver}}
}

{{range .APIs -}}
func ({{.Receiver}} *{{.TypeName}}) {{.ShortenFuncName}}({{.ArgsString}}) {{.ReturnType}} {
	return {{.Receiver}}.{{.FieldName}}({{.ArgNames}})
}
{{end}}

type {{.TypeName}}Setter struct {
	{{.LowerCaseTypeName}} *{{.TypeName}}
}

func New{{.TypeName}}Setter({{.Receiver}} *{{.TypeName}}) *{{.TypeName}}Setter {
	return &{{.TypeName}}Setter{{"{"}}{{.Receiver}}{{"}"}}
}

func (s *{{.TypeName}}Setter) SetDefault() {
	{{range .APIs -}}
		s.Set{{.ShortenFuncName}}Func({{.OriginalFuncName}})
	{{end}}
}

{{range .APIs -}}
func (s *{{.TypeName}}Setter) Set{{.ShortenFuncName}}Func(fn func({{.ArgsString}}) {{.ReturnType}}) {
	s.{{.LowerCaseTypeName}}.{{.FieldName}} = fn
}
{{end}}
`
