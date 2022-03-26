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
	A := nis(N)
	sort.Ints(A)

	for i := 0; i <= 2000; i++ {
		flag := false
		for j := 0; j < N; j++ {
			if i == A[j] {
				flag = true
				break
			}
		}
		if flag == false {
			fmt.Println(i)
			return
		}
	}
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }
func nis(n int) (a []int) {
	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni()
	}
	return a
}

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
