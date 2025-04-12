package main

import (
	"fmt"
	"strings"
)

func reverse[T any](arr []T) {
	n := len(arr)
	for i := range n / 2 {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
}

func filterContains[T fmt.Stringer](arr []T, s string) []T {
	var filtered []T
	for _, v := range arr {
		if !strings.Contains(v.String(), s) {
			filtered = append(filtered, v)
		}
	}
    return filtered
}

func move(index *int, delta int, min int, max int) {
    *index += delta
    if *index > max {
        *index = max
    } else if *index < min {
        *index = min
    }
}
