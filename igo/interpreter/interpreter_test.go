package interpreter

import (
	"testing"
)

func TestInterpret(t *testing.T) {
	t.Skip("because code architecture isn't quite there yet")
	i := &Interpreter{}
	sayHiFunc := `
func sayHi() {
	fmt.Println("Hello World, from iGo")
}`

	i.Interpret(sayHiFunc)
	i.Interpret(`sayHi()`)
}
