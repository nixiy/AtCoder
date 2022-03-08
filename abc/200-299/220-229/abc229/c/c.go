package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type kv struct {
	Umami  int
	Weight int
}

func solve(N, W int, ckv []kv) (sum int) {
	sort.Slice(ckv, func(i, j int) bool {
		return ckv[i].Umami > ckv[j].Umami
	})

	// Keyがでかい順になったので、上から貪欲に足していく
	for i := 0; i < N; i++ {
		getWeight := min(W, ckv[i].Weight) // ぎりぎり取っていい重量
		sum += getWeight * ckv[i].Umami
		W -= getWeight
		if W <= 0 {
			break
		}
	}
	return sum
}

func main() {
	N, W := ni(), ni()

	ckv := make([]kv, N)
	for i := 0; i < N; i++ {
		ckv[i] = kv{ni(), ni()}
	}
	fmt.Println(solve(N, W, ckv))
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
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
