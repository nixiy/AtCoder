package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type P struct {
	x int
	y int
}

func main() {
	N := nextInt()
	p := make([]P, N)
	for i := 0; i < N; i++ {
		p[i].x = nextInt()
		p[i].y = nextInt()
	}

	maxLen := 0.0
	// 全探索する
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			chmax(&maxLen, euclidDistance(p[i].x, p[i].y, p[j].x, p[j].y))
		}
	}
	fmt.Println(maxLen)
}

// ex chmax(&P, v)
func chmax(p interface{}, v interface{}) {
	switch v.(type) {
	case int:
		a, ok := p.(*int)
		if !ok {
			return
		}
		if vv := v.(int); *a < vv {
			*a = vv
		}
	case float32:
		a, ok := p.(*float32)
		if !ok {
			return
		}
		if vv := v.(float32); *a < vv {
			*a = vv
		}
	case float64:
		a, ok := p.(*float64)
		if !ok {
			return
		}
		if vv := v.(float64); *a < vv {
			*a = vv
		}
	}
}

func euclidDistance(x1, y1, x2, y2 int) float64 {
	return sqrt(pow(x1-x2, 2) + pow(y1-y2, 2))
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

// xのy乗
func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// √s
func sqrt(s int) float64 {
	return math.Sqrt(float64(s))
}

func init() {
	sc.Split(bufio.ScanWords)
}
