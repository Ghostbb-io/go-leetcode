package _1658

func minOperations(nums []int, x int) int {
	target := sum(nums) - x
	if target == 0 {
		return len(nums)
	}
	l := 0
	window := 0
	ans := -1
	for r, v := range nums {
		window += v
		for l < r && window > target {
			window -= nums[l]
			l++
		}
		if window == target {
			newAns := len(nums) - (r - l + 1)
			if ans == -1 || ans > newAns {
				ans = newAns
			}
		}
	}
	return ans
}

func sum(nums []int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}
