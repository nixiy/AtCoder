package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, r := ni(), ni()
	fmt.Println(comb(n, r))
}

// 組み合わせの数nCr
func comb(n, r int) int {
	if r == 0 {
		return 1
	} else {
		return comb(n, r-1) * (n - r + 1) / r
	}
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
