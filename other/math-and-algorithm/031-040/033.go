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
	a := Point{
		x: ni(),
		y: ni(),
	}
	b := Point{
		x: ni(),
		y: ni(),
	}
	c := Point{
		x: ni(),
		y: ni(),
	}

	ba := Point{
		x: a.x - b.x,
		y: a.y - b.y,
	}
	bc := Point{
		x: c.x - b.x,
		y: c.y - b.y,
	}
	ca := Point{
		x: a.x - c.x,
		y: a.y - c.y,
	}
	cb := Point{
		x: b.x - c.x,
		y: b.y - c.y,
	}

	pattern := 2
	if ba.x*bc.x+ba.y*bc.y < 0.0 {
		pattern = 1
	}
	if ca.x*cb.x+ca.y*cb.y < 0.0 {
		pattern = 3
	}

	ans := 0.0

	switch pattern {
	case 1:
		ans = sqrt(ba.x*ba.x + ba.y*ba.y)
		break
	case 2:
		S := abs(ba.x*ca.y - ba.y*ca.x)
		BClen := sqrt(bc.x*bc.x + bc.y*bc.y)
		ans = float64(S) / float64(BClen)
		break
	case 3:
		ans = sqrt(ca.x*ca.x + ca.y*ca.y)
		break
	}

	fmt.Printf("%.12f\n", ans)
}

// 絶対値で返す
func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func sqrt(x int) float64 { return math.Sqrt(float64(x)) }

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
