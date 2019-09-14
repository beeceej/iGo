package igod

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/beeceej/iGo/igo/igotypes"
	"github.com/beeceej/iGo/igo/interpreter"
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

		interpretRequest := new(igotypes.InterpretRequest)
		if err := igotypes.Unmarshall(b, interpretRequest); err != nil {
			log.Fatalln(err.Error())
		}
		result := s.Interpreter.Interpret(interpretRequest.Input)
		response := &igotypes.InterpretResponse{
			Result: igotypes.InterpretResult{
				EvaluatedTo: "EvaluatedTo: " + result + "\n",
				Info:        fmt.Sprintf("INFO: %s\n", result),
			},
		}
		if b, err = igotypes.Marshall(response); err != nil {
			log.Fatalln(err.Error())
		}
		w.Write(b)
	})

	if err := http.ListenAndServe(":9999", s); err != nil {
		log.Fatalln(err.Error())
	}
}
