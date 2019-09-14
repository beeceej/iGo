package igotypes

import (
	"github.com/beeceej/iGo/igo/igopb"
	"github.com/golang/protobuf/proto"
)

// InterpretRequest represents the network request sent to the igod
type InterpretRequest struct {
	Input string
}

func unmarshalInterpretRequest(data []byte, r *InterpretRequest) error {
	irpb := new(igopb.InterpretRequest)
	err := proto.Unmarshal(data, irpb)
	if err != nil {
		return err
	}
	r.Input = irpb.GetInput()
	return nil
}

func marshalInterpretRequest(r *InterpretRequest) ([]byte, error) {
	pbreq := new(igopb.InterpretRequest)
	pbreq.Input = r.Input
	return proto.Marshal(pbreq)
}
