package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 幅優先探索
func bfs(Height, Width, startPosi, goalPosi int, masu [][]string, passable string) int {
	// 幅優先探索の初期化 dist[h] = -1 の時未到達頂点である
	dist := make([]int, Height*Width)
	// 拡頂点に番号を振り、到達可能な頂点番号をスライスで持っている。
	G := make([][]int, Height*Width) // グラフ

	// 横方向の辺をグラフに追加
	for h := 0; h < Height; h++ {
		for w := 0; w < Width-1; w++ {
			current := h*Width + w
			next := h*Width + (w + 1)
			if masu[h][w] == passable && masu[h][w+1] == passable {
				G[current] = append(G[current], next)
				G[next] = append(G[next], current)
			}
		}
	}

	// 縦方向の辺をグラフに追加
	for h := 0; h < Height-1; h++ {
		for w := 0; w < Width; w++ {
			current := h*Width + w
			next := (h+1)*Width + w
			if masu[h][w] == passable && masu[h+1][w] == passable {
				// 双方向に追加する
				G[current] = append(G[current], next)
				G[next] = append(G[next], current)
			}
		}
	}

	const NOT_REACHED = -1
	for i := 0; i < Height*Width; i++ {
		dist[i] = NOT_REACHED
	}

	var Q intQueue
	Q.push(startPosi)   // 行動開始1をqueueにセット
	dist[startPosi] = 0 // スタート地点の最短経路は0通り

	// 幅優先探索
	for !Q.empty() {
		// debug
		// fmt.Println("==============")
		// cnt := 0
		// for i := 0; i < Height*Width; i++ {
		// 	if dist[i] == -1 {
		// 		fmt.Print("👺\t")
		// 	} else {
		// 		fmt.Print(dist[i], "\t")
		// 	}
		//
		// 	cnt++
		// 	if cnt == Width {
		// 		cnt = 0
		// 		fmt.Println()
		// 	}
		// }
		// fmt.Println("==============")

		pos := Q.pop()
		// 現在の頂点から行ける未踏の地へ行く
		for _, next := range G[pos] {
			if dist[next] == NOT_REACHED {
				dist[next] = dist[pos] + 1
				Q.push(next)
			}
		}
	}

	return dist[goalPosi]
}

func main() {
	Height, Width := ni(), ni()

	sy, sx := ni()-1, ni()-1
	startPosi := sx + sy*Width

	gy, gx := ni()-1, ni()-1
	goalPosi := gx + gy*Width

	// 盤面初期化 & 入力受け取り
	masu := make([][]string, Height)
	for h := 0; h < Height; h++ {
		masu[h] = make([]string, Width)
	}

	// 1文字ずつ盤面配列に埋めていく
	for h := 0; h < Height; h++ {
		S := strings.Split(ns(), "")
		for sIndex := 0; sIndex < len(S); sIndex++ {
			masu[h][sIndex] = S[sIndex]
		}
	}

	// 幅優先探索で解く
	shortestPass := bfs(Height, Width, startPosi, goalPosi, masu, ".")

	// ans
	fmt.Println(shortestPass)
}

type intQueue []int

func (queue *intQueue) empty() bool { return len(*queue) == 0 }
func (queue *intQueue) first() int  { return (*queue)[0] }
func (queue *intQueue) last() int   { return (*queue)[len(*queue)-1] }
func (queue *intQueue) push(i int)  { *queue = append(*queue, i) }
func (queue *intQueue) pop() int {
	result := (*queue)[0]
	*queue = (*queue)[1:]
	return result
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func ns() string        { sc.Scan(); return sc.Text() }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
