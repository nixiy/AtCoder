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
	cardCnt := make([]int, n) // 各カードのカウント(0-indexed)
	for i := 0; i < n*4-1; i++ {
		cardCnt[nextInt()-1]++
	}
	for i, cardCnt := range cardCnt {
		if cardCnt == 3 {
			fmt.Println(i + 1)
		}
	}
}
