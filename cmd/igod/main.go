package main

import (
	"net/http"

	"github.com/beeceej/iGo/igod"
	"github.com/beeceej/iGo/interpreter"
)

func main() {
	i := interpreter.Interpreter{}
	s := &igod.Server{
		ServeMux:    http.NewServeMux(),
		Interpreter: i,
	}
	s.Run()
}
