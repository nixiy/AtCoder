package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const HEIGHT_INF int = 10000000000000000 // 最大の足場が10e4の為、限りなく大きい値にしておく
var height []int
var dp []int

// 概念としてはもらうDP寄り
// 足場N-1までの最小コストは足場N-2までのコストがわかっていれば良い。
// → 最終的に足場0に帰着したとき、終了する
func rec(i int) int {
	// 既にメモされていればそれを返す
	if dp[i] < HEIGHT_INF {
		return dp[i]
	}

	// 足場0への移動コストは開始地点なので0
	if i == 0 {
		return 0
	}

	res := HEIGHT_INF

	// 現在マスを1マス前の値で埋める。
	if i-1 >= 0 {
		chmin(&res, rec(i-1)+abs(height[i]-height[i-1]))
	}

	// 現在マスを2マス前の値で埋める。
	if i-2 >= 0 {
		chmin(&res, rec(i-2)+abs(height[i]-height[i-2]))
	}

	dp[i] = res
	return res
}

func main() {
	N := ni()

	// dp計算時にout of indexの事を考えるのが面倒なので、予め2マス先までの大きさを確保しておく
	height = make([]int, N+2)
	for i := 0; i < N; i++ {
		height[i] = ni()
	}

	dp = make([]int, N+2)

	for i := 0; i < N; i++ {
		dp[i] = HEIGHT_INF
	}

	fmt.Println(rec(N - 1))
	fmt.Println(dp)
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
