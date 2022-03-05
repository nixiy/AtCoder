package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func solve(s string) string {
	strs := strings.Split(s, "")
	sort.Strings(strs)
	return strings.Join(strs, "")
}

func main() {
	s := ns()
	fmt.Println(solve(s))
}

var sc = bufio.NewScanner(os.Stdin)

func ns() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
