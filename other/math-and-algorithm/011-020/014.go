package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := ni()
	f := factorization(N)
	for _, f := range f {
		fmt.Print(f, " ")
	}
}

// 高速な素因数分解
func factorization(N int) (f []int) {
	for i := 2; i*i <= N; i++ {
		for N%i == 0 {
			N /= i
			f = append(f, i)
		}
	}
	// 素数の最小値以上が残っていれば追加
	if N >= 2 {
		f = append(f, N)
	}
	return f
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
