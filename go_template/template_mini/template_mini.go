package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func ns() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
}
