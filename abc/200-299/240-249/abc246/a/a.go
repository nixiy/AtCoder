package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	xMap := make(map[int]int)
	yMap := make(map[int]int)

	for i := 0; i < 3; i++ {
		xMap[ni()]++
		yMap[ni()]++
	}

	x, y := -1, -1

	for k, v := range xMap {
		if v == 1 {
			x = k
		}
	}

	for k, v := range yMap {
		if v == 1 {
			y = k
		}
	}

	fmt.Println(x, y)
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
