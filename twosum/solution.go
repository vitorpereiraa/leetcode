package main

import (
	"fmt"
)

func main() {
	var slice []int = []int{2, 7, 11, 15}
	var target int = 9

	fmt.Println(solution(slice, target)) // solution [0, 1]

	var slice2 []int = []int{3, 2, 4}
	var target2 int = 6

	fmt.Println(solution(slice2, target2)) // solution [1, 2]

	var slice3 []int = []int{3, 3}
	var target3 int = 6

	fmt.Println(solution(slice3, target3)) // solution [0, 1]
}

// O(n)
func solution(nums []int, target int) []int {
	aux := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		var diff int = target - nums[i]

		_, exists := aux[diff]

		if exists {
			return []int{aux[diff], i}
		} else {
			aux[nums[i]] = i
		}
	}

	return []int{}
}

// O(n2)
func solution2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
