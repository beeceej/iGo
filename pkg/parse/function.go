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
	returnExpr
)

type regexpType int
type exprMap map[regexpType]*regexp.Regexp

// Function is the struct necesary to parse a function
type Function struct {
	exprMap

	// Raw, the raw input which was determined to be a function
	Raw string

	// Identifier, the identifier of the function.
	// For example:
	// func a() {}
	// Identifier would = a
	Identifier string

	// Args is a raw string which identifies the parameters of the function
	Args string

	// Return is the return signature of the function
	Return string
}

// NewFunction returns a function with the raw text
// This will return either a reference to a Function object, or an error
// An error here does not mean there is anything wrong with the input,
// Just that it could not be recognized as a function.
func NewFunction(raw string) (fns []*Function, err error) {
	f := &Function{
		Raw: raw,
		exprMap: map[regexpType]*regexp.Regexp{
			isFunctionExpr: regexp.MustCompile(`func \(?.*\)?\{\n?[.*?|[\S\s]*?\}`),
			identifierExpr: regexp.MustCompile(`(func .* \(|func .*?)\(`),
			argsExpr:       regexp.MustCompile(`\((.*?)\)`),
			returnExpr:     regexp.MustCompile(`\) .* {`),
		},
	}

	var fnsRaw []string

	if fnsRaw, err = f.matchFunctions(); err == nil {
		for _, rawFn := range fnsRaw {
			tmp, err := newFunction(rawFn)
			if err != nil {
				return fns, errors.New("not a fn")
			}
			if tmp != nil {
				fns = append(fns, tmp)
			}
		}
	} else {
		return nil, errors.New("No functions found")
	}

	return fns, nil
}

func newFunction(raw string) (*Function, error) {
	f := &Function{
		Raw: raw,
		exprMap: map[regexpType]*regexp.Regexp{
			isFunctionExpr: regexp.MustCompile(`func \(?.*\)?\{\n?[.*?|[\S\s]*?\}`),
			identifierExpr: regexp.MustCompile(`(func .* \(|func .*?)\(`),
			argsExpr:       regexp.MustCompile(`\((.*?)\)`),
			returnExpr:     regexp.MustCompile(`\) .* {`),
		},
	}
	f.Identifier = f.matchIdentifier()
	f.Args = f.matchArgs()
	f.Return = f.matchReturn()
	return f, nil
}
func (f *Function) matchIdentifier() string {
	rawIDMatch := f.exprMap[identifierExpr].FindAllStringSubmatch(f.Raw, -1)[0][1]

	// Can probably do a fancy look-around to not pull `func` out, but this is easier
	idWithoutFunc := strings.Replace(rawIDMatch, "func", "", -1)
	return strings.Trim(idWithoutFunc, " ")
}

func (f *Function) matchArgs() string {
	rawArgsMatch := f.exprMap[argsExpr].FindAllStringSubmatch(f.Raw, -1)

	return rawArgsMatch[0][1]
}

func (f *Function) matchReturn() string {
	rawReturnMatch := f.exprMap[returnExpr].FindAllStringSubmatch(f.Raw, -1)
	if len(rawReturnMatch) == 0 {
		return ""
	}
	return strings.Trim(strings.Replace(strings.Replace(rawReturnMatch[0][0], "{", "", -1), ")", "", -1), " ")
}

func (f *Function) matchFunctions() (rawFns []string, err error) {
	if !f.exprMap[isFunctionExpr].MatchString(f.Raw) {
		return rawFns, errors.New("Not a function")
	}
	matches := f.exprMap[isFunctionExpr].FindAllStringSubmatch(f.Raw, -1)
	for _, match := range matches {
		rawFns = append(rawFns, match[0])
	}
	return rawFns, err
}

// Type is an identifier for the Function type
func (f Function) Type() T {
	return Fn
}
