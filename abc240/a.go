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

func main() {
	sc.Split(bufio.ScanWords)
	a := nextInt()
	b := nextInt()

	// 10-1のパターンの場合
	if a == 1 && b == 10 {
		fmt.Println("Yes")
	} else {
		// 10-1以外のパターンで隣り合う場合はdiffが1になる
		if b-a == 1 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
