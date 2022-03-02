package main

import (
	"reflect"
	"testing"
)

func Test_uniq(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name     string
		args     args
		wantUniq []int
	}{
		{
			name: "単純なテスト", args: args{
				input: []int{1, 1, 2, 3, 3}},
			wantUniq: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUniq := uniq(tt.args.input); !reflect.DeepEqual(gotUniq, tt.wantUniq) {
				t.Errorf("uniq() = %v, want %v", gotUniq, tt.wantUniq)
			}
		})
	}
}

func Test_pow(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{x: 0, y: 0}, want: 1},
		{name: "2", args: args{x: 0, y: 1}, want: 0},
		{name: "3", args: args{x: 1, y: 0}, want: 1},
		{name: "4", args: args{x: 1, y: 1}, want: 1},
		{name: "5", args: args{x: 2, y: 2}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0", args: args{s: ""}, want: ""},
		{name: "0", args: args{s: "a"}, want: "a"},
		{name: "0", args: args{s: "abc"}, want: "cba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.s); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itob(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{x: 0}, want: ""},
		{name: "", args: args{x: 8}, want: "1000"},
		{name: "", args: args{x: 1024}, want: "10000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := itob(tt.args.x); got != tt.want {
				t.Errorf("itob() = %v, want %v", got, tt.want)
			}
		})
	}
}
