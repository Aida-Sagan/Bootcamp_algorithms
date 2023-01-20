package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hash := make(map[byte]int)

	for ind, val := range s {
		hash[byte(val)]++
		hash[byte(t[ind])]--
	}

	for _, val := range t {
		if _, ok := hash[byte(val)]; ok {

			if hash[byte(val)] != 0 {
				return false
			}
		}
	}
	return true

}

func main() {

	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
