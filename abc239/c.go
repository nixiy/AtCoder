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

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func main() {
	sc.Split(bufio.ScanWords)
	// 2次元平面座標(x1, y1), (x2, y2)
	x1, y1 := nextInt(), nextInt()
	x2, y2 := nextInt(), nextInt()

	// 総当りしてみる
	// (Min ~ Max)^2 の範囲を検索すると4e+18となり、間に合わない。
	// AtCoderの制限時間2秒に収まるのは概ね10^7くらい。10^8だと処理内容によっては2秒を超える模様。

	// 処理を減らす方法を取るので以下を考えれば良い
	// x1, y1を中心とした4x4四方に含まれる点
	// すなわち以下の範囲のみ探索すれば良い。
	// -2 <= x1 <= 2
	// -2 <= y1 <= 2
	for x := x1 - 2; x <= x1+2; x++ {
		for y := y1 - 2; y <= y1+2; y++ {
			// fmt.Println(x, y) // TODO debug
			ansX := pow(x1-x, 2) + pow(y1-y, 2)
			ansY := pow(x2-x, 2) + pow(y2-y, 2)
			if ansX == 5 && ansY == 5 {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}
