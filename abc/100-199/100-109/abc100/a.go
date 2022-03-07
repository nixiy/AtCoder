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
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	a, b := nextInt(), nextInt()
	if a <= 8 && b <= 8 {
		fmt.Println("Yay!")
	} else {
		fmt.Println(":(")
	}
}
