package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	fact := func(n int) int {
		result := 1
		for i := 2; i <= n; i++ {
			result *= i
		}
		return result
	}
	fmt.Println(fact(n) / fact(k) / fact(n-k))
}
