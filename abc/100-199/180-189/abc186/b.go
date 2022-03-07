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
	H := nextInt()
	W := nextInt()
	min := 10000
	max := 0
	nozokuBlock := 0

	// x*yのbool型のスライス
	field := make([][]int, H)
	for i := 0; i < H; i++ {
		field[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			field[i][j] = nextInt()
			if min > field[i][j] {
				min = field[i][j]
			}

			if max < field[i][j] {
				max = field[i][j]
			}
		}
	}

	for _, v := range field {
		for _, k := range v {
			nozokuBlock += k - min
		}
	}

	if min != max {
		fmt.Println(nozokuBlock)
	} else {
		fmt.Println(0)
	}
}
