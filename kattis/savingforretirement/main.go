package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://open.kattis.com/contests/dasomz/problems/savingforretirement
func main() {
	in.Split(bufio.ScanWords)
	fmt.Println(breakEven(nextInt(), nextInt(), nextInt(), nextInt(), nextInt()))
}

func breakEven(bAge, bRetire, bSave, aAge, aSave int) int {
	bFinal := float64(bSave * (bRetire - bAge))
	aYears := bFinal / float64(aSave)
	return int(aYears+float64(aAge)) + 1
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
