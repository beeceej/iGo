package interpretor

import (
	"testing"
)

func TestThing(t *testing.T) {
	i := Interpretor{}
	str := `
	func add(a,b int) int {
		return a + b
	}
	func subtract(a,b int) int {
		return a - b
	}
	`

	i.Interpret(str)
	i.Interpret(`fmt.Println(subtract(1000,2))`)
}
