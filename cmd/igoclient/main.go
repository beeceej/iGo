package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/beeceej/iGo/igoclient"
)

type (
	host string
	port string
)

func main() {
	var (
		h, p = parseFlags()
		args = os.Args[1:]
		code string
	)
	if len(args) > 0 {
		code = args[0]
	}

	a, b := (&igoclient.Client{
		Client: http.DefaultClient,
		Host:   string(h),
		Port:   string(p),
	}).Interpret(code)
	fmt.Println(a, b)
}

func parseFlags() (host, port) {
	h := flag.String("host", "", "the host igod is running on")
	p := flag.String("port", "", "the port igod is running on")
	flag.Parse()
	if *h == "" {
		*h = "localhost"
	}
	if *p == "" {
		*p = "9999"
	}
	return host(*h), port(*p)
}
