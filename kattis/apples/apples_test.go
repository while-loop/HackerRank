package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		in, out [][]byte
	}{
		{
			str2arr(3, 3, "aaa#....#"),
			str2arr(3, 3, "a..#.a.a#"),
		},
		{
			str2arr(4, 5, "aaa.aaa.a.a.a.....a."),
			str2arr(4, 5, ".....a....aaaa.aaaaa"),
		},
	}
	for _, tc := range tcs {
		solve(tc.in)
		assert.True(t, reflect.DeepEqual(tc.out, tc.in))
	}
}

func str2arr(rows, cols int, arrStr string) [][]byte {
	arr := make([][]byte, rows)
	for r := 0; r < rows; r++ {
		arr[r] = make([]byte, cols)
		for c := 0; c < cols; c++ {
			arr[r][c] = arrStr[r*cols+c]
		}
	}
	return arr
}
