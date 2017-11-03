package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)
var FEE = 0.10

/**
 * https://open.kattis.com/contests/ze6qt6/problems/notamused
 */
func main() {
	in.Split(bufio.ScanWords)

	ppl := map[string]int{}
	day := 1
	for {
		word := Next()
		if word == "" {
			break
		}

		switch word {
		case "OPEN":
			ppl = make(map[string]int)
			break
		case "ENTER":
			ppl[Next()] -= NextInt()
			break
		case "EXIT":
			ppl[Next()] += NextInt()
			break
		case "CLOSE":
			printPrices(day, ppl)
			day++
			break
		}
	}
}
func printPrices(day int, ppl map[string]int) {
	var keys []string
	for k := range ppl {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// To perform the opertion you want
	if day != 1 {
		fmt.Println()
	}
	fmt.Println("Day", day)
	for _, k := range keys {
		fmt.Printf("%s $%.2f\n", k, float64(ppl[k])*FEE)
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
