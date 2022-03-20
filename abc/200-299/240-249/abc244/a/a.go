package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	ni()
	s := ns()
	fmt.Println(string(s[len(s)-1]))
}

var sc = bufio.NewScanner(os.Stdin)

func ns() string        { sc.Scan(); return sc.Text() }
func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
