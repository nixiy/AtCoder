package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Coordinate struct {
	x  int
	y  int
	lr string // LR
}

func solve(N int, cSlice []Coordinate, lrStr string) bool {
	yCordToCoodSlice := make(map[int][]Coordinate) // 同一Y軸ごとにまとめる

	// 走査するy軸の上限下限を決める
	yMin := 10000000000
	yMax := -1

	for i := 0; i < N; i++ {
		chmin(&yMin, cSlice[i].y)
		chmax(&yMax, cSlice[i].y)
	}

	// 各座標がどっち方向に移動するか追加する
	for i, lrByte := range lrStr {
		cSlice[i].lr = string(lrByte)
	}

	// y軸でまとめる
	for i := 0; i < N; i++ {
		yCordToCoodSlice[cSlice[i].y] = append(yCordToCoodSlice[cSlice[i].y], cSlice[i])
	}

	// まとめたy軸でループ
	for _, cSlice := range yCordToCoodSlice {
		if len(cSlice) > 1 { // 対象y軸の座標が1個以上の場合のみ検証(1個以下は衝突しないため)
			lCnt := 0
			rCnt := 0

			// R方向に移動する最小値rMinと、L方向に移動する最大値lMaxを取る
			rMin := 10000000000 // 10^9よりでかくしておく
			lMax := -1
			for _, c := range cSlice {
				if c.lr == "L" {
					lCnt++
					chmax(&lMax, c.x)
				} else {
					rCnt++
					chmin(&rMin, c.x)
				}
			}

			// 双方向に移動するy軸があれば衝突の可能性あり
			if lCnt > 0 && rCnt > 0 {
				// rMin < lMax であれば衝突する。でなければ、反対方向に永遠に進むので衝突しない
				if rMin < lMax {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	N := ni()
	cSlice := make([]Coordinate, N)

	for i := 0; i < N; i++ {
		cSlice[i].x = ni()
		cSlice[i].y = ni()
	}

	lrStr := ns()

	if solve(N, cSlice, lrStr) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

var sc = bufio.NewScanner(os.Stdin)

func ns() string        { sc.Scan(); return sc.Text() }
func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

// ex chmin(&p, v)
func chmin(p *int, v int) bool {
	if *p > v {
		*p = v
		return true
	}
	return false
}

// ex chmax(&p, v)
func chmax(p *int, v int) bool {
	if *p < v {
		*p = v
		return true
	}
	return false
}

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
