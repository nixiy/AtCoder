package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	N := ni()

	// dp計算時にout of indexの事を考えるのが面倒なので、予め2マス先までの大きさを確保しておく
	height := make([]int, N+2)
	for i := 0; i < N; i++ {
		height[i] = ni()
	}

	// 🐸の行動としては以下の2択になる
	// 現在地からi+1の足場に行く
	// 現在地からi+2の足場に行く
	// dp配列を用意し、その足場に到達するときのコストを埋める(既に値がある場合はより小さい値に書き換える
	// こちらを実施すると、dp[N-1]の値が支払うコストの総和の最小値になる
	dp := make([]int, N+2)
	HEIGHT_INF := 1000000000000000000 // 最大の足場が10e4の為、限りなく大きい値にしておく
	dp[0] = 0                         // 初期位置から始まるため、当然初期位置に移動するコストは0である
	for i := 1; i < N; i++ {
		dp[i] = HEIGHT_INF
	}

	for i := 0; i < N; i++ {
		// 1マス先の値を埋める → min(1マス先のdp値, 現在マスのdp値+高低差の絶対値(現在地高さ+1マス先高さ)
		chmin(&dp[i+1], dp[i]+abs(height[i]-height[i+1]))

		// 2マス先の値を埋める → min(2マス先のdp値, 現在マスのdp値+高低差の絶対値(現在地高さ+2マス先高さ)
		chmin(&dp[i+2], dp[i]+abs(height[i]-height[i+2]))
	}

	fmt.Println(dp[N-1])
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

// ex chmin(&p, v)
func chmin(p interface{}, v interface{}) {
	switch v.(type) {
	case int:
		a, ok := p.(*int)
		if !ok {
			return
		}
		if vv := v.(int); *a > vv {
			*a = vv
		}
	case float32:
		a, ok := p.(*float32)
		if !ok {
			return
		}
		if vv := v.(float32); *a > vv {
			*a = vv
		}
	case float64:
		a, ok := p.(*float64)
		if !ok {
			return
		}
		if vv := v.(float64); *a > vv {
			*a = vv
		}
	}
}

func init() {
	sc.Split(bufio.ScanWords)
}
