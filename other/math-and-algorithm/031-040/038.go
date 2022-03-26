package main

import (
	"bufio"
	"os"
	"strconv"
)

type Day struct {
	L int
	R int
}

func main() {
	defer wr.Flush() // 最大100000行出力するため高速化
	N, Q := ni(), ni()
	A := make([]int, N)
	D := make([]Day, Q)

	for i := 0; i < N; i++ {
		A[i] = ni()
	}

	for i := 0; i < Q; i++ {
		D[i].L = ni() - 1 // 0indexedにしておく
		D[i].R = ni() - 1 // 0indexedにしておく
	}

	// 累積和の作成
	sumA := make([]int, N)
	for i := 0; i < N; i++ {
		sumA[i] += A[i]
		if i-1 >= 0 {
			sumA[i] += sumA[i-1]
		}
	}

	// 累積和の配列をsumとすると、LからRまでの和は sum[R] - sum[L-1] のO(1)で処理可能
	for i := 0; i < Q; i++ {
		l := D[i].L
		r := D[i].R
		if l-1 >= 0 {
			wr.WriteString(strconv.Itoa(sumA[r]-sumA[l-1]) + "\n")
		} else {
			wr.WriteString(strconv.Itoa(sumA[r]) + "\n")
		}
	}
}

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriterSize(os.Stdout, 1024*1024) // 表示量が非常に多い時用

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
