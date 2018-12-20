package interpreter

import (
	"testing"
)

func TestThing(t *testing.T) {

	i := &Interpreter{}
	str := `
func sayHi() { 
	fmt.Println("Hello World, from iGo") 
}`

	i.Interpret(str)
	i.Interpret(`sayHi()`)
}
