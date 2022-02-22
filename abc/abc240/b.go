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
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// 数値配列をuniqして返す
func uniq(input []int) (uniq []int) {
	m := make(map[int]bool)
	for _, elm := range input {
		if !m[elm] {
			m[elm] = true
			uniq = append(uniq, elm)
		}
	}
	return uniq
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()          // 入力数値総数
	input := make([]int, n) // 入力数値配列

	for i := 0; i < n; i++ {
		input[i] = nextInt()
	}

	// 数値のユニーク個数を出力
	uniq := uniq(input)
	fmt.Println(len(uniq))
}
