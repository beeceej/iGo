package parse

import "testing"

func Test_ast(t *testing.T) {
	a := &AST{RawText: `
	type A struct {}
	func (a A) DoIt(a, c string, b []int) string {
		return ""
	}

	func (a A) DoIt(asdf string) (string, int) {
		return ""
	}
	func main() {
		a := func() string {return "sdf"}
    http.ListenAndServe(":8080",nil)

	}`}

	a.Parse()
}
