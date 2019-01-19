package interpreter

// Eval is largely a POC at this point,
// What you see here is hacking together a concept until it worked, and no more.
// Though it works (yay), a large refactoring is to come.

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

var mainTemplate = template.Must(template.New("_").Parse(`package main

func main() {
	{{ .M }}
}

{{range .F }}
	{{.}}
{{end}}
`))

// EvalData is
type EvalData struct {
	M string
	F []string
}

const path = "/tmp/igo/main.go"

// Eval will evaluate the text
func (i *Interpreter) Eval(text string) {
	var ed *EvalData

	if !pathExists(path) {
		dir, _ := filepath.Split(path)
		os.MkdirAll(dir, os.ModePerm)
	}
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

	if err := mainTemplate.Execute(f, ed); err != nil {
		fmt.Println(err.Error())
	}

	cmd := exec.Command("goimports", "-w", path)
	b, err := cmd.Output()

	if err != nil {
		fmt.Println("Error calling goimports", err.Error())
		fmt.Println(string(b))
		return
	}
	f.Sync()

	cmd = exec.Command("go", "build", "-o", path, path)
	b, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(b))

		lookup := strings.TrimFunc(ed.M, func(r rune) bool {
			return r == ')' || r == '('
		})
		fmt.Println(i.Functions[lookup].Raw)
		return
	}

	cmd = exec.Command(path)
	b, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing", err.Error())
	}
	fmt.Println(string(b))
	f.Sync()
}

func pathExists(path string) bool {
	pathPart, _ := filepath.Split(path)
	info, err := ioutil.ReadDir(pathPart)

	if err != nil {
		return false
	}
	return len(info) > 0
}
