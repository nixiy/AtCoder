package main

import (
	"bufio"
	"os"
	"strconv"
)

func prime(N int) (p []int) {
	isPrimeSlice := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		isPrimeSlice[i] = true
	}

	// エラトステネスの篩
	for i := 2; i*i <= N; i++ {
		if isPrimeSlice[i] {
			for x := 2 * i; x <= N; x += i {
				isPrimeSlice[x] = false
			}
		}
	}

	for i := 2; i <= N; i++ {
		if isPrimeSlice[i] {
			p = append(p, i)
		}
	}

	return p
}

func main() {
	defer wr.Flush()
	for _, p := range prime(ni()) {
		wr.WriteString(itoa(p) + " ")
	}
	wr.WriteByte('\n')
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
