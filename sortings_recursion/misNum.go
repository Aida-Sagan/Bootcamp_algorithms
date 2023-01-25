package main

import "fmt"

func missingNumber(nums []int) int {
	//cycle sort
	var i int

	for i < len(nums) {
		prevSym := nums[i]
		if nums[i] < len(nums) && nums[i] != nums[prevSym] {
			//fmt.Println("nums[i]:", nums[i])
			swap(nums, i, prevSym)

		} else {
			i++
		}
	}
	//find the missing number
	for j := 0; j < len(nums); j++ {
		if nums[j] != j {
			return j
		}
	}
	return len(nums)
}

func swap(nums []int, i int, num int) {
	temp := nums[i]
	nums[i] = nums[num]
	nums[num] = temp
}

func main() {
	var n []int = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(n))
}
