package igotypes

import (
	"fmt"
)

// Unmarshal is a function which attempts to unmarshal
// protobuf bytes into its go type
func Unmarshal(data []byte, i interface{}) (err error) {
	switch t := i.(type) {
	case *InterpretRequest:
		err = unmarshalInterpretRequest(data, t)
	case *InterpretResponse:
		err = unmarshalInterpretResponse(data, t)
	default:
		err = fmt.Errorf("unable to unmarshal %v", i)
	}
	return err
}

// Marshal is a function which attempts to marshal
// an igo type into it's protobuf byte form
func Marshal(i interface{}) ([]byte, error) {
	switch t := i.(type) {
	case *InterpretRequest:
		return marshalInterpretRequest(t)
	case *InterpretResponse:
		return marshalInterpretResponse(t)
	default:
		return nil, fmt.Errorf("unable to marshall %v", i)
	}
}
