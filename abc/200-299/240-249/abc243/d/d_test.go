package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		X int
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{X: 2, S: "URL"}, want: 6},
		{args: args{X: 500000000000000000, S: "RRUU"}, want: 500000000000000000},
		{args: args{X: 123456789, S: "LRULURLURLULULRURRLRULRRRUURRU"}, want: 126419752371},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.X, tt.args.S); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
