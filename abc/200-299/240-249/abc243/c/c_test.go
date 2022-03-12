package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		N           int
		coordinates []Coordinate
		LRs         string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{
			N: 3,
			coordinates: []Coordinate{
				{x: 2, y: 3},
				{x: 1, y: 1},
				{x: 4, y: 1},
			},
			LRs: "RRL",
		}, want: true},

		{args: args{
			N: 2,
			coordinates: []Coordinate{
				{x: 1, y: 1},
				{x: 2, y: 1},
			},
			LRs: "RR",
		}, want: false},

		{args: args{
			N: 10,
			coordinates: []Coordinate{
				{x: 1, y: 3},
				{x: 1, y: 4},
				{x: 0, y: 0},
				{x: 0, y: 2},
				{x: 0, y: 4},
				{x: 3, y: 1},
				{x: 2, y: 4},
				{x: 4, y: 2},
				{x: 4, y: 4},
				{x: 3, y: 3},
			},
			LRs: "RLRRRLRLRR",
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.N, tt.args.coordinates, tt.args.LRs); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
