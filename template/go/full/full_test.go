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
	assert.Equal(t, 0, pow(2, -2), "負の指数")
	assert.Equal(t, 1, pow(0, 0))
	assert.Equal(t, 0, pow(0, 1))
	assert.Equal(t, 1, pow(1, 0))
	assert.Equal(t, 1, pow(1, 1))
	assert.Equal(t, 4, pow(2, 2))
	assert.Equal(t, 4, pow(-2, 2))
}

func Test_reverse(t *testing.T) {
	assert.Equal(t, reverse(""), "")
	assert.Equal(t, reverse("a"), "a")
	assert.Equal(t, reverse("abc"), "cba")
}

func Test_itob(t *testing.T) {
	assert.Equal(t, itob(0), "")
	assert.Equal(t, itob(8), "1000")
	assert.Equal(t, itob(1024), "10000000000")
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

		{name: "",
			args: args{
				a:      []int{2, 2, 2, 3, 3, 3, 4},
				target: 5,
			}, want: 7},
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
	assert.Equal(t, 55,
		sumArithmeticProgression_l(10, 1, 10))
	assert.Equal(t, 5050,
		sumArithmeticProgression_l(100, 1, 100))
}

func Test_sumArithmeticProgression_d(t *testing.T) {
	assert.Equal(t, 100,
		sumArithmeticProgression_d(10, 1, 2))
	assert.Equal(t, 10000,
		sumArithmeticProgression_d(100, 1, 2))
}

func Test_intStack_pushAndPop(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var stack intStack
		stack.push(1)
		stack.push(2)
		stack.push(3)
		assert.True(t, reflect.DeepEqual(stack, intStack{1, 2, 3}))
		assert.Equal(t, stack.first(), 1)
		assert.Equal(t, stack.last(), 3)
		assert.Equal(t, 3, stack.pop())
		assert.False(t, stack.empty())

		assert.Equal(t, stack.first(), 1)
		assert.Equal(t, stack.last(), 2)
		assert.Equal(t, 2, stack.pop())
		assert.False(t, stack.empty())

		assert.Equal(t, stack.first(), 1)
		assert.Equal(t, stack.last(), 1)
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
	assert.Equal(t, 2, max(1, 2))
	assert.Equal(t, 2, max(2, 1))
	assert.Equal(t, 2, max(2, 2))
	assert.Equal(t, 9223372036854775807, max(2, 9223372036854775807))
	assert.Equal(t, 2, max(2, -9223372036854775808))
}

func Test_min(t *testing.T) {
	assert.Equal(t, 1, min(1, 2))
	assert.Equal(t, 1, min(2, 1))
	assert.Equal(t, 2, min(2, 2))
	assert.Equal(t, 2, min(2, 9223372036854775807))
	assert.Equal(t, -9223372036854775808, min(2, -9223372036854775808))
}

func Test_abs(t *testing.T) {
	assert.Equal(t, 0, abs(0))
	assert.Equal(t, 1, abs(1))
	assert.Equal(t, 1, abs(-1))
	assert.Equal(t, 9223372036854775807, abs(9223372036854775807))
	assert.Equal(t, 9223372036854775807, abs(-9223372036854775807))
}

func Test_gcd(t *testing.T) {
	assert.Equal(t, 3, gcd(6, 3))
	assert.Equal(t, 14, gcd(28, 42))
	assert.Equal(t, 15, gcd(465, 360))
}

func Test_lcm(t *testing.T) {
	assert.Equal(t, 6, lcm(2, 3))
	assert.Equal(t, 12, lcm(4, 6))
	assert.Equal(t, 4064, lcm(127, 32))
	assert.Equal(t, 5000020927008107344,
		lcm(1000000432, 10000037534),
		"オーバーフローチェック")
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
	assert.Equal(t, "", getRune("abcde", 10))
	assert.Equal(t, "d", getRune("abcde", 3))
	assert.Equal(t, "", getRune("abcde", -1))
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

func Test_atoi(t *testing.T) {
	assert.Equal(t, 1, atoi("1"))
	assert.Equal(t, 10, atoi("10"))
	assert.Equal(t, 1, atoi("01"))
	assert.Equal(t, -1, atoi("-1"))
	assert.Equal(t, 0, atoi("0xFF"))
	assert.Equal(t, 0, atoi("0b11111111"))
}

func Test_itoa(t *testing.T) {
	assert.Equal(t, "1", itoa(1))
	assert.Equal(t, "10", itoa(10))
	assert.Equal(t, "-1", itoa(-1))
	assert.Equal(t, "255", itoa(0xFF))
	assert.Equal(t, "255", itoa(0b11111111))
}

func Test_prime(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name  string
		args  args
		wantP []int
	}{
		{name: "", args: args{N: -1}, wantP: []int{}},
		{name: "", args: args{N: 0}, wantP: []int{}},
		{name: "", args: args{N: 1}, wantP: []int{}},
		{name: "", args: args{N: 2}, wantP: []int{2}},
		{name: "", args: args{N: 10}, wantP: []int{2, 3, 5, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotP := prime(tt.args.N); !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("prime() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {
	assert.False(t, isPrime(-1))
	assert.False(t, isPrime(0))
	assert.False(t, isPrime(1))

	assert.True(t, isPrime(2))
	assert.True(t, isPrime(7))

	assert.True(t, isPrime(67280421310721))
}

func Test_divisor(t *testing.T) {
	type args struct {
		N      int
		isSort bool
	}
	tests := []struct {
		name    string
		args    args
		wantDiv []int
	}{
		{args: args{
			N:      100,
			isSort: false,
		}, wantDiv: []int{1, 100, 2, 50, 4, 25, 5, 20, 10}},

		{args: args{
			N:      100,
			isSort: true,
		}, wantDiv: []int{1, 2, 4, 5, 10, 20, 25, 50, 100}},

		{args: args{
			N:      827847039317,
			isSort: false,
		}, wantDiv: []int{1, 827847039317, 909859, 909863}},

		{args: args{
			N:      827847039317,
			isSort: true,
		}, wantDiv: []int{1, 909859, 909863, 827847039317}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiv := divisor(tt.args.N, tt.args.isSort); !reflect.DeepEqual(gotDiv, tt.wantDiv) {
				t.Errorf("divisor() = %v, want %v", gotDiv, tt.wantDiv)
			}
		})
	}
}

func Test_factorization(t *testing.T) {
	assert.Equal(t, []int{2, 5}, factorization(10))
	assert.Equal(t, []int{2, 2, 3, 3}, factorization(36))
}

func Test_multiGcd(t *testing.T) {
	assert.Equal(t, -1, multiGcd([]int{1}))
	assert.Equal(t, 6, multiGcd([]int{12, 18}))
	assert.Equal(t, 6, multiGcd([]int{12, 18, 24}))
}

func Test_multiLcm(t *testing.T) {
	assert.Equal(t, 252, multiLcm([]int{12, 18, 14}))
	assert.Equal(t, 4680, multiLcm([]int{120, 156, 180}))
}

func Test_comb(t *testing.T) {
	assert.Equal(t, 1, comb(4, 0))
	assert.Equal(t, 6, comb(4, 2))
	assert.Equal(t, 1, comb(4, 4))
	assert.Equal(t, 41417124750, comb(1000, 4))
}
