package igoclient

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	igo "github.com/beeceej/iGo"
	"github.com/beeceej/iGo/igopb"
	"github.com/golang/protobuf/proto"
)

// Client is an http client which knows how to talk to an igod server
type Client struct {
	*http.Client
	Host string
	Port string
}

// Interpret takes the code and sends it off to igod
// then returns the response
func (c *Client) Interpret(code string) (string, string) {
	var (
		b       []byte
		err     error
		h       = http.DefaultClient
		request = new(igopb.InterpretRequest)
	)
	request.In = code
	if b, err = proto.Marshal(request); err != nil {
		log.Fatal(err.Error())
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"http://"+c.Host+":"+c.Port+"/interpret",
		bytes.NewReader(b),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	result, err := h.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		if err = result.Body.Close(); err != nil {
			panic(err.Error())
		}
	}()

	b, err = ioutil.ReadAll(result.Body)
	response := new(igo.InterpretResponse)
	if err = response.FromProtoBytes(b); err != nil {
		log.Fatal(err.Error())
	}
	return response.Result.Info, response.Result.EvaluatedTo
}
