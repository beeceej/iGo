package parse

type (
	// T is a constant identifier for concepts in Go
	T int

	// Classifier allows types to identify themselves
	Classifier interface {
		Type() T
	}
)

const (
	// Tfunction represents a parse.Function
	Tfunction T = iota

	// Texpression represents a parse.Expression
	Texpression
)
