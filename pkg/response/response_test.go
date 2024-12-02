package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseMarshaling(t *testing.T) {
	tests := []struct {
		name           string
		input          Response
		expectedOutput string
	}{
		{
			name: "Complete Response with Error",
			input: Response{
				Status: "fail",
				Message: Message{
					Endpoint: "/test-endpoint",
					Value:    false,
				},
				Error: "An error occurred",
			},
			expectedOutput: `{"status":"fail","message":{"endpoint":"/test-endpoint","value":false},"error":"An error occurred"}`,
		},
		{
			name: "Response Without Error",
			input: Response{
				Status: "success",
				Message: Message{
					Endpoint: "/test-endpoint",
					Value:    true,
				},
			},
			expectedOutput: `{"status":"success","message":{"endpoint":"/test-endpoint","value":true}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.input)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expectedOutput, string(data))
		})
	}
}

func TestResponseUnmarshaling(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput Response
	}{
		{
			name:  "Complete Response with Error",
			input: `{"status":"fail","message":{"endpoint":"/test-endpoint","value":false},"error":"An error occurred"}`,
			expectedOutput: Response{
				Status: "fail",
				Message: Message{
					Endpoint: "/test-endpoint",
					Value:    false,
				},
				Error: "An error occurred",
			},
		},
		{
			name:  "Response Without Error",
			input: `{"status":"success","message":{"endpoint":"/test-endpoint","value":true}}`,
			expectedOutput: Response{
				Status: "success",
				Message: Message{
					Endpoint: "/test-endpoint",
					Value:    true,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output Response
			err := json.Unmarshal([]byte(tt.input), &output)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}
