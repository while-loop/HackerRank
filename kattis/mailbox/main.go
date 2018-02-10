package main

import (
	"bufio"
	"strconv"
	"os"
	"fmt"
)

func init() {
	in.Split(bufio.ScanWords)
}

// https://open.kattis.com/contests/bw8c4i/problems/mailbox
func main() {
	cases := nextInt()
	for i := 0; i < cases; i++ {
	    fmt.Println(totalFireworks(nextInt(), nextInt()))
	}
}

func totalFireworks(mailboxes, limit int) int {
	return (limit*(limit+1))/2
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
