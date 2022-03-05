package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		a int
		b int
		c int
		x int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "必ずシャツがもらえる範囲を超えているので、確率を計算する", args: args{
			a: 30,
			b: 500,
			c: 20,
			x: 103,
		}, want: 0.0425531914893617},

		{name: "必ずもらえる順位以上であるため、確実にもらえる", args: args{
			a: 50,
			b: 500,
			c: 100,
			x: 1,
		}, want: 1.0},

		{name: "b位以下であるためもらえない", args: args{
			a: 1,
			b: 2,
			c: 1,
			x: 1000,
		}, want: 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.a, tt.args.b, tt.args.c, tt.args.x); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
