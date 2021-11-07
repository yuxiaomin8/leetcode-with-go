package leetcode

import (
	"sort"
)

func ThreeSumClose(nums []int, target int) int {
	sort.Ints(nums)
	n, sumCloset := len(nums), nums[0]+nums[1]+nums[2]
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			if abs(sum-target) < abs(sumCloset-target) {
				sumCloset = sum
			}
			if sum > target {
				k--
			} else {
				j++

			}
		}

	}
	return sumCloset
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
