package parseast

import (
	"go/ast"
)

type (
	funcDeclarationHandler func(*ast.FuncDecl) error
	funcLiteralHandler     func(*ast.FuncLit) error
	funcTypeHandler        func(*ast.FuncType) error
)

func ifFunctionDeclaration(h funcDeclarationHandler, n ast.Node) error {
	if fnDecl, ok := n.(*ast.FuncDecl); ok {
		return h(fnDecl)
	}
	return nil
}

func ifFunctionLiteral(h funcLiteralHandler, n ast.Node) error {
	if fnLiteral, ok := n.(*ast.FuncLit); ok {
		return h(fnLiteral)
	}
	return nil
}

func ifFunctionType(h funcTypeHandler, n ast.Node) error {
	if fnType, ok := n.(*ast.FuncType); ok {
		return h(fnType)
	}
	return nil
}

// compose declaration handlers into one function?
func compose(fns ...funcDeclarationHandler) funcDeclarationHandler {
	return func(n *ast.FuncDecl) error {
		for _, fn := range fns {
			if err := fn(n); err != nil {
				return err
			}
		}
		return nil
	}
}
