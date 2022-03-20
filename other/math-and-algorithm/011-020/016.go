package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := ni()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = ni()
	}

	fmt.Println(multiGcd(A))
}

// 3つ以上の値の最大公約数
func multiGcd(target []int) (ans int) {
	if len(target) <= 1 {
		ans = -1
	} else if len(target) == 2 {
		ans = gcd(target[0], target[1])
	} else {
		pastGcd := gcd(target[0], target[1])
		for i := 2; i < len(target); i++ {
			pastGcd = gcd(pastGcd, target[i])
		}
		ans = pastGcd
	}
	return ans
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
