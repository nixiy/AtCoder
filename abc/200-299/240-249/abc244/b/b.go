package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Coordinate struct {
	x   int
	y   int
	dir int
}

const EAST = 0
const SOUTH = 1
const WEST = 2
const NORTH = 3

func main() {
	ni()
	S := ns()
	coordinate := Coordinate{}

	for _, s := range S {
		s := string(s)
		if s == "S" {
			if coordinate.dir == EAST {
				coordinate.x++
			} else if coordinate.dir == SOUTH {
				coordinate.y--
			} else if coordinate.dir == WEST {
				coordinate.x--
			} else if coordinate.dir == NORTH {
				coordinate.y++
			}
		} else { // rotate
			coordinate.dir++
			coordinate.dir %= 4 // 北→東で3->0に戻るため
		}
	}

	fmt.Println(coordinate.x, coordinate.y)
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
