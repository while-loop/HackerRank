package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBalanced(t *testing.T) {
	inputs := []string{"()[]{}", "([{}])", "([]{})", "([{}])", "([()])", "{[{{[]}}]}"}
	for _, ele := range inputs {
		assert.Equal(t, "True", IsBalanced(ele), "String was not balanced", ele)
	}
}

func TestUnbalanced(t *testing.T) {
	inputs := []string{"([)]", "([]", "[])", "([})", "([()]"}
	for _, ele := range inputs {
		assert.Equal(t, "False", IsBalanced(ele), "String was not balanced", ele)
	}
}
