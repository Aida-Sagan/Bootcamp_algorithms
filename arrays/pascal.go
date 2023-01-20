package main

import "fmt"

func generate(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < len(res[i]); j++ {
			if j != 0 && i != j {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			} else {
				res[i][j] = 1
			}
		}
	}
	return res
}

func main() {
	n := 4
	fmt.Println(generate((n)))
}

/*
**Time Complexity:** **O(n^2)**, where n — length of each row and col.

**Space Complexity:** **O(n^2)** because n -- col, n -- row
 */
