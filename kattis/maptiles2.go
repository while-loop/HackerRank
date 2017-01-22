package main

import (
	"bufio"
	"os"
	"fmt"
	"math"
)

/**
  * https://open.kattis.com/contests/e3hoii/problems/maptiles2
  */
func main231561() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	in.Scan()
	coords := Key2Coord(in.Text())

	fmt.Printf("%d %d %d", coords.level, coords.x, coords.y)
}

func Key2Coord(key string) Coord {
	coord := Coord{level:len(key)}

	x, y := 0.0, 0.0
	for l := coord.level; l > 0; l-- {
		num := float64(key[coord.level - l] - '0')
		value := math.Pow(2, float64(l - 1))
		switch num {
		case 0:
			break
		case 1:
			x += value
			break
		case 2:
			y += value
			break
		case 3:
			x += value
			y += value
			break
		}
	}

	coord.x = int(x)
	coord.y = int(y)
	return coord
}

type Coord struct {
	level, x, y int
}
