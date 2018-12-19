package parse

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"regexp"
	"strings"
)

var hasPackageStatementRegexp = regexp.MustCompile("^package.*")

// ASTParse is
type ASTParse struct {
	Raw       string
	Functions []*Function

	fset *token.FileSet
	root ast.Node
}

type node struct {
	ast.Node
}

func (n *node) asString() (string, error) {
	var b *bytes.Buffer
	err := printer.Fprint(b, nil, n)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

// Parse is
func (a *ASTParse) Parse() {
	a.Setup()
	ast.Inspect(a.root, func(n ast.Node) bool {
		err := ifFunctionDeclaration(
			compose(
				a.ParseFn,
			), n)
		if err != nil {
			log.Println(err.Error())
			return false
		}
		return true
	})
}

// Setup is
func (a *ASTParse) Setup() {
	if !hasPackageStatementRegexp.MatchString(a.Raw) {
		const withPkg = `package _
		%s`
		a.Raw = fmt.Sprintf(withPkg, a.Raw)
	}
	a.Functions = []*Function{}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "/tmp/tmp.go", a.Raw, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	a.fset = fset
	a.root = node
}

// ParseFn is
func (a *ASTParse) ParseFn(n *ast.FuncDecl) error {
	var (
		identifier      string
		params          string
		returnSignature string
		body            string
		err             error
	)
	identifier = n.Name.Name
	params = a.getFunctionParameters(n)

	returnSignature = a.getFunctionReturnSignature(n)

	if body, err = (&node{n.Body}).asString(); err != nil {
		return err
	}

	a.Functions = append(a.Functions, &Function{
		Identifier: identifier,
		Params:     params,
		// This is actually an ast.Node in-itself, we could parse sub functions recursively,
		// not sure that's completely necessary just yet though
		Body:   body,
		Return: returnSignature,
	})
	return nil
}

func (a *ASTParse) getFunctionParameters(n *ast.FuncDecl) string {
	var (
		params []string
		buffer bytes.Buffer
	)
	if n.Type.Results == nil {
		return ""
	}

	offset := n.Pos()
	printer.Fprint(&buffer, a.fset, n)
	rawFn := buffer.String()
	for _, pGroup := range n.Type.Params.List {
		p := pGroup.Type.Pos() - offset
		e := pGroup.Type.End() - offset
		groupType := rawFn[p:e]
		params = append(params, paramGroup(pGroup.Names, groupType))
	}

	return strings.Join(params, ", ")
}

func (a *ASTParse) getFunctionReturnSignature(n *ast.FuncDecl) string {
	var (
		returns []string
		fnRaw   string
		err     error
	)
	if n.Type.Results == nil {
		return ""
	}
	offset := n.Pos()

	if fnRaw, err = (&node{n}).asString(); err != nil {
		fnRaw = ""
	}

	for _, rGroup := range n.Type.Results.List {
		p := rGroup.Type.Pos() - offset
		e := rGroup.Type.End() - offset
		returns = append(returns, fnRaw[p:e])
	}

	return strings.Join(returns, ", ")

}

func paramGroup(idents []*ast.Ident, t string) string {
	if len(idents) < 1 {
		return ""
	}
	paramNames := []string{}
	for _, i := range idents {
		if i.Name == "" {
			continue
		}
		paramNames = append(paramNames, i.Name)
	}

	return fmt.Sprint(strings.Join(paramNames, ", "), " ", t)
}
