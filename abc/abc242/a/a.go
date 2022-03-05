package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(a, b, c, x int) float64 {
	nin := b - a

	if a >= x {
		return 1
	} else if a < x && x <= b {
		return float64(c) / float64(nin)
	} else {
		return 0
	}
}

func main() {
	a, b, c, x := ni(), ni(), ni(), ni()
	fmt.Println(solve(a, b, c, x))
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func init() {
	sc.Split(bufio.ScanWords)
}
