package parse

// T is a constant identifier for concepts in Go
type T int

const (
	// Tfunction represents a parse.Function
	Tfunction T = iota

	// Texpression represents a parse.Expression
	Texpression
)

// Classifier allows types to identify themselves
type Classifier interface {
	Type() T
}
