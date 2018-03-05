package main

import (
	"fmt"
	hs "hackerank/strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type pair struct {
	left, right int
}

func main() {
	str_len := 0
	n_queries := 0
	fmt.Scanf("%d %d", &str_len, &n_queries)

	str := ""
	fmt.Scanf("%s", &str)

	l := min(len(str), str_len)
	str = str[0:l]

	queries := make([]pair, n_queries)
	for i := 0; i < n_queries; i++ {
		fmt.Scanf("%d %d\n", &queries[i].left, &queries[i].right)
	}
	//fmt.Printf("l => %d nq => %d str =>%s\n", l, n_queries, str)
	for _, v := range queries {
		//result := hs.ArrCountUniqueSubstrings(str, v.left, v.right)
		result := hs.ForCountUniqueSubstrings(str, v.left, v.right)
		fmt.Printf("%v\n", result)
	}
}
