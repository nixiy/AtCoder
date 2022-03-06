package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := ns()
	if strings.Contains(strings.Repeat("oxx", 100000), s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
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
