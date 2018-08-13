package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beeceej/iGo/pkg/interpreter"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	port := flag.String("port", "9999", "Port is the port the interpreter will run on")
	i := interpreter.Interpreter{}
	http.HandleFunc("/interpret", func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		var m struct {
			Raw string `json:"raw"`
		}
		json.Unmarshal(b, &m)
		i.Interpret(m.Raw)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%s", *port), nil); err != nil {
		fmt.Println(err.Error())
		spew.Dump(i)
	}
}
