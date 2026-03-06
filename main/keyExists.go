package main

func keyExists(m map[string]int, keystring string) bool {
	_, ok := m[keystring]
	return ok
}
