package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beeceej/iGo/igod/igodpb"
	"github.com/beeceej/iGo/interpreter"
)

func main() {
	_ = igodpb.DebugRequest{}
	port := flag.String("port", "9999", "Port is the port the interpreter will run on")
	i := interpreter.Interpreter{}
	http.HandleFunc("/interpret", func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		var m struct {
			Raw string `json:"raw"`
		}
		json.Unmarshal(b, &m)
		result := i.Interpret(m.Raw)
		rm := map[string]string{}
		rm["raw"] = result
		b, _ = json.Marshal(rm)
		w.Write(b)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%s", *port), nil); err != nil {
		fmt.Println(err.Error())
	}
}
