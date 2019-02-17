package igod

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	igo "github.com/beeceej/iGo"
	"github.com/beeceej/iGo/interpreter"
)

// Server is
type Server struct {
	Interpreter interpreter.Interpreter
	*http.ServeMux
}

// Run is
func (s *Server) Run() {
	s.HandleFunc("/interpret", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer r.Body.Close()
		interpretReq := new(igo.InterpretRequest)
		interpretReq.FromProtoBytes(b)
		if err := interpretReq.FromProtoBytes(b); err != nil {
			log.Fatalln(err.Error())
		}
		result := s.Interpreter.Interpret(interpretReq.Input)
		res := new(igo.InterpretResponse)
		res.Result = new(igo.InterpretResult)
		res.Result.EvaluatedTo = "EvaluatedTo: " + result + "\n"
		res.Result.Info = "INFO: " + result + "\n"
		if b, err = res.ToProtoBytes(); err != nil {
			log.Fatalln(err.Error())
		}

		w.Write(b)
	})
	s.HandleFunc("/inspect", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit Inspect")
	})
	if err := http.ListenAndServe(":9999", s); err != nil {
		log.Fatalln(err.Error())
	}
}
