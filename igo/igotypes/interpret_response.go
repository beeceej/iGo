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
	pbresponse := &igopb.InterpretResponse{
		Result: &igopb.Result{
			EvaluatedTo: r.Result.EvaluatedTo,
			Info:        r.Result.Info,
		},
	}
	return proto.Marshal(pbresponse)
}
