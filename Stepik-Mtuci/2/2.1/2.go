package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	result := 1
	for i := 2; i <= n; i++ {
		result = result % 1000007 * i
	}
	fmt.Println(result % 1000007)
}
