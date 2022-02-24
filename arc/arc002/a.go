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
	y := nextInt()
	uruu := false

	if y%4 == 0 {
		uruu = true
	}
	if y%100 == 0 {
		uruu = false
	}
	if y%400 == 0 {
		uruu = true
	}

	if uruu {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
