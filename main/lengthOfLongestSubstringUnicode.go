package main

func lengthOfLongestSubstringUnicode(s string) int {
	maxSubStr := 0
	counter := 0
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] == runes[j] {
				break
			}
			counter++
		}
		if counter > maxSubStr {
			maxSubStr = counter
		}
		counter = 0
	}

	return maxSubStr
}
