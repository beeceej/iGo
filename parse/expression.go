package parse

// Expression is anything that is not a function
type Expression struct {
	Raw string
}

// Type is an identifier for the Function type
func (e *Expression) Type() T {
	return Texpression
}

func (e Expression) String() string {
	return firstChar.ReplaceAllString(e.Raw, "... ")
}
