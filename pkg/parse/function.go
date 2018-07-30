package parse

import (
	"errors"
	"regexp"
	"strings"
)

const (
	isFunctionExpr regexpType = iota
	identifierExpr
	argsExpr
)

type regexpType int
type exprMap map[regexpType]*regexp.Regexp

// Function is the struct necesary to parse a function
type Function struct {
	exprMap
	Raw        string
	Identifier string
	Args       []string
}

// NewFunction returns a function with the raw text
func NewFunction(raw string) (f *Function, err error) {
	f = &Function{
		Raw: raw,
		exprMap: map[regexpType]*regexp.Regexp{
			isFunctionExpr: regexp.MustCompile(`func \(?.*\)?\{\n?[.*?|[\S\s]*?\}`),
			identifierExpr: regexp.MustCompile(`(func .* \(|func .*?)\(`),
			argsExpr:       regexp.MustCompile(`\((.*?)\)`),
		},
	}

	if !f.isFunction() {
		return nil, errors.New("Not a function")
	}

	f.Identifier = f.matchIdentifier()
	f.Args = f.matchArgs()

	return f, nil
}

func (f *Function) matchIdentifier() string {
	rawIDMatch := f.exprMap[identifierExpr].FindAllStringSubmatch(f.Raw, -1)[0][1]

	// Can probably do a fancy look-around to not pull `func` out, but this is easier
	idWithoutFunc := strings.Replace(rawIDMatch, "func", "", -1)
	return strings.Trim(idWithoutFunc, " ")
}

func (f *Function) matchArgs() []string {
	return []string{}
}

func (f *Function) isFunction() bool {
	return f.exprMap[isFunctionExpr].MatchString(f.Raw)
}
