package main

import (
	"fmt"
)

/*
algorithm:

1. create hast table, where key is element from array nums1, value has false/true.
True — the hash element is contained in the array result, false — vice versa.

2. traverse through array nums1. Check: the element in nums1 is contained in hash table, if false — add element in hash.

3. traverse through array nums2. Check: the element in nums2 is contained in hash table and value of hash
is false ⇒ append in array res then change value on true.

**Time Complexity**: O(n + m),  where n — length of array1, m — length of array2.

**Space Complexity**: O(n), because we use additional space — hash table for saving elements.

*/

func intersection(nums1 []int, nums2 []int) []int {
	hash := make(map[int]bool)
	var res []int

	for _, el := range nums1 {
		if _, ok := hash[el]; !ok {
			hash[el] = false //false --- element is  not contained in res
		}
	}

	for _, el := range nums2 {
		if _, ok := hash[el]; ok {
			if !hash[el] {
				res = append(res, el)
				hash[el] = true
			}
		}
	}
	return res

}

func main() {
	var arr []int = []int{1, 2, 2}
	var arr2 []int = []int{2, 2, 2, 3, 4}

	fmt.Println(intersection(arr, arr2))
}
