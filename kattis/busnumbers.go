package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/**
 * https://open.kattis.com/contests/e3hoii/problems/busnumbers
 */
func main564() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	in.Scan()
	N, _ := strconv.Atoi(in.Text())
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		in.Scan()
		arr[i], _ = strconv.Atoi(in.Text())
	}

	fmt.Println(Compress(arr))
}

func Compress(arr []int) string {
	sort.Ints(arr)
	buf := bytes.Buffer{}
	arrLen := len(arr)

	for i := 0; i < arrLen; {
		offset := i
		for (offset+1 < arrLen) && (arr[offset] == arr[offset+1]-1) {
			offset++
		}

		if i == offset {
			buf.WriteString(fmt.Sprintf("%d ", arr[i]))
		} else if offset-i == 1 {
			buf.WriteString(fmt.Sprintf("%d ", arr[i]))
			offset--
		} else {
			buf.WriteString(fmt.Sprintf("%d-%d ", arr[i], arr[offset]))
		}

		i += offset - i + 1
	}

	return strings.TrimSpace(buf.String())
}
