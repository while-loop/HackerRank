package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type point struct {
	x, y float64
}

type line struct {
	p1, p2 point
}

// https://en.wikipedia.org/wiki/Line–line_intersection#Given_two_points_on_each_line
func (t line) intersects(o line) bool {
	x1, y1 := t.p1.x, t.p1.y
	x2, y2 := t.p2.x, t.p2.y
	x3, y3 := o.p1.x, o.p1.y
	x4, y4 := o.p2.x, o.p2.y
	denom := ((x1 - x2) * (y3 - y4)) - ((y1 - y2) * (x3 - x4))
	if denom == 0 {
		return false
	}

	x := (((y1 - y3) * (x4 - x3)) - ((x1 - x3) * (y4 - y3))) / denom
	y := (((y1 - y3) * (x2 - x1)) - ((x1 - x3) * (y2 - y1))) / denom
	return (0 <= x && x <= 1) && (0 <= y && y <= 1)
}

type triangle struct {
	sides [3]line
}

// https://open.kattis.com/contests/bfe8oz/problems/countingtriangles
func main() {
	in.Split(bufio.ScanWords)

	for ls := nextInt(); ls > 0; ls = nextInt() {
		var lines []line
		for i := 0; i < ls; i++ {
			lines = append(lines, line{point{nextFloat64(), nextFloat64()}, point{nextFloat64(), nextFloat64()}})
		}
		fmt.Println(len(makeTriangles(lines)))
	}
}

func makeTriangles(lines []line) []*triangle {
	var tris []*triangle
	for i := 0; i < len(lines); i++ {
		for j := i; j < len(lines); j++ {
			for k := j; k < len(lines); k++ {
				if lines[i].intersects(lines[j]) && lines[j].intersects(lines[k]) && lines[k].intersects(lines[i]) {
					tris = append(tris, &triangle{[3]line{lines[i], lines[j], lines[k]}})
				}
			}
		}
	}
	return tris
}

var in = bufio.NewScanner(os.Stdin)

func nextInt() int {
	in.Scan()
	tmp, err := strconv.Atoi(in.Text())
	if err != nil {
		panic(err)
	}
	return tmp
}

func nextFloat64() float64 {
	in.Scan()
	tmp, err := strconv.ParseFloat(in.Text(), 64)
	if err != nil {
		panic(err)
	}
	return tmp
}
