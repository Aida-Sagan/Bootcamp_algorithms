package main

import "fmt"

/*

i+1 < len(nums) --- чтобы не выйти за пределы массива

nums[ind] = nums[len(nums)-1] --- последний элемент жб записывается в массив


**Time Complexity**: O(n), where n — length of array. We have one loop ‘for’ which traverse through array.

**Space Complexity**: O(1) * n = O(1). We dont use additional space, because the array was given to us initially.

*/

func removeDuplicates(nums []int) int {
	var ind int = 0

	for i := 0; i < len(nums)-1; i++ {
		if i+1 < len(nums) && nums[i] != nums[i+1] {
			nums[ind] = nums[i]
			ind++
		}
	}
	nums[ind] = nums[len(nums)-1]
	return ind + 1
}

func main() {
	var arr []int = []int{1, 1, 2}
	//var arr []int = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr))
}
