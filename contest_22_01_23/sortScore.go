package main

import (
	"fmt"
	"sort"
)

func sortTheStudents(score [][]int, k int) [][]int {

	/*var exam []int
	for i := 0; i < len(score); i++ {
		exam[i] = score[i][k]
	}

	var temp int

	for i := 0; i < len(exam); i++ {
		for j := i + 1; j < len(exam); j++ {

			if exam[j] > exam[i] {
				temp = exam[j]
				exam[j] = exam[i]
				exam[i] = temp
			}

			for l := 0; l < len(score[0]); l++ {
				temp = score[i][l]
				score[i][l] = score[j][l]
				score[j][l] = score[i][l]
			}
		}
	}
	return score*/

	sort.SliceStable(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})
	return score
}

func main() {
	score := [][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}
	k := 2
	fmt.Println(sortTheStudents(score, k))
}
