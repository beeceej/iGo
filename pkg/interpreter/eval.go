package interpreter

// Eval is largely a POC at this point,
// What you see here is hacking together a concept until it worked, and no more.
// Though it works (yay), a large refactoring is to come.

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

var path = filepath.Join(os.Getenv("PWD"), "cmd/iGoBin/exe.go")

// Eval will evaluate the text
func (i *Interpreter) Eval(text string) {
	var ed *EvalData
	t := template.Must(template.New("mainTmpl").Parse(tmpl))
	f, err := os.Create(path)
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

	cmd := exec.Command(filepath.Join("goimports"), "-w", path)
	_, err = cmd.Output()

	if err != nil {
		fmt.Println("Error calling goimports", err.Error())
		return
	}
	f.Sync()

	cmd = exec.Command("go", "run", path)
	b, err := cmd.Output()
	if err != nil {
		fmt.Println("Error calling go run", err.Error())
	}

	fmt.Println(fmt.Sprintf(">> %s", string(b)))
	f.Sync()
	os.Remove(path)
}
