package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	A, B := ni(), ni()
	fmt.Println(gcd(A, B))
}

// 最大公約数: greatest common divisor
func gcd(a, b int) int {
	if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
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
