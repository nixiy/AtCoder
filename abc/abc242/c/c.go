package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 998244353

// 第一引数側に加算代入する
// 加算するたびにMOD取らないと簡単にオーバーフローするので注意
func add(a *int, b int) {
	*a += b
	*a %= MOD
}

func solve(N int) int {
	dp := make([][10]int, N+1)
	// 1桁目の取りうるパターンは1通りしかないためdpテーブルを初期化
	for d := 1; d <= 9; d++ {
		dp[1][d] = 1
	}

	// 2桁目からN桁目まで試す
	for n := 2; n <= N; n++ {
		for digit := 1; digit <= 9; digit++ {
			// 左上からの加算
			if digit-1 >= 1 {
				add(&dp[n][digit], dp[n-1][digit-1])
			}
			// 左からの加算
			add(&dp[n][digit], dp[n-1][digit])

			// 左下からの加算
			if digit+1 <= 9 {
				add(&dp[n][digit], dp[n-1][digit+1])
			}
		}
	}

	// N桁目の1~9の総和を足す
	sum := 0
	for digit := 1; digit <= 9; digit++ {
		add(&sum, dp[N][digit])
	}

	return sum
}
func main() {
	n := ni()
	fmt.Println(solve(n))
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
