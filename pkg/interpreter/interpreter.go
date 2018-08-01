package interpreter

import (
	"fmt"

	"github.com/beeceej/iGo/pkg/parse"
)

// Interpreter houses the function references and input history
type Interpreter struct {
	Functions map[string]*parse.Function
	History   []string
}

// Interpret will take some text and
// Classify it as either an expression or a function
// If it is a function it will store the reference of the Function in a map
// If the text is classified as an expression, it will evaluate the expression,
// using the function reference map if needed
func (i *Interpreter) Interpret(text string) {
	if i.Functions == nil {
		i.Functions = make(map[string]*parse.Function)
	}
	i.History = append(i.History, text)
	t := i.classify(text)
	for _, tv := range t {
		switch v := tv.(type) {
		case *parse.Function:
			i.Functions[v.Identifier] = v
			fmt.Printf("# %s\n", v.String())
			break
		case *parse.Expression:
			fmt.Printf(">> %s\n", text)
			i.Eval(text)
			break
		}
	}
}

func (i *Interpreter) classify(text string) []parse.Classifier {
	var t []parse.Classifier
	fns := parse.Functions(text)
	if len(fns) == 0 {
		t = []parse.Classifier{&parse.Expression{Raw: text}}
	}
	for _, fi := range fns {
		t = append(t, fi)
	}
	return t
}
