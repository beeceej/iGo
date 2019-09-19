package igoclient

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/beeceej/iGo/igo/igotypes"
)

// Client is an http client which knows how to talk to an igod server
type Client struct {
	*http.Client
	Host string
	Port string
}

// Interpret takes the code and sends it off to igod
// then returns the response
func (c *Client) Interpret(code string) (igotypes.InterpretResponse, error) {
	var (
		b        []byte
		err      error
		h        = http.DefaultClient
		request  = new(igotypes.InterpretRequest)
		response = new(igotypes.InterpretResponse)
	)
	request.Input = code
	if b, err = igotypes.Marshal(request); err != nil {
		log.Fatal(err.Error())
	}
	req, err := http.NewRequest(
		http.MethodPost,
		"http://"+c.Host+":"+c.Port+"/interpret",
		bytes.NewReader(b),
	)
	if err != nil {
		return igotypes.InterpretResponse{}, err
	}
	result, err := h.Do(req)
	if err != nil {
		return igotypes.InterpretResponse{}, err
	}
	defer func() {
		if err = result.Body.Close(); err != nil {
			panic(err.Error())
		}
	}()

	if b, err = ioutil.ReadAll(result.Body); err != nil {
		return igotypes.InterpretResponse{}, err
	}

	if err = igotypes.Unmarshal(b, response); err != nil {
		return igotypes.InterpretResponse{}, err
	}
	return *response, nil
}
