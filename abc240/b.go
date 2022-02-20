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
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	input := make([]int, n)
	uniq := []int{}

	for i := 0; i < n; i++ {
		input[i] = nextInt()
	}

	m := make(map[int]bool, n)

	for _, elem := range input {
		if !m[elem] {
			m[elem] = true
			uniq = append(uniq, elem)
		}
	}

	fmt.Println(len(uniq))
}
