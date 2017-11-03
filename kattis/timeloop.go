package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

/**
 * https://open.kattis.com/contests/e3hoii/problems/timeloop
 */
func main894() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	in.Scan()
	N, _ := strconv.Atoi(in.Text())
	buf := bytes.Buffer{}
	for i := 1; i <= N; i++ {
		buf.WriteString(fmt.Sprintf("%d Abracadabra\n", i))
	}

	fmt.Print(buf.String())
}
