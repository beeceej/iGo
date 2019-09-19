package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/beeceej/iGo/igo/igod"
	"github.com/beeceej/iGo/igo/interpreter"
)

type configuration struct {
	Port string `toml:port`
}

func main() {
	i := interpreter.Interpreter{}
	s := &igod.Server{
		ServeMux:    http.NewServeMux(),
		Interpreter: i,
		Port:        getPort(),
	}
	s.Run()
}

//getPort returns the port based on the server configuration
func getPort() string {
	portPtr := flag.String("port", "", "port on which server should run")
	flag.Parse()
	var port string
	if *portPtr != "" {
		port = *portPtr
	} else if os.Getenv("IGO_PORT") != "" {
		port = os.Getenv("IGO_PORT")
	} else {
		var configuration configuration
		if _, err := toml.DecodeFile(filepath.Join(os.Getenv("HOME"), ".config", "igo", "config.toml"), &configuration); err != nil {
			log.Fatal(err)
		}
		port = configuration.Port
	}
	return port
}
