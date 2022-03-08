package main

import (
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		N   int
		W   int
		ckv []kv
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{args: args{
			N: 3,
			W: 5,
			ckv: []kv{
				{3, 1},
				{4, 2},
				{2, 4},
			},
		}, wantSum: 15},

		{args: args{
			N: 4,
			W: 100,
			ckv: []kv{
				{6, 2},
				{1, 5},
				{3, 9},
				{8, 7},
			},
		}, wantSum: 100},

		{args: args{
			N: 10,
			W: 3141,
			ckv: []kv{
				{314944731, 649},
				{140276783, 228},
				{578012421, 809},
				{878510647, 519},
				{925326537, 943},
				{337666726, 611},
				{879137070, 306},
				{87808915, 39},
				{756059990, 244},
				{228622672, 291},
			},
		}, wantSum: 2357689932073},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := solve(tt.args.N, tt.args.W, tt.args.ckv); gotSum != tt.wantSum {
				t.Errorf("solve() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
