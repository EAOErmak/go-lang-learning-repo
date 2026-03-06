package main

func mergeMaps(m1, m2 map[string]int) map[string]int {
	myMap := m1

	for key, value := range m2 {
		_, ok := myMap[key]
		if ok {
			myMap[key] += value
		} else {
			myMap[key] = value
		}
	}
	return myMap
}
