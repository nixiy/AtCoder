package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	n, x := ni(), ni()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni()
	}
	sort.Ints(a) // 昇順ソート

	lowerIndex := lowerBound(a, x)
	// あくまで境界値のindexを返すため、その値と比較してみる必要がある
	if lowerIndex != n && a[lowerIndex] == x {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// target以上を満たす最小の位置を0-basedの添字で返す
// リスト内のどの要素よりも大きい場合、len(a)を返す
// リスト内のどの要素よりも小さい場合、0を返す
func lowerBound(a []int, target int) int { return sort.SearchInts(a, target) }

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriterSize(os.Stdout, 1024*1024) // 表示量が非常に多い時用

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
