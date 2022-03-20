package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func bitAll(N, S int, A []int) bool {
	for i := 0; i < (1 << N); i++ {
		sum := 0
		for j := 1; j <= N; j++ {
			if (i & (1 << (j - 1))) != 0 {
				sum += A[j]
			}
		}
		if sum == S {
			return true
		}
	}
	return false
}
func main() {
	N, S := ni(), ni()
	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		A[i] = ni()
	}

	if bitAll(N, S, A) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func ns() string        { sc.Scan(); return sc.Text() }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }
func itoa(i int) string { return strconv.Itoa(i) }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
