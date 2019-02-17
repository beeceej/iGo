package igoclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	igo "github.com/beeceej/iGo"
	"github.com/beeceej/iGo/igopb"
	"github.com/golang/protobuf/proto"
)

type Client struct {
	http.Client
}

func Test() {
	h := http.DefaultClient
	_ = igo.InterpretRequest{}
	sampleBody := new(igopb.InterpretRequest)
	sampleBody.In = `func hello() string {
    return "Hello"
}`
	b, err := proto.Marshal(sampleBody)
	if err != nil {
		log.Fatal(err.Error())
	}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9999/interpret", bytes.NewReader(b))
	if err != nil {
		log.Fatal(err.Error())
	}
	result, err := h.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer result.Body.Close()
	b, err = ioutil.ReadAll(result.Body)
	response := new(igo.InterpretResponse)
	err = response.FromProtoBytes(b)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(response.Result.Info, response.Result.EvaluatedTo)
}
