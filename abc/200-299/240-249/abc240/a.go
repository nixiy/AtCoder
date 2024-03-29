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

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	a := nextInt()
	b := nextInt()

	// 10-1のパターンの場合
	if a == 1 && b == 10 {
		fmt.Println("Yes")
	} else {
		// 10-1以外のパターンはdiffが1になる
		if b-a == 1 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
