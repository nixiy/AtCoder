package main

import (
	"bufio"
	"os"
	"strconv"
)

type Cp struct {
	c int
	p int
}

type Sum struct {
	one int
	two int
}

func main() {
	defer wr.Flush()
	N := ni()
	cpSlice := make([]Cp, N+1)
	for i := 1; i <= N; i++ {
		cpSlice[i].c = ni()
		cpSlice[i].p = ni()
	}

	sumSlice := make([]Sum, N+1)

	// それぞれの累積和を取る
	for i := 1; i <= N; i++ {
		// 1個前の累積を引き継ぐ
		sumSlice[i].one = sumSlice[i-1].one
		sumSlice[i].two = sumSlice[i-1].two
		// 現在の数を累積
		if cpSlice[i].c == 1 {
			sumSlice[i].one += cpSlice[i].p
		} else {
			sumSlice[i].two += cpSlice[i].p
		}
	}

	Q := ni()
	for i := 1; i <= Q; i++ {
		l, r := ni(), ni()
		ans1 := sumSlice[r].one - sumSlice[l-1].one
		ans2 := sumSlice[r].two - sumSlice[l-1].two
		wr.WriteString(itoa(ans1) + " " + itoa(ans2) + "\n")
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
