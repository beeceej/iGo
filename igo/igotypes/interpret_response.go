package igotypes

import (
	"github.com/beeceej/iGo/igo/igopb"
	"github.com/golang/protobuf/proto"
)

type (

	// InterpretResponse represents the response back from igod
	InterpretResponse struct {
		Result InterpretResult
	}

	// InterpretResult represents the result of the response back grom igod
	InterpretResult struct {
		EvaluatedTo string
		Info        string
	}
)

func unmarshalInterpretResponse(data []byte, r *InterpretResponse) error {
	irpb := new(igopb.InterpretResponse)
	err := proto.Unmarshal(data, irpb)
	if err != nil {
		return err
	}
	pbresult := irpb.GetResult()
	r.Result = InterpretResult{
		EvaluatedTo: pbresult.GetEvaluatedTo(),
		Info:        pbresult.GetInfo(),
	}
	return nil
}

func marshalInterpretResponse(r *InterpretResponse) ([]byte, error) {
	pbresponse := new(igopb.InterpretResponse)
	pbresult := new(igopb.Result)
	pbresult.EvaluatedTo = r.Result.EvaluatedTo
	pbresult.Info = r.Result.Info
	pbresponse.Result = pbresult
	return proto.Marshal(pbresponse)
}
