package column

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractAggregateFunctionParams(t *testing.T) {
	testCases := []struct {
		name           string
		params         string
		expectedFunc   string
		expectedBase   string
	}{
		{
			name:         "simple case",
			params:       "sum, Float64",
			expectedFunc: "sum",
			expectedBase: "Float64",
		},
		{
			name:         "with spaces",
			params:       "argMin, Float64, UInt64",
			expectedFunc: "argMin",
			expectedBase: "Float64",
		},
		{
			name:         "with tuple",
			params:       "groupArray, Tuple(String, String)",
			expectedFunc: "groupArray",
			expectedBase: "Tuple(String, String)",
		},
		{
			name:         "complex with tuple and extra param",
			params:       "groupArray, Tuple(String, String), UInt32",
			expectedFunc: "groupArray",
			expectedBase: "Tuple(String, String)",
		},
		{
			name:         "nested tuples",
			params:       "groupArray, Tuple(Tuple(UInt16, UInt16), Tuple(UInt16, UInt16))",
			expectedFunc: "groupArray",
			expectedBase: "Tuple(Tuple(UInt16, UInt16), Tuple(UInt16, UInt16))",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			funcName, baseType := ExtractAggregateFunctionParams(tc.params)
			assert.Equal(t, tc.expectedFunc, funcName)
			assert.Equal(t, tc.expectedBase, baseType)
		})
	}
}
