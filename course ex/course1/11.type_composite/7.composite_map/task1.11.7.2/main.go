package main

import "fmt"

func mergeMaps(map1, map2 map[string]int) map[string]int {
	resultMap := make(map[string]int)
	for key, value := range map1 {
		resultMap[key] = value
	}

	for key, value := range map2 {
		resultMap[key] = value
	}

	return resultMap
}

func main() {
	map1 := map[string]int{"apple": 3, "banana": 2}
	map2 := map[string]int{"orange": 5, "grape": 4}
	mergedMap := mergeMaps(map1, map2)
	for key, value := range mergedMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
