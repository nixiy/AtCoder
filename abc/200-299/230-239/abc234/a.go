package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func f(t int) int {
	return pow(t, 2) + 2*t + 3
}

func main() {
	t := nextInt()
	fmt.Println(f(f(f(t)+t) + f(f(t))))
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func init() {
	sc.Split(bufio.ScanWords)
}
