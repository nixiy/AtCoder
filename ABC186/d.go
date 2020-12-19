package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	pastSum := 0

	field := make([]int, N)

	for i := 0; i < N; i++ {
		field[i] = nextInt()
	}

	sort.Ints(field)

	for j := 0; j < N; j++ {
		total += field[j] * j
		total -= pastSum
		pastSum += field[j]
	}

	fmt.Println(total)
}
