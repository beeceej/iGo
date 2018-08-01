package interpreter

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestThing(t *testing.T) {
	// t.Skip("More of an integration test for now")
	i := Interpreter{}
	str := `
"func sayHi() { 
	fmt.Println("Hello World, from iGo") 
}"`

	i.Interpret(str)
	spew.Dump(i.Functions)
	i.Interpret(`sayHi()`)
}
