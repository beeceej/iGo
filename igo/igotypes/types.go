package igotypes

import (
	"fmt"
)

// Unmarshall is a function which attempts to unmarshal
// protobuf bytes into its go type
func Unmarshall(data []byte, i interface{}) error {
	switch t := i.(type) {
	case *InterpretRequest:
		unmarshalInterpretRequest(data, t)
	case *InterpretResponse:
		unmarshalInterpretResponse(data, t)
	default:
		return fmt.Errorf("unable to unmarshall %v", i)
	}
	return nil
}

// Marshall is a function which attempts to marshal
// an igo type into it's protobuf byte form
func Marshall(i interface{}) ([]byte, error) {
	switch t := i.(type) {
	case *InterpretRequest:
		return marshalInterpretRequest(t)
	case *InterpretResponse:
		return marshalInterpretResponse(t)
	default:
		return nil, fmt.Errorf("unable to marshall %v", i)
	}
}
