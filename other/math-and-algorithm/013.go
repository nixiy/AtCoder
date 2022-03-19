package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	N := ni()
	div := divisor(N, true)
	for _, d := range div {
		fmt.Println(d)
	}
}

// 高速な約数列挙(未ソート)
func divisor(N int, isSort bool) (div []int) {
	for i := 1; i*i <= N; i++ {
		if N%i != 0 {
			continue
		}
		div = append(div, i)
		if i != N/i {
			div = append(div, N/i)
		}
	}
	if isSort {
		sort.Ints(div)
	}
	return div
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
