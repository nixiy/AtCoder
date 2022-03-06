package main

import (
	"fmt"
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
	t.Run("", func(t *testing.T) {
		assert.Equal(t, 0, pow(2, -2), "負の指数")
		assert.Equal(t, 1, pow(0, 0))
		assert.Equal(t, 0, pow(0, 1))
		assert.Equal(t, 1, pow(1, 0))
		assert.Equal(t, 1, pow(1, 1))
		assert.Equal(t, 4, pow(2, 2))
		assert.Equal(t, 4, pow(-2, 2))
	})
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
		assert.True(t, chmin(&min, target))
		assert.Equal(t, 3, min)
	})
	t.Run("minが置き換わらないテスト", func(t *testing.T) {
		min := 5
		target := 6
		assert.False(t, chmin(&min, target))
		assert.Equal(t, 5, min)
	})
}

func Test_chmax(t *testing.T) {
	t.Run("maxが置き換わるテスト", func(t *testing.T) {
		max := 5
		target := 10
		assert.True(t, chmax(&max, target))
		assert.Equal(t, 10, max)
	})
	t.Run("maxが置き換わらないテスト", func(t *testing.T) {
		max := 5
		target := 2
		assert.False(t, chmax(&max, target))
		assert.Equal(t, 5, max)
	})
}

func Test_sumArithmeticProgression_l(t *testing.T) {
	t.Run("", func(t *testing.T) {
		assert.Equal(t, 55,
			sumArithmeticProgression_l(10, 1, 10))
	})

	t.Run("", func(t *testing.T) {
		assert.Equal(t, 5050,
			sumArithmeticProgression_l(100, 1, 100))
	})
}

func Test_sumArithmeticProgression_d(t *testing.T) {
	t.Run("", func(t *testing.T) {
		assert.Equal(t, 100,
			sumArithmeticProgression_d(10, 1, 2))
	})

	t.Run("", func(t *testing.T) {
		assert.Equal(t, 10000,
			sumArithmeticProgression_d(100, 1, 2))
	})
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
	t.Run("", func(t *testing.T) {
		assert.Equal(t, 2, max(1, 2))
		assert.Equal(t, 2, max(2, 1))
		assert.Equal(t, 2, max(2, 2))
		assert.Equal(t, 9223372036854775807, max(2, 9223372036854775807))
		assert.Equal(t, 2, max(2, -9223372036854775808))
	})
}

func Test_min(t *testing.T) {
	t.Run("", func(t *testing.T) {
		assert.Equal(t, 1, min(1, 2))
		assert.Equal(t, 1, min(2, 1))
		assert.Equal(t, 2, min(2, 2))
		assert.Equal(t, 2, min(2, 9223372036854775807))
		assert.Equal(t, -9223372036854775808, min(2, -9223372036854775808))
	})
}

func Test_abs(t *testing.T) {
	assert.Equal(t, 0, abs(0))
	assert.Equal(t, 1, abs(1))
	assert.Equal(t, 1, abs(-1))
	assert.Equal(t, 9223372036854775807, abs(9223372036854775807))
	assert.Equal(t, 9223372036854775807, abs(-9223372036854775807))
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

func Test_swap(t *testing.T) {
	a := 2
	b := 3
	swap(&a, &b)
	assert.Equal(t, 3, a)
	assert.Equal(t, 2, b)
}

func Test_getRune(t *testing.T) {
	t.Run("", func(t *testing.T) {
		assert.Equal(t, "", getRune("abcde", 10))
	})

	t.Run("", func(t *testing.T) {
		assert.Equal(t, "d", getRune("abcde", 3))
	})

	t.Run("", func(t *testing.T) {
		assert.Equal(t, "", getRune("abcde", -1))
	})
}

func Test_UnionFind(t *testing.T) {
	N := 7
	uf := NewUnionFind(N)
	assert.Equal(t, uf.Groups(),
		map[int][]int{0: {0}, 1: {1}, 2: {2}, 3: {3}, 4: {4}, 5: {5}, 6: {6}})

	uf.Merge(1, 2)
	assert.Equal(t, uf.Groups(),
		map[int][]int{0: {0},
			1: {1, 2},
			3: {3}, 4: {4}, 5: {5}, 6: {6}})
	fmt.Println(uf)

	uf.Merge(2, 3)
	assert.Equal(t, uf.Groups(),
		map[int][]int{0: {0},
			1: {1, 2, 3},
			4: {4}, 5: {5}, 6: {6}})
	fmt.Println(uf)

	uf.Merge(5, 6)
	assert.Equal(t, uf.Groups(),
		map[int][]int{0: {0},
			1: {1, 2, 3},
			4: {4},
			5: {5, 6}})
	fmt.Println(uf)

	// 既に同じ集合にある場合、変化なし
	uf.Merge(1, 3)
	assert.Equal(t, uf.Groups(),
		map[int][]int{0: {0},
			1: {1, 2, 3},
			4: {4},
			5: {5, 6}})
	fmt.Println(uf)

	assert.True(t, uf.IsSame(1, 3))
	assert.False(t, uf.IsSame(2, 5))

	uf.Merge(1, 6)
	assert.Equal(t, uf.Groups(),
		map[int][]int{0: {0},
			1: {1, 2, 3, 5, 6},
			4: {4}})

	uf.Merge(4, 3)
	fmt.Println(uf.Groups())

	// 自身が属する集合の個数
	assert.Equal(t, 1, uf.Size(0))

	assert.Equal(t, 6, uf.Size(1))
	assert.Equal(t, 6, uf.Size(2))
	assert.Equal(t, 6, uf.Size(3))
	assert.Equal(t, 6, uf.Size(4))
	assert.Equal(t, 6, uf.Size(5))
	assert.Equal(t, 6, uf.Size(6))

	fmt.Println(uf)
}

func Test_sqrt(t *testing.T) {
	assert.Equal(t, 0.0, sqrt(0))
	assert.Equal(t, 2.0, sqrt(4))
	assert.Equal(t, 1.4142135623730951, sqrt(2))
}

func Test_diff(t *testing.T) {
	assert.Equal(t, 0, diff('a', 'a'))
	assert.Equal(t, 1, diff('a', 'b'))
	assert.Equal(t, 2, diff('a', 'c'))
	assert.Equal(t, 1, diff('b', 'c'))
	assert.Equal(t, 1, diff('z', 'a'))
	assert.Equal(t, 25, diff('a', 'z'))
}

func Test_byteShift(t *testing.T) {
	assert.Equal(t, "a", byteShift('a', 0))  // 同値
	assert.Equal(t, "a", byteShift('a', 26)) // 同値
	assert.Equal(t, "b", byteShift('a', 1))
	assert.Equal(t, "a", byteShift('z', 1))
	assert.Equal(t, "z", byteShift('a', -1))
}

func Test_strShift(t *testing.T) {
	assert.Equal(t, "abcABC", strShift("abcABC", 0))  // 同値
	assert.Equal(t, "abcABC", strShift("abcABC", 26)) // 同値
	assert.Equal(t, "bcdBCD", strShift("abcABC", 1))
	assert.Equal(t, "zabZAB", strShift("abcABC", -1))
	assert.Equal(t, "WKLV LV D VHFUHW PHVVDJH",
		strShift("THIS IS A SECRET MESSAGE", 3))
}

func Test_reverseString(t *testing.T) {
	assert.Equal(t, "cba", reverseString("abc"))
}
