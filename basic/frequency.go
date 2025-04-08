package basic

import (
	"fmt"
	"sort"
)

type Pair struct {
	char  rune
	count int
}

func pritnFrequency() {
	input := "mississippi"
	set := make(map[rune]int)
	for _, v := range input {
		set[v]++
	}

	var pair []Pair
	for ch, s := range set {
		pair = append(pair, Pair{char: ch, count: s})
	}

	sort.Slice(pair, func(x, y int) bool {
		return pair[x].count > pair[y].count
	})

	for _, p := range pair {
		fmt.Printf("%c: %d\n", p.char, p.count)
	}
}
