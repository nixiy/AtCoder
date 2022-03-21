package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x    int
	y    int
	posi int // (x,y)について(0,0):0, (1,0):1, (0,1):y*Widthのように番号をふり、どのマス目なのかを指し示す
}

func main() {
	Height, Width := ni(), ni()

	start := Coordinate{}
	start.y = ni() - 1
	start.x = ni() - 1
	start.posi = start.x + start.y*Width

	goal := Coordinate{}
	goal.y = ni() - 1
	goal.x = ni() - 1
	goal.posi = goal.x + goal.y*Width

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

	// 拡頂点に番号を振り、到達可能な頂点番号をスライスで持っている。
	G := make([][]int, Height*Width) // グラフ

	// 横方向の辺をグラフに追加
	for h := 0; h < Height; h++ {
		for w := 0; w < Width-1; w++ {
			idx1 := h*Width + w
			idx2 := h*Width + (w + 1)
			if masu[h][w] == "." && masu[h][w+1] == "." {
				G[idx1] = append(G[idx1], idx2)
				G[idx2] = append(G[idx2], idx1)
			}
		}
	}

	// 四辺が壁に囲まれているため、index1から開始している
	for h := 0; h < Height-1; h++ {
		for w := 0; w < Width; w++ {
			idx1 := h*Width + w
			idx2 := (h+1)*Width + w
			if masu[h][w] == "." && masu[h+1][w] == "." {
				G[idx1] = append(G[idx1], idx2)
				G[idx2] = append(G[idx2], idx1)
			}
		}
	}

	// 幅優先探索の初期化 dist[h] = -1 の時未到達頂点である
	dist := make([]int, Height*Width)
	for i := 0; i < Height*Width; i++ {
		dist[i] = -1
	}
	var Q intQueue
	Q.enqueue(start.posi) // 行動開始1をqueueにセット
	dist[start.posi] = 0  // スタート地点の最短経路は0通り

	// 幅優先探索
	for !Q.empty() {
		pos := Q.dequeue()
		for i := 0; i < len(G[pos]); i++ {
			next := G[pos][i]
			if dist[next] == -1 {
				dist[next] = dist[pos] + 1
				Q.enqueue(next)
			}
		}
	}

	// ans
	fmt.Println(dist[goal.posi])
}

type intQueue []int

func (queue *intQueue) empty() bool   { return len(*queue) == 0 }
func (queue *intQueue) first() int    { return (*queue)[0] }
func (queue *intQueue) last() int     { return (*queue)[len(*queue)-1] }
func (queue *intQueue) enqueue(i int) { *queue = append(*queue, i) }
func (queue *intQueue) dequeue() int {
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
