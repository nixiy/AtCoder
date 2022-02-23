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

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	sc.Split(bufio.ScanWords)
}

// string[i]のように取得するとbyteで取得できてしまう
// 中間処理でruneを使用して部分文字を取得する
func getStrForIndex(str string, index int) string {
	rs := []rune(str)
	return string(rs[index])
}

func main() {
	inputStr := nextStr()
	first, second := nextInt(), nextInt()
	for i := 0; i < len(inputStr); i++ {
		if i == first-1 {
			fmt.Print(getStrForIndex(inputStr, second-1))
		} else if i == second-1 {
			fmt.Print(getStrForIndex(inputStr, first-1))
		} else {
			fmt.Print(getStrForIndex(inputStr, i))
		}
	}
	fmt.Println()
}
