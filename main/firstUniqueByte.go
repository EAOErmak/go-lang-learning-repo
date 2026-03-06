package main

func firstUniqueByte(s string) int {
	arr := make([]int, 256)

	bytes := []byte(s)

	for _, sym := range bytes {
		arr[sym]++
	}

	idx := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			return idx
		} else {
			for arr[i] > 0 {
				arr[i]--
				idx++
			}
			continue
		}
	}

	return 0
}
