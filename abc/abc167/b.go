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
	one, zero, minus, tryCnt := nextInt(), nextInt(), nextInt(), nextInt()
	sum := 0
	cards := []int{1, 0, -1}

	// 試行回数文、大きい数値を取っていけばよい
	for i, cardCnt := range []int{one, zero, minus} {
		if tryCnt >= cardCnt {
			sum += cards[i] * cardCnt
			tryCnt -= cardCnt
		} else {
			sum += cards[i] * tryCnt
			break
		}
	}
	fmt.Println(sum)
}
