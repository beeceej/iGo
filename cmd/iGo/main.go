package main

import (
	"os"
	"text/template"
)

const mainTmpl = `package main
func main() {
	{{.}}
}
`

func main() {
	t := template.Must(template.New("mainTmpl").Parse(mainTmpl))
	t.Execute(os.Stdout, `fmt.Println("iGoIsCool")`)
}
