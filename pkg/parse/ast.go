package parse

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"regexp"
)

var hasPackageStatementRegexp = regexp.MustCompile("^package.*")

// AST is
type AST struct {
	RawText   string
	Functions []*Function
}

// Parse is
func (a *AST) Parse() {
	a.Setup()
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, "tmp.go", a.RawText, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	var buffer bytes.Buffer
	var pretty string
	ast.Inspect(node, func(n ast.Node) bool {
		if aNode, ok := n.(*ast.FuncDecl); ok {
			printer.Fprint(&buffer, fset, aNode)
			pretty = buffer.String()
			a.Functions = append(
				a.Functions,
				getFunction(aNode, fset, pretty),
			)
		}

		if aNode, ok := n.(*ast.FuncLit); ok {

			printer.Fprint(&buffer, fset, aNode)
			pretty = buffer.String()
			a.Functions = append(a.Functions, &Function{
				Raw:        pretty,
				Identifier: "",
				// Params:     nil,
			})
			// printer.Fprint(os.Stdout, fset, aNode.Recv.)
		}

		if aNode, ok := n.(*ast.FuncType); ok {

			printer.Fprint(&buffer, fset, aNode)
			pretty = buffer.String()
			a.Functions = append(a.Functions, &Function{
				Raw:        pretty,
				Identifier: "aNode.Name.Name",
				// Params:     nil,
			})
			// printer.Fprint(os.Stdout, fset, aNode.Recv.)
		}
		return true
	})
	fmt.Println(a)
}

// Setup is
func (a *AST) Setup() {
	if !hasPackageStatementRegexp.MatchString(a.RawText) {
		a.RawText = fmt.Sprintf(`package _
		%s`, a.RawText)
	}

	if a.Functions == nil {
		a.Functions = make([]*Function, 1, 1)
	}
}

func getFunction(fdecl *ast.FuncDecl, fset *token.FileSet, src string) *Function {
	var offset = fdecl.Pos()
	namesByString := func(idents []*ast.Ident) string {
		var str string
		if len(idents) == 1 {
			str = fmt.Sprintf("%s", idents[0])
		} else if len(idents) > 0 {
			for _, n := range idents {
				str = fmt.Sprintf("%s, %s", str, n)
			}
			str = str[2:]
		}
		return str
	}

	var paramsString string
	var name string
	var body string
	var returnSignature string
	var buffer bytes.Buffer
	ast.Inspect(fdecl, func(n ast.Node) bool {
		printer.Fprint(&buffer, fset, fdecl.Body)

		body = buffer.String()
		if fd, ok := n.(*ast.FuncDecl); ok {
			name = fd.Name.Name
			if len(fd.Type.Params.List) > 0 {
				for _, param := range fd.Type.Params.List {
					paramsString += fmt.Sprintf("%s %s, ", namesByString(param.Names), src[param.Type.Pos()-offset:param.Type.End()-offset])
				}
			}
			fmt.Println("AAAAAAA")
			printer.Fprint(os.Stdout, fset, fd.Type)
			if fd.Type.Results != nil {

				if len(fd.Type.Results.List) > 0 {
					for _, result := range fd.Type.Results.List {
						returnSignature += fmt.Sprintf("%s, ", result.Type)
					}
				}
			}
		}

		return true
	})
	if paramsString != "" {
		paramsString = paramsString[0 : len(paramsString)-2]
	}
	f := &Function{
		Identifier: name,
		Params:     paramsString,
		Body:       body,
		Return:     returnSignature,
	}
	return f
}
