package main

import "fmt"

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1) //O(n)

	for i := 0; i < rowIndex+1; i++ { //O(n)
		for j := i; j >= 0; j-- { // O(n)
			if i != j && j != 0 {
				row[j] = row[j-1] + row[j]
			} else {
				row[j] = 1
			}
		}
	}
	return row
}

func main() {
	ind := 3
	fmt.Println(getRow(ind))
}

/*
Time complexity: O(n^2), where n -- length row and col

space complexity: O(n), because we create additional space for n elements
*/
