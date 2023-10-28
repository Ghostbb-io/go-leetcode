package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchRange([]int{0, 0, 1, 1, 1, 4, 5, 5}, 2))
	fmt.Println(searchRange([]int{1, 2, 3}, 3))
	fmt.Println(searchRange([]int{2, 2}, 3))

	fmt.Println(searchRange([]int{1}, 1))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if nums[r] != target {
		return []int{-1, -1}
	}
	l = r
	for (l-1 >= 0 && nums[l-1] == target) || (r+1 < len(nums) && nums[r+1] == target) {
		if l-1 >= 0 && nums[l-1] == target {
			l--
		}
		if r+1 < len(nums) && nums[r+1] == target {
			r++
		}
	}
	return []int{l, r}
}
