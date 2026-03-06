package main

import "fmt"

func fillMap() map[string]int {
	var myInputeKey string
	var myInputeValue int
	myMap := map[string]int{}
	for {
		fmt.Scanln(&myInputeKey, &myInputeValue)
		if myInputeKey == "111" {
			break
		}
		myMap[myInputeKey] = myInputeValue
	}
	return myMap
}
