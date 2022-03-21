package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Goods struct {
	w int // 重さ
	v int // 価値
}

func main() {
	N, W := ni(), ni()
	goods := make([]Goods, N+2)
	for i := 1; i <= N; i++ {
		goods[i].w = ni()
		goods[i].v = ni()
	}

	// dp[N個の品物][重さ] = 価値
	dp := make([][]int, N+2)
	for i := 0; i < N+2; i++ {
		dp[i] = make([]int, W+2)
	}
	// dp[0]の価値の最大値は全て0(初期値が0の為特に初期化は行わない)

	// dp埋め開始
	for n := 1; n <= N; n++ {
		for w := 1; w <= W; w++ {
			if w-goods[n].w >= 0 { // out of rangeにならないように
				dp[n][w] = max(dp[n-1][w], dp[n-1][w-goods[n].w]+goods[n].v)
			} else {
				dp[n][w] = dp[n-1][w]
			}
		}
	}
	fmt.Println(dp[N][W])
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
