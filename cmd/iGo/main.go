package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"text/template"

	iGo "github.com/beeceej/iGo/pkg/parse"
)

const mainTmpl = `package main
func main() {
	{{ .M }}
}

{{range .F }}
	{{.}}
{{end}}
`

// In is
type In struct {
	M string
	F []string
}

const path = "/Users/brianjones/development/golib/src/github.com/beeceej/iGo/cmd/iGoBin/exe.go"

func main() {
	t := template.Must(template.New("mainTmpl").Parse(mainTmpl))
	history := make(map[string]string)
	var instruct In
	for {
		fmt.Print("\n$ ")
		f, err := os.Create("/Users/brianjones/development/golib/src/github.com/beeceej/iGo/cmd/iGoBin/exe.go")
		if err != nil {
			panic(err.Error())
		}
		defer f.Close()
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		for _, v := range history {
			var found bool
			for _, fns := range instruct.F {
				if fns == v {
					found = true
				}
			}
			if !found {
				instruct.F = append(instruct.F, v)
			}
		}

		instruct.M = text

		if err := t.Execute(f, instruct); err != nil {
			fmt.Println(err.Error())
		}

		if fn, err := iGo.NewFunction(text); err != nil {
			cmd := exec.Command("/Users/brianjones/development/golib/bin/goimports", "-w", "/Users/brianjones/development/golib/src/github.com/beeceej/iGo/cmd/iGoBin/exe.go")
			b, err := cmd.Output()

			if err != nil {
				fmt.Println("Error calling goimports", err.Error())
				continue
			}

			fmt.Println(string(b))
			f.Sync()
			cmd = exec.Command("go", "run", "/Users/brianjones/development/golib/src/github.com/beeceej/iGo/cmd/iGoBin/exe.go")
			b, err = cmd.Output()
			if err != nil {
				fmt.Println("Error calling go run", err.Error())
				continue
			}
			fmt.Println(fmt.Sprintf(">> %s", string(b)))
			f.Sync()
			os.Remove(path)
		} else {
			history[fn.Identifier] = fn.Raw
			fmt.Printf("# %s :: Function (%s) -> %s\n", fn.Identifier, fn.Args, fn.Return)
		}
	}
}
