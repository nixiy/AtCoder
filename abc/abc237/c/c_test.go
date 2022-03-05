package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "prefixにaを付けたら回文になるパターン", args: args{s: "kasaka"}, want: true},
		{name: "どうやっても回文にならないパターン", args: args{s: "atcoder"}, want: false},
		{name: "付け足さずとも回文のパターン", args: args{s: "php"}, want: true},
		{name: "aのみで構成された文字列", args: args{s: "aaa"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.s); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
