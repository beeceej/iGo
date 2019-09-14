package parse

import (
	"fmt"
	"regexp"
)

var firstChar = regexp.MustCompile(`(?m)^`)

// Function is the in memory representation of a function
type Function struct {
	// Raw, the raw input which was determined to be a function
	Raw string

	// Identifier, the identifier of the function.
	// For example:
	// func a() {}
	// Identifier would = a
	Identifier string

	// Params is a raw string which identifies the parameters of the function
	Params string

	// Return is the return signature of the function
	Return string

	// Body is everything which is contained in the function
	Body string
}

func (f Function) String() (str string) {
	if f.Return == "" {
		str = fmt.Sprintf("%s :: Function (%s) -> ()", f.Identifier, f.Params)
	} else {
		str = fmt.Sprintf("%s :: Function (%s) -> (%s)", f.Identifier, f.Params, f.Return)
	}

	return fmt.Sprintf("%s\n%s",
		firstChar.ReplaceAllString(f.Raw, "... "),
		firstChar.ReplaceAllString(str, ">>> "),
	)
}

// Type is an identifier for the Function type
func (f Function) Type() T {
	return Tfunction
}
