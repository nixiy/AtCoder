package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		N      int
		Coupon int
		X      int
		A      []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{
			N:      5,
			Coupon: 4,
			X:      7,
			A:      []int{8, 3, 10, 5, 13},
		}, want: 12},

		{args: args{
			N:      5,
			Coupon: 100,
			X:      7,
			A:      []int{8, 3, 10, 5, 13},
		}, want: 0},

		{args: args{
			N:      20,
			Coupon: 815,
			X:      60,
			A:      []int{2066, 3193, 2325, 4030, 3725, 1669, 1969, 763, 1653, 159, 5311, 5341, 4671, 2374, 4513, 285, 810, 742, 2981, 202},
		}, want: 112},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.N, tt.args.Coupon, tt.args.X, tt.args.A); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
