package main

import (
	"fmt"
	"math"
	"math/big"
)

/*
Input:
2
1 10
3 5

Output:
2
3
5
7

3
5
*/
func main() {
	//test()
	//return
	var t int
	fmt.Scanf("%d", &t)
	var num1, num2 int
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &num1)
		fmt.Scanf("%d", &num2)

		for i := num1; i <= num2; i++ {
			if prime(i) {
				fmt.Println(i)
			}
		}
		fmt.Println()
	}
}

var primes = map[int]bool{}

func prime(n int) bool {
	if n == 2 {
		return true
	} else if n%2 == 0 || n == 1 {
		return false
	}

	isPrime, exists := primes[n]
	if exists {
		return isPrime
	}

	primes[n] = true
	sqrtn := math.Sqrt(float64(n))
	for i := 3; float64(i) <= sqrtn; i += 2 {
		if n%i == 0 || n%(i+2) == 0 {
			primes[n] = false
			break
		}
	}

	return primes[n]
}

func test() {
	MAX := 1000000000
	for i := 1; i <= MAX; i++ {
		mine := prime(i)
		yours := big.NewInt(int64(i)).ProbablyPrime(4)
		if i%10000 == 0 {
			fmt.Println(i)
		}
		if mine != yours {
			fmt.Println("Mismatched prime", i, mine, yours)
			break
		}
	}
}
