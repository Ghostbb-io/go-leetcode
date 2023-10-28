package main

import (
	"fmt"
	"sort"
)

func main() {
}

func minOperations(nums []int) int {
	n := len(nums)
	ans := 100000000
	// sort
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	// unique
	nums = unique(nums)

	// windows
	// 因為排序+唯一，所以max - min = 中間所需要的數量
	// max - min < n，如果大於n代表中間所取的的數量>規定數量
	// ans = len(nums) - n
	for i := 0; i < len(nums); i++ {
		j := i + 1
		for j < len(nums) && nums[j]-nums[i] < n {
			j++
		}
		ans = min(ans, n-(j-i))
	}
	return ans
}

func unique(nums []int) []int {
	keys := make(map[int]bool)
	result := make([]int, 0)
	for _, n := range nums {
		if _, ok := keys[n]; !ok {
			result = append(result, n)
			keys[n] = true
		}
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}