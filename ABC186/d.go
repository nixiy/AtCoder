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
	N := nextInt()
	total := 0

	field := make([]int, N)

	for i := 0; i < N; i++ {
		field[i] = nextInt()
	}

	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			ans := field[i] - field[j]
			if ans < 0 {
				ans *= -1
			}
			total += ans
		}
	}

	fmt.Println(total)
}
