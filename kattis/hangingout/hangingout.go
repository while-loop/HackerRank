package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)

type sem struct {
	max    int
	denies int
	cur    int
}

func newSem(max int) *sem {
	return &sem{max: max}
}

func (s *sem) change(x int) {
	if s.cur+x > s.max {
		s.denies++
		return
	}

	s.cur += x
}

/**
 * https://open.kattis.com/problems/hangingout
 */
func main() {
	in.Split(bufio.ScanWords)
	max, lines := nextInt(), nextInt()
	mSem := newSem(max)

	for i := 0; i < lines; i++ {
		action, n := next(), nextInt()
		switch action {
		case "enter":
			mSem.change(n)
		case "leave":
			mSem.change(-n)
		}
	}

	fmt.Println(mSem.denies)
}

func next() string {
	in.Scan()
	return in.Text()
}

func nextInt() int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp
}
