package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := ni()
	h := make([]int, N+2)
	dp := make([]int, N+2) // 2マス先まで飛ぶバッファ
	for i := 0; i < N; i++ {
		h[i] = ni()
	}

	// dp初期化 (最小値を取るため、大きい数値を入れておく)
	// 0マス目からスタートでコスト0なので、1以降を初期化
	for i := 1; i < N+2; i++ {
		dp[i] = 10000000000
	}

	// dp
	for i := 0; i < N; i++ {
		chmin(&dp[i+1], dp[i]+abs(h[i]-h[i+1]))
		chmin(&dp[i+2], dp[i]+abs(h[i]-h[i+2]))
	}

	fmt.Println(dp[N-1])
}

// 絶対値で返す
func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// ex chmin(&p, v)
func chmin(p *int, v int) bool {
	if *p > v {
		*p = v
		return true
	}
	return false
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
