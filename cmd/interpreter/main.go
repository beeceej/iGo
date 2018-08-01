package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beeceej/iGo/pkg/interpreter"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	i := interpreter.Interpreter{}
	http.HandleFunc("/interpret", func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		var m struct {
			Text string `json:"text"`
		}
		json.Unmarshal(b, &m)
		i.Interpret(m.Text)
	})

	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Println(err.Error())
		spew.Dump(i)
	}
}
