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

	fmt.Println(multiLcm(A))
}

// 最大公約数: greatest common divisor
func gcd(a, b int) int {
	if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
}

// 最小公倍数
func lcm(a, b int) int { return b / gcd(a, b) * a }

// 3つ以上の値の最小公倍数
func multiLcm(target []int) (ans int) {
	if len(target) <= 1 {
		ans = -1
	} else {
		pastLcm := lcm(target[0], target[1])
		for i := 2; i < len(target); i++ {
			pastLcm = lcm(pastLcm, target[i])
		}
		ans = pastLcm
	}
	return ans
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
