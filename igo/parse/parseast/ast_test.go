package parseast

import (
	"fmt"
	"testing"
)

func Test_ast(t *testing.T) {
	a := &Parser{Raw: `
	type A struct {}
	func (a A) DoIt(a, c string, b []int) string {
		return ""
	}

	func (a A) DoIt(asdf string) (string, int) {
		return ""
	}
	func main() {
		a := func() string {return "sdf"}
    http.ListenAndServe(":8080", nil)
	}`}

	a.Parse()
	fns := a.Functions
	if fns != nil {
		for i, fn := range fns {
			if fn != nil {
				fmt.Println(fn.String())
			} else {
				fmt.Println(i, "is null")
			}
		}
	}
}
