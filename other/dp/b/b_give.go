package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N, K int, h []int) int {
	// 🐸の行動としては以下のK択になる
	// 現在地からi+1の足場に行く
	// 現在地からi+2の足場に行く
	// 現在地からi+Kの足場に行く
	// dp配列を用意し、その足場に到達するときのコストを埋める(既に値がある場合はより小さい値に書き換える
	// こちらを実施すると、dp[N-1]の値が支払うコストの総和の最小値になる
	dp := make([]int, N+K)
	HEIGHT_INF := 2 << 60 // 最小値を求めるので、足して64bitに収まる程度の大きな数値にしておく
	dp[0] = 0             // 初期位置から始まるため、当然初期位置に移動するコストは0である
	for i := 1; i < N+K; i++ {
		dp[i] = HEIGHT_INF
	}

	for i := 0; i < N; i++ {
		// K個先(1indexedな事に注意)
		for k := 1; k <= K; k++ {
			// kマス先の値を埋める → min(kマス先のdp値, 現在マスのdp値+高低差の絶対値(現在地高さ+kマス先高さ)
			if i+k < N {
				chmin(&dp[i+k], dp[i]+abs(h[i]-h[i+k]))
			}
		}
	}

	return dp[N-1]
}

func main() {
	N, K := ni(), ni()

	// dp計算時にout of indexの事を考えるのが面倒なので、予めKマス先までの大きさを確保しておく
	h := make([]int, N+K)
	for i := 0; i < N; i++ {
		h[i] = ni()
	}

	fmt.Println(solve(N, K, h))
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return x * -1
}

// ex chmin(&p, v)
func chmin(p *int, v int) {
	if *p > v {
		*p = v
	}
}

func init() {
	sc.Split(bufio.ScanWords)
}
