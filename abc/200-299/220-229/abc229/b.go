package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	A, B := ni(), ni()

	for A != 0 && B != 0 {
		if (A%10)+(B%10) >= 10 {
			fmt.Println("Hard")
			return
		}
		A /= 10
		B /= 10
	}

	fmt.Println("Easy")
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
