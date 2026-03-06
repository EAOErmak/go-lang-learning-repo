package main

func countFrequency(nums []int) map[int]int {
	myMap := map[int]int{}

	for _, num := range nums {
		myMap[num]++
	}

	return myMap
}
