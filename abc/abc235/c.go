package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

// 数値配列をuniqして返す
func uniq(input []int) (uniq []int) {
	m := make(map[int]bool)
	for _, elm := range input {
		if !m[elm] {
			m[elm] = true
			uniq = append(uniq, elm)
		}
	}
	return uniq
}

// 等差数列の和
// ただし初項a、項数n、末項lがわかっている場合
func sumArithmeticProgression_l(n, a, l int) int {
	fmt.Println("n, a, l: ", n, a, l)
	return n * (a + l) / 2
}

// 等差数列の和
// ただし初項a、項数n、公差dがわかっている場合
func sumArithmeticProgression_d(n, a, d int) int {
	return (n / 2) * (2*a + (n-1)*d)
}

func yes() {
	fmt.Println("Yes")
}

func no() {
	fmt.Println("No")
}

// 頻出するYes No出力用
func printYesNo(b bool) {
	if b {
		yes()
	} else {
		no()
	}
}

// string[i]のように取得するとbyteで取得できてしまう
// 中間処理でruneを使用して部分文字を取得する
func getRune(str string, index int) string {
	rs := []rune(str)
	return string(rs[index])
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	seqLen, pairLen := nextInt(), nextInt()
	seq := make([]int, seqLen)
	x := make([]int, pairLen)
	k := make([]int, pairLen)
	// 1 1 2 3 1 2 の場合
	// 以下の様な数値がそれぞれ何番目に登場するのかを記録しておく
	// 2次元配列で宣言すると10億x10億の配列を作ることになり大変な事になる(やってしまった)
	// map[0:[1 2 5] 1:[3 6] 2:[4]]
	mat := make(map[int][]int)

	for i := 0; i < seqLen; i++ {
		seq[i] = nextInt()
	}
	for i := 0; i < pairLen; i++ {
		x[i] = nextInt()
		k[i] = nextInt()
	}

	for i := 0; i < seqLen; i++ {
		// matrix[今見ている数]にi+1を追加する
		mat_i := seq[i] - 1
		mat[mat_i] = append(mat[mat_i], i+1)
	}

	fmt.Println(mat)

	for pi := 0; pi < pairLen; pi++ {
		xp := x[pi] - 1
		kp := k[pi] - 1
		// 対象の数値がseqに存在するか
		// 対象のkpが配列のindex内に収まるか
		if len(mat[xp]) != 0 && len(mat[xp]) > kp {
			fmt.Println(mat[xp][kp])
		} else {
			fmt.Println(-1)
		}
	}
}
