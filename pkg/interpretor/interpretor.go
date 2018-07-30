package interpretor

import "github.com/beeceej/iGo/pkg/parse"

// Interpretor is
type Interpretor struct {
	Functions map[string]*parse.Function
	History   []string
}

// Interpret will
func (i *Interpretor) Interpret(text string) {
	if i.Functions == nil {
		i.Functions = make(map[string]*parse.Function)
	}
	i.History = append(i.History, text)
	t := i.classify(text)
	for _, tv := range t {
		switch v := tv.(type) {
		case *parse.Function:
			i.Functions[v.Identifier] = v
		default:
			i.Eval(text)
		}
	}
}

func (i *Interpretor) classify(text string) []parse.Classifier {
	var t []parse.Classifier

	f, err := parse.NewFunction(text)
	if err == nil {
		for _, fi := range f {
			t = append(t, fi)
		}
		return t
	}
	t = []parse.Classifier{&parse.Expression{Raw: text}}
	return t
}
