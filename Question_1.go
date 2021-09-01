package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	SortBy_A()
}

func SortBy_A() {

	x := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}

	sort.Slice(x, func(i, j int) bool {
		if len(x[i]) != len(x[j]) {
			return len(x[i]) > len(x[j])
		} else {
			return x[i] < x[j]
		}

	})
	fmt.Println(x)
	sort.Slice(x, func(i, j int) bool {
		return strings.Count(x[i], "a") > strings.Count(x[j], "a")
	})

	fmt.Println(x)
}
