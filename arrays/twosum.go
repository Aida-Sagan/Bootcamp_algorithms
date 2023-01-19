package main

import "fmt"

func twoSum(array []int, target int) []int {
	hash := make(map[int]int)
	res := make([]int, 0, 2)

	for ind, val := range array {
		if _, ok := hash[val]; !ok {
			hash[val] = ind
		}
	}

	for ind, val := range array {
		diff := target - val
		if _, ok := hash[diff]; ok {
			if hash[diff] != ind {
				res = append(res, ind, hash[diff])
				return res
			}
		}
	}
	return res
}

func main() {
	var arr []int = []int{2, 7, 11, 15}
	tar := 9

	fmt.Println(twoSum(arr, tar))
}
