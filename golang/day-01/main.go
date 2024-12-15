package main

import (
	"slices"
	"sort"
)

func reduce[T1, T2 any](s []T1, accumulator T2, reducer func(T2, T1) T2) T2 {
	r := accumulator
	for _, v := range s {
		r = reducer(r, v)
	}
	return r
}

func PrepareGifts(gifts []int) []int {
	if len(gifts) == 0 {
		return []int{}
	}

	uniqueGifts := reduce(gifts, []int{}, func(accumulator []int, element int) []int {
		if slices.Contains(accumulator, element) {
			return accumulator
		}
		accumulator = append(accumulator, element)
		return accumulator
	})

	sort.Slice(uniqueGifts, func(i, j int) bool { return uniqueGifts[i] < uniqueGifts[j] })

	return uniqueGifts
}
