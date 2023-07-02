package main

func CountFruits(fruits []string) map[string]int {
	m := make(map[string]int, len(fruits))
	for _, fruit := range fruits {
		m[fruit]++
	}
	return m
}

func main() {
	fruits := []string{"apple", "orange", "banana"}
	m := CountFruits(fruits)
	for fruit, count := range m {
		println(fruit, count)
	}
}
