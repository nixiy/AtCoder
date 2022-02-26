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

	for i := 0; i < seqLen; i++ {
		seq[i] = nextInt()
	}
	for i := 0; i < pairLen; i++ {
		x[i] = nextInt()
		k[i] = nextInt()
	}

	// 全探索
	for pi := 0; pi < pairLen; pi++ {
		// xがk回目に出てくるのは何番目か
		c := 0
		for si := 0; si < seqLen; si++ {
			// 数列とxの箇所が同じ
			if seq[si] == x[pi] {
				c++
				// xの登場回数がk回目であれば
				if c == k[pi] {
					fmt.Println(si + 1)
					break
				}
			}
		}
		if c == 0 || c < k[pi] {
			fmt.Println(-1)
		}
	}
}