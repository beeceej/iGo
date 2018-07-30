package parse

// Expression is anything that is not a function (for now)
type Expression struct {
	Raw string
}

// Type is an identifier for the Function type
func (e *Expression) Type() T {
	return Exprsn
}


