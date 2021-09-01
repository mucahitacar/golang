package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	SortBy_A()
}

//todo optimizing
func SortBy_A() {

	x := []string{"apple", "red", "apple", "red", "red", "red"}

	max := 0
	maxstring := ""
	count := 0
	for i := 0; i < len(x); i++ {
		for y := 0; y < len(x); y++ {
			if x[i] == x[y] {
				count++
			}
		}
		fmt.Println("count", count)
		if count > max {
			max = count
			maxstring = x[i]
		}
		count = 0

	}

	fmt.Println("max", maxstring)
	sort.Slice(x, func(i, j int) bool {
		return strings.Count(x[i], "a") > strings.Count(x[j], "a")
	})

	fmt.Println(x)
}
