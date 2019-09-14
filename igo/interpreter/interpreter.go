package interpreter

import (
	"github.com/beeceej/iGo/igo/parse"
	"github.com/beeceej/iGo/igo/parse/parseast"
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
func (i *Interpreter) Interpret(text string) string {
	if i.Functions == nil {
		i.Functions = make(map[string]*parse.Function)
	}
	i.History = append(i.History, text)
	t := i.classify(text)
	for _, tv := range t {
		switch v := tv.(type) {
		case *parse.Function:
			i.Functions[v.Identifier] = v
			return v.String()
		case *parse.Expression:
			return i.Eval(text)
		}
	}
	return ""
}

func (i *Interpreter) classify(text string) []parse.Classifier {
	var t []parse.Classifier
	p := &parseast.Parser{
		Raw: text,
	}
	p.Parse()
	fns := p.Functions
	if len(fns) == 0 {
		t = []parse.Classifier{&parse.Expression{Raw: text}}
	}

	for _, fi := range fns {
		t = append(t, fi)
	}
	return t
}
