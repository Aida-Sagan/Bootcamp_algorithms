package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	size := len(nums)
	left := 0
	right := size - 1
	res := make([]int, size)

	for i := 0; i < size; i++ {
		if nums[i]%2 == 0 {
			res[left] = nums[i]
			left++
		} else {
			res[right] = nums[i]
			right--
		}

	}
	return res
}

func main() {
	n := []int{3, 2, 5, 4}
	fmt.Println(sortArrayByParity(n))
}
