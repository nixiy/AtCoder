package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	S := ns()
	a, _ := strconv.Atoi(string(S[0]))
	b, _ := strconv.Atoi(string(S[2]))
	fmt.Println(a * b)
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
