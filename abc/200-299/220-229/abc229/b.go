package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	A, B := ni(), ni()
	minLen := len(strconv.Itoa(min(A, B)))

	for i := 0; i < minLen; i++ {
		digA := A % 10
		digB := B % 10
		if digA+digB >= 10 {
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
