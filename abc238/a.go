package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func pow(x, y int) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	// TODO debug
	// fmt.Println("2n: ", pow(2, n))
	// fmt.Println("n2: ", pow(n, 2))

	// 単純にint64を用いても、nが62を超えるとオーバーフローしてしまう
	// n=1: 2^n=2, n^2=1
	// n=2: 2^n=4, n^2=4   ✅
	// n=3: 2^n=8, n^2=9   ✅
	// n=4: 2^n=16, n^2=16 ✅
	// n=5: 2^n=32, n^2=25
	// n=6: 2^n=64, n^2=36
	// 上記から2 <= n <= 4以外の範囲を除いて2^nの方が大きくなるため、そのまま判定する
	if 2 <= n && n <= 4 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
