package _229

func majorityElement(nums []int) []int {
	// 先用"摩爾投票算法(Boyer-Moore Voting Algorithm)"算出出現次數最多的兩個數字
	// 再重新遍歷數組計算這兩個數字出現的次數
	k1, k2, v1, v2 := nums[0], nums[0], 0, 0
	for _, n := range nums {
		if n == k1 {
			v1++
			continue
		} else if n == k2 {
			v2++
			continue
		} else if v1 == 0 {
			k1 = n
			v1 = 1
		} else if v2 == 0 {
			k2 = n
			v2 = 1
		} else {
			v1--
			v2--
		}
	}
	v1, v2 = 0, 0
	for _, n := range nums {
		if n == k1 {
			v1++
		} else if n == k2 {
			v2++
		}
	}
	ans := make([]int, 0)
	if v1 > len(nums)/3 {
		ans = append(ans, k1)
	}
	if v2 > len(nums)/3 {
		ans = append(ans, k2)
	}
	return ans
}
