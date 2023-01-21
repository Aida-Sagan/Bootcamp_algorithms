package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	hash := make(map[[26]int][]string) // key -- array with frequency of each symbol word (counter), value --- array of word-anagram(with equal frequency)

	for _, w := range strs {
		var key [26]int
		for _, sym := range w {
			key[sym-'a']++ //find the frequency letter in word
		}
		hash[key] = append(hash[key], w)
	}

	var res [][]string
	for _, val := range hash {
		res = append(res, val)
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams((strs)))
}

//OUTPUT: [[eat tea ate] [tan nat] [bat]]
/*
**Time Complexity: O(26 * n * m) = O(n * m),  where n — length array, m — length word, 26 — number of characters between ‘a’ - ‘z’**

**Space Complexity: O(n*m), because we create hash map with size n*m.**

 */
