package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	k := nextInt()

	// 文字列の2進数変換を行う
	s := itob(k)

	// 2進数の1部分を2に置き換える
	s = strings.ReplaceAll(s, "1", "2")

	fmt.Println(s)
}

// 10進数xを文字列2進数で返す
func itob(x int) string {
	s := ""
	for x > 0 {
		s = strconv.Itoa(x%2) + s
		x /= 2
	}
	return s
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func init() {
	sc.Split(bufio.ScanWords)
}
