package main

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxValue(hash map[byte]int) int {
	res := 0
	for _, v := range hash {
		res = max(res, v)
	}
	return res
}

func characterReplacement(s string, k int) int {
	maxLen := 0
	hash := make(map[byte]int)

	left := 0
	for right := 0; right < len(s); right++ {
		hash[s[right]]++

		windowLen := right - left + 1

		statement := windowLen - maxValue(hash)

		if statement <= k {
			maxLen = max(maxLen, windowLen)
		} else {
			hash[s[left]]--
			left++
		}

	}

	return maxLen
}

/*
time complexity: O(26*n) = O(n), n - length string s
space: O(N), we created hash
*/
