package main

func isPalindromeUnicode(s string) bool {
	runes := []rune(s)

	size := len(runes)

	for i := 0; i < size/2; i++ {
		if runes[i] != runes[size-i-1] {
			return false
		}
	}

	return true
}
