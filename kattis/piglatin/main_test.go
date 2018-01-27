package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	cases := []struct {
		line, exp string
	}{
		{"i cant speak pig latin", "iyay antcay eakspay igpay atinlay"},
		{"the quick brown fox", "ethay uickqay ownbray oxfay"},
		{"jumps over the lazy dog", "umpsjay overyay ethay azylay ogday"},
		{"and ordinary foxes", "andyay ordinaryyay oxesfay"},
		{"dont jump over lazy dogs", "ontday umpjay overyay azylay ogsday"},
	}

	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tc.exp, pigLatin(tc.line))
		})
	}
}
