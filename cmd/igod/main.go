package main

import (
	"net/http"

	"github.com/beeceej/iGo/igod"
	"github.com/beeceej/iGo/interpreter"
)

func main() {
	// port := flag.String("port", "9999", "Port is the port the interpreter will run on")
	i := interpreter.Interpreter{}
	s := igod.Server{
		ServeMux:    http.NewServeMux(),
		Interpreter: i,
	}
	s.Run()
	// http.HandleFunc("/interpret", func(w http.ResponseWriter, r *http.Request) {
	// 	b, _ := ioutil.ReadAll(r.Body)
	// 	defer r.Body.Close()
	// 	var m struct {
	// 		Raw string `json:"raw"`
	// 	}
	// 	json.Unmarshal(b, &m)
	// 	result := i.Interpret(m.Raw)
	// 	rm := map[string]string{}
	// 	rm["raw"] = result
	// 	b, _ = json.Marshal(rm)
	// 	w.Write(b)
	// }
	// )

	// if err := http.ListenAndServe(fmt.Sprintf(":%s", *port), nil); err != nil {
	// 	fmt.Println(err.Error())
	// }
}
