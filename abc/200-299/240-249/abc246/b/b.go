package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	x, y := ni(), ni()
	e := euclidDistance(0, 0, x, y)
	distanceMagnification := 1 / float64(e)
	fmt.Println(float64(x)*distanceMagnification, float64(y)*distanceMagnification)
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

// 2次元ユークリッド距離
func euclidDistance(px, py, qx, qy int) float64 {
	pxf := float64(px)
	pyf := float64(py)
	qxf := float64(qx)
	qyf := float64(qy)
	return math.Sqrt((pxf-qxf)*(pxf-qxf) + (pyf-qyf)*(pyf-qyf))
}

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
