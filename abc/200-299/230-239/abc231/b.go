package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Candidate struct {
	name string
	cnt  int
}

func main() {
	N := ni()

	nameMap := make(map[string]int)
	for i := 0; i < N; i++ {
		s := ns()
		nameMap[s]++
	}
	max := Candidate{
		name: "",
		cnt:  -1,
	}
	for s, i := range nameMap {
		if max.cnt < i {
			max = Candidate{
				name: s,
				cnt:  i,
			}
		}
	}
	fmt.Println(max.name)
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func ns() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
