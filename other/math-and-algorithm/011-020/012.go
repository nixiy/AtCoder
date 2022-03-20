package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := ni()
	if isPrime(N) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// 高速な素数判定
func isPrime(N int) bool {
	if N < 2 {
		return false
	} else {
		// 2 - √N まで調べれば良い
		for i := 2; i*i <= N; i++ {
			if N%i == 0 {
				return false
			}
		}
		return true
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
