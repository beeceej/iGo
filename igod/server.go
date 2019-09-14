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
	*http.ServeMux
	Interpreter interpreter.Interpreter
}

// Run is
func (s *Server) Run() {
	s.HandleFunc("/interpret", func(w http.ResponseWriter, r *http.Request) {
		var (
			b   []byte
			err error
		)

		if b, err = ioutil.ReadAll(r.Body); err != nil {
			log.Fatalln(err.Error())
		}

		defer func() {
			if err = r.Body.Close(); err != nil {
				panic(err.Error())
			}
		}()

		interpretRequest := new(igo.InterpretRequest)
		if err := interpretRequest.FromProtoBytes(b); err != nil {
			log.Fatalln(err.Error())
		}
		result := s.Interpreter.Interpret(interpretRequest.Input)
		res := new(igo.InterpretResponse)
		res.Result = new(igo.InterpretResult)
		res.Result.EvaluatedTo = "EvaluatedTo: " + result + "\n"
		res.Result.Info = fmt.Sprintf("INFO: %s\n", result)
		if b, err = res.ToProtoBytes(); err != nil {
			log.Fatalln(err.Error())
		}
		w.Write(b)
	})

	if err := http.ListenAndServe(":9999", s); err != nil {
		log.Fatalln(err.Error())
	}
}
