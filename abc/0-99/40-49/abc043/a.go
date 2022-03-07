package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := ni()
	fmt.Println(n * (n + 1) / 2)
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func init() {
	sc.Split(bufio.ScanWords)
}
