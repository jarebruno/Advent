package main

import (
	"fmt"
	"strings"
)

func myReduce[T1, T2 any](s []T1, accumulator T2, reducer func(T2, T1) T2) T2 {
	r := accumulator
	for _, v := range s {
		r = reducer(r, v)
	}
	return r
}

func myMap[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func CreateFrame(names []string) string {
	if len(names) == 0 {
		panic("Can not create frame for empty names")
	}

	maxLength := myReduce(names, 0, func(accumulator int, element string) int {
		if len(element) > accumulator {
			return len(element)
		}
		return accumulator
	})

	innerFrameLines := myMap(names, func(element string) string {
		a := fmt.Sprintf("* %s%s *\n", element, strings.Repeat(" ", maxLength-len(element)))
		return a
	})

	externalFrameLines := fmt.Sprintf("%s\n", strings.Repeat("*", maxLength+4))

	return externalFrameLines + strings.Join(innerFrameLines, "") + externalFrameLines
}
