package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := ni()
	dp := make([]int, N+2) // i段目までの移動方法が何通りあるかのdpテーブル(2マス先まで配るのでバッファを設けている)
	dp[0] = 1              // 0段目から開始なので1通り

	for i := 0; i < N; i++ {
		dp[i+1] += dp[i]
		dp[i+2] += dp[i]
	}

	fmt.Println(dp[N])
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
