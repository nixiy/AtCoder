package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		N int   // 足場の数
		K int   // 何個先まで飛べるか
		h []int // 各足場の高さ
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "足場 1 → 2 → 5 と移動すると、コストの総和は ∣10−30∣+∣30−20∣=30 となります。", args: args{
			N: 5,
			K: 3,
			h: []int{10, 30, 40, 50, 20},
		}, want: 30},
		{name: "足場 1 → 2 → 3 と移動すると、コストの総和は ∣10−20∣+∣20−10∣=20 となります。", args: args{
			N: 3,
			K: 1,
			h: []int{10, 20, 10},
		}, want: 20},
		{name: "足場 1 → 2 と移動すると、コストの総和は ∣10−10∣=0 となります。", args: args{
			N: 2,
			K: 100,
			h: []int{10, 10},
		}, want: 0},
		{name: "足場 1 → 4 → 8 → 10 と移動すると、コストの総和は ∣40−70∣+∣70−70∣+∣70−60∣=40 となります。", args: args{
			N: 10,
			K: 4,
			h: []int{40, 10, 20, 70, 80, 10, 20, 70, 80, 60},
		}, want: 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.N, tt.args.K, tt.args.h); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
