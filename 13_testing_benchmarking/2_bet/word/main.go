package word

import "strings"

// UseCount return a map with every word and string and its occurrence count
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

func Count(s string) int {
	return len(UseCount(s))
}
