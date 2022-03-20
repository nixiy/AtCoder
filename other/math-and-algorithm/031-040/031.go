package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Study struct {
	yes int
	no  int
}

func main() {
	N := ni()
	A := make([]int, N+2)
	for i := 0; i < N; i++ {
		A[i] = ni()
	}
	dp := make([]Study, N+2)
	dp[0].yes = 0 // 初期の実力は0
	dp[0].no = 0  // 初期の実力は0

	for i := 1; i <= N; i++ {
		// 2日連続で勉強できないため、前日勉強していない実力と本日分を加算する
		dp[i].yes = dp[i-1].no + A[i]
		// 勉強しない行為は2日連続でも可能、よって以下2つのMaxを取る
		// 1. 前日勉強した場合の実力
		// 2. 前日勉強しない場合の実力
		dp[i].no = max(dp[i-1].yes, dp[i-1].no)
	}

	// 最終的にN日目に勉強した場合としない場合のMaxを取る
	fmt.Println(max(dp[N].yes, dp[N].no))
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
