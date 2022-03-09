package main

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		N_nin  int
		K_joui int
		Point  [][3]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{args: args{
			N_nin:  3,
			K_joui: 1,
			Point: [][3]int{
				{178, 205, 132},
				{112, 220, 96},
				{36, 64, 20},
			},
		}, want: []bool{true, true, false}},

		{args: args{
			N_nin:  2,
			K_joui: 1,
			Point: [][3]int{
				{300, 300, 300},
				{200, 200, 200},
			},
		}, want: []bool{true, true}},

		{args: args{
			N_nin:  4,
			K_joui: 2,
			Point: [][3]int{
				{127, 235, 78},
				{192, 134, 298},
				{28, 56, 42},
				{96, 120, 250},
			},
		}, want: []bool{true, true, false, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.N_nin, tt.args.K_joui, tt.args.Point); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
