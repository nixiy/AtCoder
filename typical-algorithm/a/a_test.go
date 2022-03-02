package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		N int
		K int
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{
			N: 8,
			K: 4,
			A: []int{1, 3, 5, 7, 9, 11, 13, 15},
		}, want: 2},
		{args: args{
			N: 5,
			K: 1000000000,
			A: []int{1, 2, 3, 4, 5},
		}, want: -1},
		{args: args{
			N: 8,
			K: 10,
			A: []int{0, 5, 10, 10, 10, 11, 13, 17},
		}, want: 2},
		{args: args{
			N: 8,
			K: 2,
			A: []int{3, 5, 10, 10, 10, 11, 13, 17},
		}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.N, tt.args.K, tt.args.A); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
