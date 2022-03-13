package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	defer wr.Flush()
	H, W := ni(), ni()
	A := make([][]int, H)
	for i := 0; i < H; i++ {
		A[i] = make([]int, W)
	}

	hTotalSlice := make([]int, H)
	wTotalSlice := make([]int, W)

	// 各行と列の総和を保持しておく
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			v := ni()
			A[i][j] = v
			hTotalSlice[i] += v
			wTotalSlice[j] += v
		}
	}

	// 自身の属する行と列の総和から自身を引くと合計値になる
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			wr.WriteString(itoa(hTotalSlice[i] + wTotalSlice[j] - A[i][j]))
			wr.WriteByte(' ')
		}
		wr.WriteByte('\n')
	}
}

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriterSize(os.Stdout, 1024*1024) // 表示量が非常に多い時用

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }
func itoa(i int) string { return strconv.Itoa(i) }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
