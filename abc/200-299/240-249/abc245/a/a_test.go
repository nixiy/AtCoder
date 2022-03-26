package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		takaH int
		takaM int
		aokiH int
		aokiM int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{
			takaH: 0, takaM: 0, aokiH: 0, aokiM: 0,
		}, want: true},

		{args: args{
			takaH: 0, takaM: 1, aokiH: 0, aokiM: 0,
		}, want: false},

		{args: args{
			takaH: 0, takaM: 0, aokiH: 0, aokiM: 1,
		}, want: true},

		{args: args{
			takaH: 0, takaM: 0, aokiH: 0, aokiM: 59,
		}, want: true},

		{args: args{
			takaH: 0, takaM: 59, aokiH: 0, aokiM: 0,
		}, want: false},

		{args: args{
			takaH: 0, takaM: 59, aokiH: 1, aokiM: 0,
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.takaH, tt.args.takaM, tt.args.aokiH, tt.args.aokiM); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
