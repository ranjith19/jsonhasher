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
		lhs:         `{"x": 1, "y": "b", "z": null}`,
		rhs:         `{"y": "b", "z": null,   "x":              1}`,
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
	{
		description: "lhs, rhs should match",
		lhs:         "{ \"foo\": 1.23e1, \"bar\": { \"baz\": true, \"abc\": 12 } }",
		rhs:         "{ \"bar\": { \"abc\": 12, \"baz\": true }, \"foo\": 12.3 }",
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

			lhsHash, err = HashJsonStringSha1(tc.lhs)
			require.Nil(t, err)
			rhsHash, err = HashJsonStringSha1(tc.rhs)
			require.Nil(t, err)
			if tc.match {
				require.Equal(t, lhsHash, rhsHash)
			} else {
				require.NotEqual(t, lhsHash, rhsHash)
			}

			lhsHash, err = HashJsonStringSha512(tc.lhs)
			require.Nil(t, err)
			rhsHash, err = HashJsonStringSha512(tc.rhs)
			require.Nil(t, err)
			if tc.match {
				require.Equal(t, lhsHash, rhsHash)
			} else {
				require.NotEqual(t, lhsHash, rhsHash)
			}

			h1, _ := HashJsonStringSha1(tc.lhs)
			h2, _ := HashJsonStringSha256(tc.lhs)
			h3, _ := HashJsonStringSha512(tc.rhs)
			require.NotEqual(t, h1, h2)
			require.NotEqual(t, h1, h3)
			require.NotEqual(t, h2, h3)
		})
	}
}

type hashInterfaceTestCase struct {
	description string
	lhs         interface{}
	rhs         interface{}
	match       bool
}

type dummyStruct struct {
	X string
	Y bool
	Z float64
	i int
}

var hashInterfaceTestCases = []hashInterfaceTestCase{
	{
		description: "lhs, rhs should match",
		lhs:         dummyStruct{"A", false, 10.1, 1},
		rhs:         dummyStruct{"A", false, 10.1, 0},
		match:       true,
	},
}

func TestHashInterface(t *testing.T) {
	for _, tc := range hashInterfaceTestCases {
		t.Run(tc.description, func(t *testing.T) {
			lhsHash, err := HashInterface(tc.lhs)
			require.Nil(t, err)
			rhsHash, err := HashInterface(tc.rhs)
			require.Nil(t, err)
			if tc.match {
				require.Equal(t, lhsHash, rhsHash)
			} else {
				require.NotEqual(t, lhsHash, rhsHash)
			}

			lhsHash, err = HashInterfaceSha1(tc.lhs)
			require.Nil(t, err)
			rhsHash, err = HashInterfaceSha1(tc.rhs)
			require.Nil(t, err)
			if tc.match {
				require.Equal(t, lhsHash, rhsHash)
			} else {
				require.NotEqual(t, lhsHash, rhsHash)
			}

			lhsHash, err = HashInterfaceSha256(tc.lhs)
			require.Nil(t, err)
			rhsHash, err = HashInterfaceSha256(tc.rhs)
			require.Nil(t, err)
			if tc.match {
				require.Equal(t, lhsHash, rhsHash)
			} else {
				require.NotEqual(t, lhsHash, rhsHash)
			}

			lhsHash, err = HashInterfaceSha512(tc.lhs)
			require.Nil(t, err)
			rhsHash, err = HashInterfaceSha512(tc.rhs)
			require.Nil(t, err)
			if tc.match {
				require.Equal(t, lhsHash, rhsHash)
			} else {
				require.NotEqual(t, lhsHash, rhsHash)
			}

			h1, _ := HashInterfaceSha1(tc.lhs)
			h2, _ := HashInterfaceSha256(tc.lhs)
			h3, _ := HashInterfaceSha512(tc.rhs)
			require.NotEqual(t, h1, h2)
			require.NotEqual(t, h1, h3)
			require.NotEqual(t, h2, h3)
		})
	}
}
