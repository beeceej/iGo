package interpreter

import (
	"testing"
)

func TestThing(t *testing.T) {
	t.Skip("More of an integration test for now")
	i := Interpreter{}
	str := `
"func sayHi() { 
	fmt.Println("Hello World, from iGo") 
}`

	i.Interpret(str)
	i.Interpret(`sayHi()`)
}
