package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func factorial(n int) int {
	if n == 1 {
		return n
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	fmt.Println(factorial(ni()))
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
