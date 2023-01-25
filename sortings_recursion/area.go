package main

import (
	"math"
)

//using recursion
func rectangle(length int, area int) []int {
	if area%length == 0 {
		return []int{area / length, length}
	}
	var res []int
	res = append(res, rectangle(length-1, area)...)
	return res
}

func constructRectangle(area int) []int {
	if area == 0 {
		return []int{0, 0}
	}
	sqrt := int(math.Sqrt(float64(area)))
	return rectangle(sqrt, area)
}
