package parse

type T int

const (
	// Fn represents a parse.Function
	Fn T = iota

	// Exprsn represents a parse.Expression
	Exprsn
)

// Classifier allows types to identify themselves
type Classifier interface {
	Type() T
}
