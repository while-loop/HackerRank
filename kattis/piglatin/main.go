package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// https://open.kattis.com/contests/bfe8oz/problems/piglatin
func main() {
	in.Split(bufio.ScanLines)

	for in.Scan() {
		fmt.Println(pigLatin(in.Text()))
	}
}

var in = bufio.NewScanner(os.Stdin)

func pigLatin(line string) string {
	out := bytes.NewBufferString("")

	for _, word := range strings.Split(line, " ") {
		if isVowel(word[0]) {
			word = vowel(word)
		} else {
			word = conso(word)
		}

		out.WriteString(word + " ")
	}

	return strings.TrimSpace(out.String())
}

func conso(word string) string {
	idx := 0
	for i := range word {
		if isVowel(word[i]) {
			idx = i
			break
		}
	}

	return word[idx:] + word[:idx] + "ay"
}

func vowel(word string) string {
	return word + "yay"
}

func isVowel(u uint8) bool {
	switch u {
	case 'a', 'e', 'i', 'o', 'u', 'y':
		return true
	default:
		return false
	}
}
