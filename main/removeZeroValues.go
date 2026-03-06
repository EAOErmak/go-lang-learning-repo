package main

func removeZeroValues(m map[string]int) {
	for key, value := range m {
		if value == 0 {
			delete(m, key)
		}
	}
}
