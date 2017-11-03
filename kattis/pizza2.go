package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

/**
 * https://open.kattis.com/contests/e3hoii/problems/pizza2
 */
func main212() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	in.Scan()
	pizzaR, _ := strconv.ParseFloat(in.Text(), 64)
	in.Scan()
	crustR, _ := strconv.ParseFloat(in.Text(), 64)

	cheeseA := math.Pow(float64(pizzaR-crustR), 2.0)
	pizzaA := math.Pow(float64(pizzaR), 2.0)

	fmt.Printf("%.6f", cheeseA/pizzaA*100)
}
