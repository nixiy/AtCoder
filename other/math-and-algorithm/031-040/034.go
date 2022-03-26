package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Point struct {
	x int
	y int
}

func main() {
	N := ni()
	p := make([]Point, N)

	for i := 0; i < N; i++ {
		p[i].x = ni()
		p[i].y = ni()
	}

	minDis := 100000000.0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i != j {
				chmin(&minDis, euclidDistance(p[i].x, p[i].y, p[j].x, p[j].y))
			}
		}
	}

	fmt.Println(minDis)
}

// 2次元ユークリッド距離
func euclidDistance(px, py, qx, qy int) float64 {
	pxf := float64(px)
	pyf := float64(py)
	qxf := float64(qx)
	qyf := float64(qy)
	return math.Sqrt((pxf-qxf)*(pxf-qxf) + (pyf-qyf)*(pyf-qyf))
}

// ex chmin(&p, v)
func chmin(p *float64, v float64) bool {
	if *p > v {
		*p = v
		return true
	}
	return false
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
