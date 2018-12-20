package interpreter

import (
	"testing"
)

func TestThing(t *testing.T) {
	t.Skip("because code architecture is terrible right now")
	i := &Interpreter{}
	str := `
func sayHi() { 
	fmt.Println("Hello World, from iGo") 
}`

	i.Interpret(str)
	i.Interpret(`sayHi()`)
}
