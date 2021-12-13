package day03

import (
	"fmt"
	"sort"
)

type EarnIndex struct {
	index, earned int
}

func deleteAndEarn(nums []int) int {
	hist := historgram(nums)
	keys := make([]int, len(hist))
	i := 0
	for k := range hist {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	fmt.Printf("hist %v %v\n", hist, keys)
	if len(keys) < 2 {
		return hist[keys[0]]
	}
	if keys[0]+1 != keys[1] {
		memo := make(map[EarnIndex]int)
		return maxEarn(hist, keys, 0, memo)
	}
	memo1 := make(map[EarnIndex]int)
	memo2 := make(map[EarnIndex]int)
	return Max(
		maxEarn(hist, keys, 1, memo2),
		maxEarn(hist, keys, 0, memo1),
	)
}

func maxEarn(hist map[int]int, keys []int, index int, memo map[EarnIndex]int) int {
	if index >= len(keys) {
		return 0
	}
	i := EarnIndex{index, hist[keys[index]]}
	val, ok := memo[i]
	fmt.Printf("hist %d, %v\n", keys[index], hist)
	if ok {
		fmt.Printf("val %d %d\n", keys[index], val)
		return val
	} else {
		b := 0
		if index+2 <= len(keys) {
			if keys[index]+1 == keys[index+1] {
				b = Max(
					maxEarn(hist, keys, index+2, memo),
					maxEarn(hist, keys, index+3, memo),
				)
			} else {
				b = Max(
					maxEarn(hist, keys, index+1, memo),
					maxEarn(hist, keys, index+2, memo),
				)
			}
		}
		memo[i] = hist[keys[index]] + b
		fmt.Printf("memo %d %d\n", keys[index], memo[i])
		return memo[i]
	}
}

func historgram(nums []int) map[int]int {
	h := make(map[int]int)
	for _, v := range nums {
		n, ok := h[v]
		if ok {
			h[v] = n + v
		} else {
			h[v] = v
		}
	}
	return h
}
