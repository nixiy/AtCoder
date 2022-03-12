package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		S int
		T int
		X int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{7, 20, 12}, want: true},
		{args: args{20, 7, 12}, want: false},
		{args: args{23, 0, 23}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.S, tt.args.T, tt.args.X); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
