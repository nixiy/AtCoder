package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		N       int
		X       int
		friends []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{
			N:       4,
			X:       2,
			friends: []int{3, 1, 1, 2},
		}, want: 3},

		{args: args{
			N:       20,
			X:       12,
			friends: []int{7, 11, 10, 1, 7, 20, 14, 2, 17, 3, 2, 5, 19, 20, 8, 14, 18, 2, 10, 10},
		}, want: 7},

		{args: args{
			N:       3,
			X:       1,
			friends: []int{2, 3, 1},
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.N, tt.args.X, tt.args.friends); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
