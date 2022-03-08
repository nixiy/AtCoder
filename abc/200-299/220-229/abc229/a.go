package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s1 := ns()
	s2 := ns()

	if s1 == ".#" && s2 == "#." || s1 == "#." && s2 == ".#" {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
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
