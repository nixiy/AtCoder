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

	x := nextInt()
	if x >= 0 {
		fmt.Println(x)
	} else {
		fmt.Println(0)
	}
}
