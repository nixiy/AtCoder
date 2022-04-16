package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve(N int) {

	if N == 1 {
		wr.WriteString(itoa(1) + " ")
	}

	if N >= 2 {
		solve(N - 1)
		wr.WriteString(itoa(N) + " ")
		solve(N - 1)
	}
}

func main() {
	defer wr.Flush()
	N := ni()
	solve(N)
}

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriterSize(os.Stdout, 1024*1024) // 表示量が非常に多い時用

func ns() string        { sc.Scan(); return sc.Text() }
func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }
func itoa(i int) string { return strconv.Itoa(i) }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
