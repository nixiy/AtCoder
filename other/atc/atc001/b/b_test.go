package main

import (
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		N       int
		Q       int
		queries []Query
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{
			N: 8,
			Q: 9,
			queries: []Query{
				{0, 1, 2},
				{0, 3, 2},
				{1, 1, 3},
				{1, 1, 4},
				{0, 2, 4},
				{1, 4, 1},
				{0, 4, 2},
				{0, 0, 0},
				{1, 0, 0},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.N, tt.args.Q, tt.args.queries)
		})
	}
}
