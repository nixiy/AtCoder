package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = nextInt()
	}

	current := 0
	for i := 0; i < 3; i++ {
		current = a[current]
	}

	fmt.Println(current)
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func init() {
	sc.Split(bufio.ScanWords)
}
