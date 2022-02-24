package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := nextInt()
	heights := make([]int, n+1)
	heights[n] = -1 // 番兵
	for i := 0; i < n; i++ {
		heights[i] = nextInt()
	}

	for i := 0; i < n; i++ {
		// 1個先の台が同じ高さまたは小さかったらそこでやめて、現在位置の高さを表示する
		if heights[i] >= heights[i+1] {
			fmt.Println(heights[i])
			return
		}
	}
}
