package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := nextStr()
	a := string(n[0])
	b := string(n[1])
	c := string(n[2])
	abc := a + b + c
	bca := b + c + a
	cab := c + a + b
	fmt.Println(atoi(abc) + atoi(bca) + atoi(cab))
}
