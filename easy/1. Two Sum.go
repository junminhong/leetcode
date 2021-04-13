package main

import "fmt"

func main() {
	numbs := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbs, target))
}

func twoSum(numbs []int, target int) []int {
	for index, value := range numbs {
		for indexTwo, valueTwo := range numbs {
			if index != indexTwo {
				if (value + valueTwo) == target {
					return []int{index, indexTwo}
				}
			}
		}
	}
	return nil
}
