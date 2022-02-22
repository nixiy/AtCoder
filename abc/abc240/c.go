package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	jump_n := nextInt() // 飛ぶ回数
	ans_p := nextInt()  // 正解の座標の位置
	first := make([]int, jump_n)
	second := make([]int, jump_n)

	// 飛び方の標準入力を受け取る
	for i := 0; i < jump_n; i++ {
		first[i] = nextInt()
		second[i] = nextInt()
	}

	// 条件
	// 複数のジャンプの仕方で同じ座標に到着する場合、それらは同一視してよい
	// 途中で座標がanswerより大きくなるようなジャンプの仕方は無視して良い

	// 次のような動的計画法を考える
	// dp[jump][p]: i回ジャンプを行った時点で座標pの位置にいる事が可能なら1、不可能なら0
	// 0 <= jump <= jump_n
	// 0 <= p <= ans_p
	// O(NX)

	// 初期化
	dp := make([][]bool, jump_n+1)
	for i := 0; i < jump_n+1; i++ {
		dp[i] = make([]bool, ans_p+1)
		for j := 0; j < ans_p+1; j++ {
			dp[i][j] = false
		}
	}
	dp[0][0] = true // 初期位置にいる事は可能なため

	// Solve
	// 全てのpについてdp[k+1][p] := 0 に初期化する
	// dp[k][p] = 1となる全てのpについて以下の処理を行う
	// p + first_k <= ans_p ならば dp[k+1][p+first_k] = 1とする
	// p + second_k <= ans_p ならば dp[k+1][p+second_k] = 1とする

	// 飛ぶ回数分
	for jump := 0; jump < jump_n; jump++ {
		// 正解の座標の位置分
		for p := 0; p <= ans_p; p++ {
			// 到達可能であれば
			if dp[jump][p] {
				// 現在位置p+ジャンプ数で正解座標に収まれば、移動先のフラグを立てる
				if p+first[jump] <= ans_p {
					dp[jump+1][p+first[jump]] = true
				}
				if p+second[jump] <= ans_p {
					dp[jump+1][p+second[jump]] = true
				}
			}
		}
	}

	// TODO
	// 生成したdp確認用
	// fmt.Println("* ans *")
	// for _, v := range dp {
	// 	fmt.Println(v)
	// }

	if dp[jump_n][ans_p] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
