package main

func Cakes(recipe, available map[string]int) int {
	// Better: count := 1e9
	count := 1_000_000_000

	for ingredient := range recipe {
		n := available[ingredient] / recipe[ingredient]
		if n < count {
			count = n
		}
	}
	return count
}
