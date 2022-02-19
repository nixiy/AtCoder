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

func main() {
	sc.Split(bufio.ScanWords)

	h := nextInt()
	w := nextInt()

	// generate 2d array
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
	}

	for i_h := 0; i_h < h; i_h++ {
		for i_w := 0; i_w < w; i_w++ {
			a[i_h][i_w] = nextInt()
		}
	}

	for i_w := 0; i_w < w; i_w++ {
		for i_h := 0; i_h < h; i_h++ {
			fmt.Print(a[i_h][i_w])
			if i_h+1 != h {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
