package parse

import (
	"testing"
)

func Test_extract_a_bunch_of_combinations(t *testing.T) {
	fns := `
func a     () string {
	return ""
}
func b() int{
	return 1
}
func c() string{
	return "asdf"
}

func abcdefg() (string str, err error) (abc str, cde err){
	return ""
}

func nnn() (string str, err ...error) (str, int){
	return str, int
}

func b() int{
	return 1
}
func c() string{
	return "asdf"
}`

	fnSlice := extract(fns)

	if len(fnSlice) != 7 {
		t.Fail()
	}

}

func Test_extract_expectFunctionParsedCorrect(t *testing.T) {
	fnRaw := `
func a     () string {
	return ""
}
`
	fnSlice := extract(fnRaw)
	if fnSlice[0].Body != `return ""` {
		t.Fail()
	}
	if fnSlice[0].Identifier != `a` {
		t.Fail()
	}

	if fnSlice[0].Params != `` {
		t.Fail()
	}

	if fnSlice[0].Return != `string` {
		t.Fail()
	}
}

func Test_extract_expectThreeExtracted(t *testing.T) {

	threeFns := `
func a() string {
	return ""
}
func b() int{
	return 1
}
func c(s string, b string) string{
	return c + b
}
`

	raw := extract(threeFns)

	if len(raw) != 3 {
		t.Fail()
	}
}

func Test_extract_expectNoneExtracted(t *testing.T) {
	fns := `
f c() {

}
`

	raw := extract(fns)

	if len(raw) != 0 {
		t.Fail()
	}
}

func Test_extract_expectTwoExtracted(t *testing.T) {
	fns := `
func one(a,b,c string) string {
	return ""
}

function hi(a string) string {
	return a
}

func three(i int) int {
	return i + i
}
`

	raw := extract(fns)

	if len(raw) != 2 {
		t.Fail()
	}
}

func Test_extract_expectOneExtracted(t *testing.T) {
	fns := `
func abc(i int) int {
	return func() int{
		return i
	}()
}
`

	raw := extract(fns)

	if len(raw) != 1 {
		t.Fail()
	}
}

func Test_matchIdentifier_expectIdentified(t *testing.T) {
	fns := `
func abc(i int) int {
	return func() int{
		return i
	}
}
`

	id := matchIdentifier(fns)
	if id != "abc" {
		t.Fail()
	}
}

func Test_matchIdentifier_expectNotIdentified(t *testing.T) {
	fns := `
func (i int) int {
	return func() int{
		return i
	}
}
`

	id := matchIdentifier(fns)
	if id != "" {
		t.Fail()
	}
}

func Test_matchReturn_expectIdentified(t *testing.T) {
	fns := `
func abc(i int) int {
	return func() int{
		return i
	}
}
`

	returnVal := matchReturn(fns)
	if returnVal != "int" {
		t.Fail()
	}
}

func Test_matchReturn_expectEmpty(t *testing.T) {
	fns := `
func (i int) {
	return func() int{
		return i
	}
}
`

	returnVal := matchReturn(fns)
	if returnVal != "" {
		t.Fail()
	}
}

func Test_firstMatchExpectFound(t *testing.T) {
	matches := [][]string{
		[]string{"a"},
	}

	match := firstMatch(matches)
	if match != "a" {
		t.Fail()
	}
}

func Test_firstMatchExpectEmpty(t *testing.T) {
	matches := [][]string{}

	match := firstMatch(matches)
	if match != "" {
		t.Fail()
	}
}

func Test_matchParamsExpectFound(t *testing.T) {
	fns := `
func (i int) {
	return func() int{
		return i
	}
}
`

	param := matchParams(fns)
	if param != "i int" {
		t.Fail()
	}

	fns = `
func (i ...int) {
	return func() int{
		return i
	}
}
`
	param = matchParams(fns)
	if param != "i ...int" {
		t.Fail()
	}
}

func Test_matchParamsExpectNone(t *testing.T) {
	fns := `
func () {
	return func() int{
		return i
	}
}
`

	param := matchParams(fns)
	if param != "" {
		t.Fail()
	}
}
func Test_getMatchExpectFound(t *testing.T) {
	matches := [][]string{
		[]string{"asbc", "1992"},
		[]string{"asbcasdfasdf"},
	}

	match := getMatch(0, 0, matches)
	if match != "asbc" {
		t.Fail()
	}

	match = getMatch(0, 1, matches)
	if match != "1992" {
		t.Fail()
	}
	match = getMatch(1, 1, matches)
	if match != "" {
		t.Fail()
	}
}

func Test_getMatchExpectExpectEmpty(t *testing.T) {
	matches := [][]string{
		[]string{"asbc", "1992"},
		[]string{"asbcasdfasdf"},
	}

	match := getMatch(1, 1, matches)
	if match != "" {
		t.Fail()
	}
}
