package main

import (
	"fmt"
	"strings"
)

func main() {
	var num string
	fmt.Scan(&num)
	if isMagic(num, "1") || isMagic(num, "14") || isMagic(num, "144") {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func isMagic(num, prefix string) bool {
	if num == "" {
		return true
	}
	num, found := strings.CutPrefix(num, prefix)
	if !found {
		return false
	}
	if isMagic(num, "1") || isMagic(num, "14") || isMagic(num, "144") {
		return true
	}
	return false
}
