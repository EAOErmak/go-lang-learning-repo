package main

func mostFrequentChar(s string) byte {
	myMap := map[byte]int{}

	for i := 0; i < len(s); i++ {
		myMap[s[i]]++
	}

	var mfreq byte

	max := 0

	for i := 0; i < len(s); i++ {
		if myMap[s[i]] > max {
			max = myMap[s[i]]
			mfreq = s[i]
		}
	}
	return mfreq
}
