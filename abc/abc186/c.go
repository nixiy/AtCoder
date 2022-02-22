package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	sevenCount := 0

	for i := 1; i <= N; i++ {
		if strings.Contains(fmt.Sprintf("%d", i), "7") || strings.Contains(fmt.Sprintf("%o", i), "7") {
			sevenCount++
		}
	}

	fmt.Println(N - sevenCount)
}
