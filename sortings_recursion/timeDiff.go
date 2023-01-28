package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func findMinDifference(timePoints []string) int {
	minutes := make([]bool, 24*60)
	min := 1440

	for _, time := range timePoints {
		parses := strings.Split(time, ":")

		hour, _ := strconv.Atoi(parses[0])
		minute, _ := strconv.Atoi(parses[1])

		newMinute := hour*60 + minute

		//if there is 2 equal time => return 0.
		if minutes[newMinute] {
			return 0
		}
		minutes[newMinute] = true
	}

	first, last := -1, -1
	for i := 0; i < len(minutes); i++ {

		if minutes[i] {
			if last != -1 {
				min = Min(min, i-last)
			}
			last = i

			if first == -1 {
				first = i
			}
		}

	}
	return Min(min, len(minutes)-last+first)
}

func main() {
	n := []string{"23:59", "00:00"}
	fmt.Println(findMinDifference(n))
}
