package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	yes = "YES"
	no  = "NO"
)

// https://open.kattis.com/contests/bfe8oz/problems/anothercandies
func main() {
	in.Split(bufio.ScanWords)

	N := nextInt()
	for i := 0; i < N; i++ {
		kids := nextUint64()
		cur := uint64(0)
		for k := uint64(0); k < kids; k++ {
			cur = (cur + nextUint64()) % kids
		}

		if cur == 0 {
			fmt.Println(yes)
		} else {
			fmt.Println(no)
		}
	}
}

var in = bufio.NewScanner(os.Stdin)

func nextInt() int {
	in.Scan()
	tmp, err := strconv.ParseInt(strings.Replace(in.Text(), ",", "", -1), 10, 0)
	if err != nil {
		panic(err)
	}
	return int(tmp)
}

func nextUint64() uint64 {
	in.Scan()
	tmp, err := strconv.ParseUint(strings.Replace(in.Text(), ",", "", -1), 10, 64)
	if err != nil {
		panic(err)
	}
	return tmp
}
