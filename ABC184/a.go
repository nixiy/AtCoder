package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scL = bufio.NewScanner(os.Stdin)
var scW = bufio.NewScanner(os.Stdin)

func nextLine() string {
	scL.Scan()
	return scL.Text()
}

func nextWords() string {
	scW.Scan()
	return scW.Text()
}

func nextInt() int {
	scW.Scan()
	i, e := strconv.Atoi(scW.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	scW.Split(bufio.ScanWords)

	A := make([]int, 4)
	for i := 0; i < 4; i++ {
		A[i] = nextInt()
	}
	fmt.Println(A[0]*A[3] - A[1]*A[2])
}
