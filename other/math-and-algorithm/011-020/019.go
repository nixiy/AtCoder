package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := ni()
	cardForNum := make(map[int]int)
	for i := 0; i < N; i++ {
		cardForNum[ni()]++
	}
	r := cardForNum[1]
	y := cardForNum[2]
	b := cardForNum[3]
	// 色それぞれに対して、rC2 + yC2 + bC2 を計算する事で全パターンを出せる
	fmt.Println(comb(r, 2) + comb(y, 2) + comb(b, 2))
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
