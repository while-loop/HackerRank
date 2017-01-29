package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var in = bufio.NewScanner(os.Stdin)

/**
  *
  */
func main() {
	in.Split(bufio.ScanWords)

	N := NextInt()
	M := NextInt()
	dict := map[string]int{}
	for i := 0; i < N; i++ {
		dict[Next()] = NextInt()
	}

	for i := 0; i < M; i++ {
		pts := 0
		for {
			value := Next()
			if value == "." {
				break
			}
			pts += dict[value]
		}
		fmt.Println(pts)
	}
}

func NextInt() int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp
}

func NextFloat64() float64 {
	in.Scan()
	tmp, _ := strconv.ParseFloat(in.Text(), 64)
	return tmp
}

func NextLine() string {
	in.Scan()
	return in.Text()
}

func Next() string {
	in.Scan()
	return in.Text()
}
