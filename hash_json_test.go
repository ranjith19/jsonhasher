package jsonhasher

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type hashJsonTestCase struct {
	description string
	lhs         string
	rhs         string
	match       bool
}

var hashJsonTestCases = []hashJsonTestCase{
	{
		description: "lhs, rhs should match",
		lhs:         `{"x": "a", "y": "b"}`,
		rhs:         `{"y": "b", "x": "a"}`,
		match:       true,
	},
	{
		description: "lhs, rhs should match",
		lhs:         `{"x": "a", "y": "b"}`,
		rhs:         `{"z": "b", "x": "a"}`,
		match:       false,
	},
	{
		description: "lhs, rhs should match",
		lhs:         `[{"x": "a", "y": "b"}]`,
		rhs:         `[{"y": "b", "x": "a"}]`,
		match:       true,
	},
	{
		description: "lhs, rhs should match",
		lhs:         `[{"x": false, "y": "b"}]`,
		rhs:         `[{"y": "b", "x":    false}]`,
		match:       true,
	},
	{
		description: "lhs, rhs should match",
		lhs:         `[{"x": "a", "y": 1}]`,
		rhs:         `[{"y": 1, "x": "a"}]`,
		match:       true,
	},
	{
		description: "lhs, rhs should match",
		lhs:         `[1,2,3,4,5]`,
		rhs:         `[1,2,3,4,5]`,
		match:       true,
	},
	{
		description: "lhs, rhs should match",
		lhs:         `1.23e1`,
		rhs:         `12.3`,
		match:       true,
	},
}

func TestHashJsonString(t *testing.T) {
	for _, tc := range hashJsonTestCases {
		t.Run(tc.description, func(t *testing.T) {
			lhsHash, err := HashJsonString(tc.lhs)
			require.Nil(t, err)
			rhsHash, err := HashJsonString(tc.rhs)
			require.Nil(t, err)
			if tc.match {
				require.Equal(t, lhsHash, rhsHash)
			} else {
				require.NotEqual(t, lhsHash, rhsHash)
			}
		})
	}
}

type typeDetermineTestCase struct {
	js       string
	expected string
}

var typeDetermineTestCases = []typeDetermineTestCase{
	{
		js:       "null",
		expected: "nil",
	},
	{
		js:       `"haha"`,
		expected: "string",
	},
	{
		js:       `11`,
		expected: "float",
	},
	{
		js:       `11.0`,
		expected: "float",
	},
	{
		js:       `{"haha": 111}`,
		expected: "dict",
	},
	{
		js:       `[{"haha": 111}]`,
		expected: "list",
	},
	{
		js:       `1.23e1`,
		expected: "float",
	},
}

func TestTypeCheck(t *testing.T) {
	for _, tc := range typeDetermineTestCases {
		t.Run(tc.expected, func(t *testing.T) {
			determined, err := determineType([]byte(tc.js))
			require.Nil(t, err)
			require.Equal(t, tc.expected, *determined)
		})
	}
}
