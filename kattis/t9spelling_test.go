package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCase0(t *testing.T) {
	in := "hi"
	exp := "44 444"
	assert.Equal(t, exp, Word2Nums(in), "FeelsUnluckyMan")
}

func TestCase1(t *testing.T) {
	in := "yes"
	exp := "999337777"
	assert.Equal(t, exp, Word2Nums(in), "FeelsUnluckyMan")
}

func TestCase2(t *testing.T) {
	in := "foo  bar"
	exp := "333666 6660 022 2777"
	assert.Equal(t, exp, Word2Nums(in), "FeelsUnluckyMan")
}

func TestCase3(t *testing.T) {
	in := "hello world"
	exp := "4433555 555666096667775553"
	assert.Equal(t, exp, Word2Nums(in), "FeelsUnluckyMan")
}

func TestCase4(t *testing.T) {
	in := "YesTERDAY I went   to the  store"
	exp := "99933777783377732999044409336680 0 086660844330 07777866677733"
	assert.Equal(t, exp, Word2Nums(in), "FeelsUnluckyMan")
}
