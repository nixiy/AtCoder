package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	a := nextInt()
	b := nextInt()

	// 10-1のパターンの場合
	if a == 1 && b == 10 {
		fmt.Println("Yes")
	} else {
		if b-a == 1 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
