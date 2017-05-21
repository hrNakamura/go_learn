package main

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

func Test_sqrtBigFloat(t *testing.T) {
	type args struct {
		v *big.Float
	}
	type testst struct {
		name string
		args args
		want *big.Float
	}
	tests := make([]testst, 0)
	{
		// TODO: Add test cases.
		var st testst
		st.name = "test"
		st.args = args{new(big.Float).SetFloat64(5.0)}
		st.want = new(big.Float).SetFloat64(math.Sqrt(5))
		tests = append(tests, st)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sqrtBigFloat(tt.args.v)
			fmt.Printf("sqrtBigFloat() = %v, want %v\n", got, tt.want)
		})
	}
}
