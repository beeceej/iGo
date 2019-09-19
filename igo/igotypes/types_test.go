package igotypes

import (
	"testing"

	"github.com/beeceej/iGo/igo/igopb"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	type test struct {
		expect         string
		givenData      []byte
		givenReference interface{}
		expected       interface{}
		errorAssertion assert.ErrorAssertionFunc
	}

	cases := map[string]test{
		`given nil byte slice`: {
			expect:         "an error",
			givenData:      func() (b []byte) { return b }(),
			givenReference: nil,
			expected:       nil,
			errorAssertion: assert.Error,
		},
		`given empty byte slice`: {
			expect:         "an error",
			givenData:      []byte{},
			givenReference: nil,
			expected:       nil,
			errorAssertion: assert.Error,
		},
		`given a reference to an unaccounted for type`: {
			expect:         "an error",
			givenData:      []byte{},
			givenReference: new(struct{ A string }),
			expected:       new(struct{ A string }),
			errorAssertion: assert.Error,
		},
		`given a reference to InterpreterRequest and invalid bytes`: {
			expect:         "an error",
			givenData:      []byte{0, 1, 2, 3, 4},
			givenReference: new(InterpretRequest),
			expected:       new(InterpretRequest),
			errorAssertion: assert.Error,
		},
		`given a reference to InterpreterResponse and invalid bytes`: {
			expect:         "an error",
			givenData:      []byte{0, 1, 2, 3, 4},
			givenReference: new(InterpretResponse),
			expected:       new(InterpretResponse),
			errorAssertion: assert.Error,
		},
		`given a reference to InterpreterRequest and valid bytes`: {
			expect: "InterpretRequest unmarshalled correctly",
			givenData: func() []byte {
				req := new(igopb.InterpretRequest)
				req.Input = "asdf"
				b, err := proto.Marshal(req)
				assert.NoError(t, err)
				return b
			}(),
			givenReference: new(InterpretRequest),
			expected: &InterpretRequest{
				Input: "asdf",
			},
			errorAssertion: assert.NoError,
		},
		`given a reference to InterpreterResponse and valid bytes`: {
			expect: "InterpretRequest unmarshalled correctly",
			givenData: func() []byte {
				res := new(igopb.InterpretResponse)
				res.Result = &igopb.Result{
					EvaluatedTo: "123",
					Info:        "some info",
				}
				b, err := proto.Marshal(res)
				assert.NoError(t, err)
				return b
			}(),
			givenReference: new(InterpretResponse),
			expected: &InterpretResponse{
				Result: InterpretResult{
					EvaluatedTo: "123",
					Info:        "some info",
				},
			},
			errorAssertion: assert.NoError,
		},
	}

	for description, c := range cases {
		t.Run(description, func(t *testing.T) {
			err := Unmarshal(c.givenData, c.givenReference)
			c.errorAssertion(t, err, c.expect)
			assert.Equal(t, c.expected, c.givenReference, c.expect)
		})
	}
}

func TestMarshal(t *testing.T) {
	type test struct {
		expect         string
		expectedData   []byte
		givenReference interface{}
		errorAssertion assert.ErrorAssertionFunc
	}

	cases := map[string]test{
		`given nil byte slice`: {
			expect:         "an error",
			expectedData:   func() (b []byte) { return b }(),
			givenReference: nil,
			errorAssertion: assert.Error,
		},
		`given empty byte slice`: {
			expect:         "an error",
			expectedData:   func() (b []byte) { return b }(),
			givenReference: nil,
			errorAssertion: assert.Error,
		},
		`given a reference to an unaccounted for type`: {
			expect:         "an error",
			expectedData:   nil,
			givenReference: new(struct{ A string }),
			errorAssertion: assert.Error,
		},
		`given an empty reference to InterpreterRequest`: {
			expect:         "empty pb message",
			givenReference: &InterpretRequest{},
			expectedData: func() []byte {
				data, err := proto.Marshal(&igopb.InterpretRequest{})
				assert.NoError(t, err)
				return data
			}(),
			errorAssertion: assert.NoError,
		},
		`given an empty reference to InterpretResponse`: {
			expect:         "empty pb message",
			givenReference: &InterpretResponse{},
			expectedData: func() []byte {
				data, err := proto.Marshal(&igopb.InterpretResponse{
					Result: &igopb.Result{},
				})
				assert.NoError(t, err)
				return data
			}(),
			errorAssertion: assert.NoError,
		},
		`given a reference to InterpreterRequest`: {
			expect: "expect correct proto bytes",
			givenReference: &InterpretRequest{
				Input: "some input",
			},
			expectedData: func() []byte {
				data, err := proto.Marshal(
					&igopb.InterpretRequest{
						Input: "some input",
					},
				)
				assert.NoError(t, err)
				return data
			}(),
			errorAssertion: assert.NoError,
		},
		`given a reference to InterpreterResponse`: {
			expect: "expect correct proto bytes",
			givenReference: &InterpretResponse{
				Result: InterpretResult{
					EvaluatedTo: "asdf",
					Info:        "some info",
				},
			},
			expectedData: func() []byte {
				data, err := proto.Marshal(
					&igopb.InterpretResponse{
						Result: &igopb.Result{
							EvaluatedTo: "asdf",
							Info:        "some info",
						},
					},
				)
				assert.NoError(t, err)
				return data
			}(),
			errorAssertion: assert.NoError,
		},
	}

	for description, c := range cases {
		t.Run(description, func(t *testing.T) {
			actualData, err := Marshal(c.givenReference)
			c.errorAssertion(t, err, c.expect)
			assert.Equal(t, c.expectedData, actualData, c.expect)
		})
	}
}
