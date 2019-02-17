package igo

import (
	"github.com/beeceej/iGo/igopb"
	"github.com/golang/protobuf/proto"
)

type InterpretRequest struct {
	Input string
}

func (i *InterpretRequest) FromProtoBytes(b []byte) (err error) {
	p := new(igopb.InterpretRequest)
	if err = proto.Unmarshal(b, p); err != nil {
		return err
	}

	i.Input = p.GetIn()

	return nil
}

type InterpretResponse struct {
	Result *InterpretResult
}

type InterpretResult struct {
	EvaluatedTo string
	Info        string
}

func (i *InterpretResponse) FromProtoBytes(b []byte) (err error) {
	p := new(igopb.InterpretResponse)
	if err = proto.Unmarshal(b, p); err != nil {
		return err
	}

	interpretResult := new(InterpretResult)
	interpretResult.Info = p.GetResult().GetInfo()
	interpretResult.EvaluatedTo = p.GetResult().GetEvaluatedTo()
	i.Result = interpretResult
	return nil
}

func (i *InterpretResponse) ToProtoBytes() (b []byte, err error) {
	p := new(igopb.InterpretResponse)
	p.Result = new(igopb.Result)
	p.Result.Info = i.Result.Info
	p.Result.EvaluatedTo = i.Result.EvaluatedTo
	return proto.Marshal(p)
}
