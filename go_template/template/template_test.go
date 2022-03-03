package main

import (
	"github.com/stretchr/testify/assert"
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
			name:     "空",
			args:     args{input: []int{1}},
			wantUniq: []int{1},
		},
		{
			name:     "Sortされてないテスト",
			args:     args{input: []int{1, 1, 3, 2, 3}},
			wantUniq: []int{1, 3, 2},
		},
		{
			name:     "単純なテスト",
			args:     args{input: []int{1, 1, 2, 3, 3}},
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
		{name: "", args: args{s: ""}, want: ""},
		{name: "", args: args{s: "a"}, want: "a"},
		{name: "", args: args{s: "abc"}, want: "cba"},
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

func Test_lowerBound(t *testing.T) {
	type args struct {
		a      []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				a:      []int{2, 2, 2, 3, 3, 3, 4},
				target: 1,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				a:      []int{2, 2, 2, 3, 3, 3, 4},
				target: 2,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				a:      []int{2, 2, 2, 3, 3, 3, 4},
				target: 3,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				a:      []int{2, 2, 2, 3, 3, 3, 4},
				target: 5,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowerBound(tt.args.a, tt.args.target); got != tt.want {
				t.Errorf("lowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_upperBound(t *testing.T) {
	type args struct {
		a      []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				a:      []int{2, 2, 2, 3, 3, 3, 4},
				target: 1,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				a:      []int{2, 2, 2, 3, 3, 3, 4},
				target: 2,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				a:      []int{2, 2, 2, 3, 3, 3, 4},
				target: 3,
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				a:      []int{2, 2, 2, 3, 3, 3, 4},
				target: 5,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := upperBound(tt.args.a, tt.args.target); got != tt.want {
				t.Errorf("upperBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chmin(t *testing.T) {
	t.Run("minが置き換わるテスト", func(t *testing.T) {
		min := 5
		target := 3
		chmin(&min, target)
		assert.Equal(t, 3, min)
	})
	t.Run("minが置き換わらないテスト", func(t *testing.T) {
		min := 5
		target := 6
		chmin(&min, target)
		assert.Equal(t, 5, min)
	})
}

func Test_chmax(t *testing.T) {
	t.Run("maxが置き換わるテスト", func(t *testing.T) {
		max := 5
		target := 10
		chmax(&max, target)
		assert.Equal(t, 10, max)
	})
	t.Run("maxが置き換わらないテスト", func(t *testing.T) {
		max := 5
		target := 2
		chmax(&max, target)
		assert.Equal(t, 5, max)
	})
}

func Test_sumArithmeticProgression_l(t *testing.T) {
	type args struct {
		n     int
		first int
		last  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:     10,
				first: 1,
				last:  10,
			},
			want: 55,
		},
		{
			args: args{
				n:     100,
				first: 1,
				last:  100,
			},
			want: 5050,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumArithmeticProgression_l(tt.args.n, tt.args.first, tt.args.last); got != tt.want {
				t.Errorf("sumArithmeticProgression_l() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumArithmeticProgression_d(t *testing.T) {
	type args struct {
		n     int
		first int
		diff  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n:     10,
				first: 1,
				diff:  2,
			}, want: 100,
		},
		{
			name: "",
			args: args{
				n:     100,
				first: 1,
				diff:  2,
			},
			want: 10000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumArithmeticProgression_d(tt.args.n, tt.args.first, tt.args.diff); got != tt.want {
				t.Errorf("sumArithmeticProgression_d() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intStack_pushAndPop(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var stack intStack
		stack.push(1)
		stack.push(2)
		stack.push(3)
		assert.True(t, reflect.DeepEqual(stack, intStack{1, 2, 3}))
		assert.Equal(t, 3, stack.pop())
		assert.False(t, stack.empty())
		assert.Equal(t, 2, stack.pop())
		assert.False(t, stack.empty())
		assert.Equal(t, 1, stack.pop())
		assert.True(t, stack.empty())
	})
}

func Test_intQueue_enqueue(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var queue intQueue
		queue.enqueue(1)
		queue.enqueue(2)
		queue.enqueue(3)
		assert.True(t, reflect.DeepEqual(queue, intQueue{1, 2, 3}))
		assert.Equal(t, 1, queue.dequeue())
		assert.False(t, queue.empty())
		assert.Equal(t, 2, queue.dequeue())
		assert.False(t, queue.empty())
		assert.Equal(t, 3, queue.dequeue())
		assert.True(t, queue.empty())
	})
}

func Test_max(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{x: 1, y: 2}, want: 2},
		{args: args{x: 2, y: 1}, want: 2},
		{args: args{x: 2, y: 2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{x: 1, y: 2}, want: 1},
		{args: args{x: 2, y: 1}, want: 1},
		{args: args{x: 1, y: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{x: 0}, want: 0},
		{args: args{x: 1}, want: 1},
		{args: args{x: -1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{a: 6, b: 3}, want: 3},
		{args: args{a: 28, b: 42}, want: 14},
		{args: args{a: 465, b: 360}, want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lcm(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{a: 2, b: 3}, want: 6},
		{args: args{a: 4, b: 6}, want: 12},
		{args: args{a: 127, b: 32}, want: 4064},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcm(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("lcm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{nums: []int{}},
			want: [][]int{{}},
		},
		{
			args: args{nums: []int{1}},
			want: [][]int{
				{1},
			},
		},
		{
			args: args{nums: []int{1, 2}},
			want: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			args: args{nums: []int{1, 2, 3}},
			want: [][]int{
				{1, 2, 3},
				{2, 1, 3},
				{3, 1, 2},
				{1, 3, 2},
				{2, 3, 1},
				{3, 2, 1},
			},
		},
		{
			args: args{nums: []int{1, 2, 2}},
			want: [][]int{
				{1, 2, 2},
				{2, 1, 2},
				{2, 1, 2},
				{1, 2, 2},
				{2, 2, 1},
				{2, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
