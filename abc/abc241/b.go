package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	pastaNum := nextInt()
	days := nextInt()

	pastaLen := make([]int, pastaNum)
	pastaLenMap := make(map[int]int)
	for i := 0; i < pastaNum; i++ {
		pastaLen[i] = nextInt()
		pastaLenMap[pastaLen[i]]++
	}

	pastaLenDays := make([]int, days)
	for i := 0; i < days; i++ {
		pastaLenDays[i] = nextInt()
	}

	for i := 0; i < days; i++ {
		if pastaLenMap[pastaLenDays[i]] == 0 {
			no()
			return
		}
		pastaLenMap[pastaLenDays[i]]--
	}
	yes()
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func yes() {
	fmt.Println("Yes")
}

func no() {
	fmt.Println("No")
}

func init() {
	sc.Split(bufio.ScanWords)
}
