package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	result := float64(0)
	for i := 0; i <= n; i++ {
		result += math.Pow(-1.0, float64(i)) / float64(2*i+1)
	}
	fmt.Println(4.0 * result)
}
