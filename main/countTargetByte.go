package main

func countTargetByte(s string, target byte) int {
	bytes := []byte(s)

	count := 0

	for _, sym := range bytes {
		if sym == target {
			count++
		}
	}

	return count
}
