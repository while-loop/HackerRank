package main

import (
	"fmt"
	"reflect"
)

func main() {
	cases := []struct {
		in  string
		exp []string
	}{
		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"abc", "acb", "bac", "bca", "cba", "cab"}},
	}

	for _, tc := range cases {
		act := permutation(tc.in)
		if reflect.DeepEqual(tc.exp, act) {
			fmt.Printf("pass\n%v == %v\n", tc.exp, act)
		} else {
			fmt.Printf("fail:\npermutation(%v) = %v - want: %v\n", tc.in, act, tc.exp)
		}
		fmt.Println()
	}
}

func permutation(str string) []string {
	return permutationR(str, 0)
}

func permutationR(str string, idx int) []string {
	if idx >= len(str) {
		return []string{str}
	}

	ret := make([]string, 0)
	for i := idx; i < len(str); i++ {
		ret = append(ret, permutationR(swap(str, idx, i), idx+1)...)
	}

	return ret
}

func swap(s string, a, b int) string {
	str := []byte(s)
	str[a], str[b] = str[b], str[a]
	return string(str)
}
