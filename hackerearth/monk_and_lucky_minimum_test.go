package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bufio"
	"strings"
)

func TestCase0(t *testing.T) {
	input := `10
5
27 84 50 18 70
3
19 20 19
3
11 11 27
1
26
9
17 17 17 76 17 17 17 17 17
8
12 12 12 12 12 26 12 12
7
78 34 16 59 13 13 13
3
34 34 34
9
31 53 70 21 96 89 17 88 17
5
78 21 21 21 21`

	exp := "Lucky\nUnlucky\nUnlucky\nLucky\nUnlucky\nLucky\nLucky\nLucky\nUnlucky\nUnlucky\n"

	scan := bufio.NewScanner(strings.NewReader(input))
	assert.Equal(t, exp, solve(scan), "FeelsUnluckyMan")
}