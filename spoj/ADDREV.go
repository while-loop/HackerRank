package main

import "fmt"

/*

Sample input:

3
24 1
4358 754
305 794

Sample output:

34
1998
1
*/
func main() {
	var n int
	fmt.Scanf("%d", &n)
	var num1, num2 int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &num1)
		fmt.Scanf("%d", &num2)

		fmt.Println(reverse(reverse(num1) + reverse(num2)))
	}
}

func reverse(num int) int {
	tmp := 0
	for num > 0 {
		tmp = (tmp * 10) + (num % 10)
		num /= 10
	}
	return tmp
}
