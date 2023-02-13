package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, target int
	Ai, AjMin := 0, math.MaxInt
	fmt.Fscan(in, &n, &target)
	Aslice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &Aslice[i])
	}

Loop:
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if Aslice[i]+Aslice[j] < target {
				Ai, AjMin = i, j
				continue Loop
			}
		}
	}

	if Ai == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(Ai+1, AjMin+1)
	}
}
