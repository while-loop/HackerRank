package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAlphasCase0(t *testing.T) {
	in := "All your base are belong to us."
	exp := "@11 `/0|_||Z 8@$3 @|Z3 8310[]\\[]6 ']['0 |_|$."
	assert.Equal(t, exp, ConvAlphabet(in), "FeelsUnluckyMan")
}

func TestAlphasCase1(t *testing.T) {
	in := "What's the Frequency, Kenneth?"
	exp := "\\/\\/[-]@'][''$ ']['[-]3 #|Z3(,)|_|3[]\\[](`/, |<3[]\\[][]\\[]3']['[-]?"
	assert.Equal(t, exp, ConvAlphabet(in), "FeelsUnluckyMan")
}

func TestAlphasCase2(t *testing.T) {
	in := "A new alphabet!"
	exp := "@ []\\[]3\\/\\/ @1|D[-]@83']['!"
	assert.Equal(t, exp, ConvAlphabet(in), "FeelsUnluckyMan")
}

func TestAlphasCase3(t *testing.T) {
	in := "Hello World!"
	exp := "[-]3110 \\/\\/0|Z1|)!"
	assert.Equal(t, exp, ConvAlphabet(in), "FeelsUnluckyMan")
}
