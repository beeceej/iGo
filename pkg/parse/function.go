package parse

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	// holds double duty for classifying, and extracting functions from raw text
	isFunctionExpr regexpType = iota
	identifierExpr
	argsExpr
	returnExpr
)

var expressions = map[regexpType]*regexp.Regexp{
	isFunctionExpr: regexp.MustCompile(`func \(?.*\)?\{\n?(.*|\s|\S)*?(\})`),
	identifierExpr: regexp.MustCompile(`(func .* \(|func .*?)\(`),
	argsExpr:       regexp.MustCompile(`\((.*?)\)`),
	returnExpr:     regexp.MustCompile(`\) .* {`),
}

type regexpType int

// Function is the in memory representation of a function,
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
}

// Functions returns a function with the raw text
// This will return either a reference to a Function object, or an error
// An error here does not mean there is anything wrong with the input,
// Just that it could not be recognized as a function.
func Functions(raw string) (fns []*Function) {
	var rawFns []string
	if rawFns = extract(raw); len(rawFns) == 0 {
		return []*Function{}
	}

	for _, pFun := range rawFns {
		fns = append(fns, newFn(pFun))
	}

	return fns
}

func newFn(raw string) (f *Function) {
	f = &Function{
		Raw:        raw,
		Identifier: matchIdentifier(raw),
		Params:     matchParams(raw),
		Return:     matchReturn(raw),
	}

	return f
}

func matchIdentifier(raw string) string {
	rawIDMatch := expressions[identifierExpr].FindAllStringSubmatch(raw, -1)[0][1]
	// Can probably do a fancy look-around to not pull `func` out, but this is easier
	idWithoutFunc := strings.Replace(rawIDMatch, "func", "", -1)
	return strings.Trim(idWithoutFunc, " ")
}

func matchParams(raw string) string {
	rawParamMatch := expressions[argsExpr].FindAllStringSubmatch(raw, -1)
	return getMatch(0, 1, rawParamMatch)
}

func matchReturn(raw string) string {
	rmOpenBrace := func(raw string) string {
		return strings.Replace(raw, "{", "", -1)
	}
	rmCloseParen := func(raw string) string {
		return strings.Replace(raw, ")", "", -1)
	}
	rawReturnMatches := expressions[returnExpr].FindAllStringSubmatch(raw, -1)
	if len(rawReturnMatches) == 0 {
		return ""
	}
	rawReturn := firstMatch(rawReturnMatches)
	return strings.Trim(rmCloseParen(rmOpenBrace(rawReturn)), " ")
}

func extract(raw string) (rawFns []string) {
	if !expressions[isFunctionExpr].MatchString(raw) {
		return []string{}
	}
	matches := expressions[isFunctionExpr].FindAllStringSubmatch(raw, -1)
	for _, match := range matches {
		lastBracket := match[1] // Need to clean this expression up to cleanly extract the last bracket
		rawFns = append(rawFns, match[0]+lastBracket)
	}

	return rawFns
}

// firstMatch will return the first match, of the first capture group if it exists,
// if it doesn't exist it will return an empty string
// match[i] will give the capture groups associated with the i'th match,
// match[i][j] will give the j'th capture group associated with the i'th match,
func firstMatch(matches [][]string) (match string) {
	if len(matches) != 0 && len(matches[0]) != 0 {
		match = matches[0][0]
	}
	return match
}

// getMatch will return empty string if i,j is not found in the matches slice
// or it will return the matches/groups at matches[i][j]
func getMatch(i, j int, matches [][]string) (match string) {
	if len(matches) <= i {
		match = ""
	} else if len(matches[i]) <= j {
		match = ""
	} else {
		match = matches[i][j]
	}
	return match
}

// Type is an identifier for the Function type
func (f Function) Type() T {
	return Tfunction
}

func (f Function) String() (str string) {
	if f.Return == "" {
		str = fmt.Sprintf("%s :: Function (%s) -> ()\n", f.Identifier, f.Params)
	} else {
		str = fmt.Sprintf("%s :: Function (%s) -> %s\n", f.Identifier, f.Params, f.Return)
	}
	return str
}
