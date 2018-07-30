package interpretor

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

const tmpl = `package main
func main() {
	{{ .M }}
}

{{range .F }}
	{{.}}
{{end}}
`

// EvalData is
type EvalData struct {
	M string
	F []string
}

const path = "/Users/brianjones/development/golib/src/github.com/beeceej/iGo/cmd/iGoBin/exe.go"

// Eval will evaluate the text
func (i *Interpretor) Eval(text string) {
	var ed *EvalData
	t := template.Must(template.New("mainTmpl").Parse(tmpl))
	f, err := os.Create("/Users/brianjones/development/golib/src/github.com/beeceej/iGo/cmd/iGoBin/exe.go")
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	fns := make([]string, len(i.Functions))
	for _, fn := range i.Functions {
		fns = append(fns, fn.Raw)
	}
	ed = &EvalData{
		M: text,
		F: fns,
	}
	if err := t.Execute(f, ed); err != nil {
		fmt.Println(err.Error())
	}

	cmd := exec.Command("/Users/brianjones/development/golib/bin/goimports", "-w", "/Users/brianjones/development/golib/src/github.com/beeceej/iGo/cmd/iGoBin/exe.go")
	_, err = cmd.Output()

	if err != nil {
		fmt.Println("Error calling goimports", err.Error())
		return
	}
	f.Sync()

	cmd = exec.Command("go", "run", "/Users/brianjones/development/golib/src/github.com/beeceej/iGo/cmd/iGoBin/exe.go")
	b, err := cmd.Output()
	if err != nil {
		fmt.Println("Error calling go run", err.Error())
	}
	fmt.Println(fmt.Sprintf(">> %s", string(b)))
	f.Sync()
	os.Remove(path)
}
