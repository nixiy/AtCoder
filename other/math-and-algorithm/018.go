package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := ni()
	goodsForPrice := make(map[int]int)
	for i := 0; i < N; i++ {
		goodsForPrice[ni()]++
	}

	// 500円になるのは、以下の2組のみ
	// - 100 400
	// - 200 300
	// 積の法則を用いて、ad + bc通りで計算する事ができる。
	fmt.Println(goodsForPrice[100]*goodsForPrice[400] + goodsForPrice[200]*goodsForPrice[300])
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
