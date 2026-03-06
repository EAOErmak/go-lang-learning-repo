package main

import "fmt"

func fillSlice() []int {
	var myInpute int
	mySlice := []int{}
	for {
		fmt.Scanln(&myInpute)
		if myInpute == 111 {
			break
		}
		mySlice = append(mySlice, myInpute)
	}
	return mySlice
}
