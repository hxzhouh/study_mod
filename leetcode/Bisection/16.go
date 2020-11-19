package Bisection

import (
	"math"
	"sort"
)

/**
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
示例：
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
提示：
3 <= nums.length <= 10^3
-10^3 <= nums[i] <= 10^3
-10^4 <= target <= 10^4
*/
// 做题要仔细呀
func threeSumClosest(nums []int, target int) int {
	// 对nums 排序
	sort.Ints(nums)
	var (
		n    = len(nums)
		best = math.MaxInt32
	)
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}
	for i := 0; i < n; i++ { //枚举 A
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target { //相等就直接返回。
				return target
			}
			update(sum)
			if sum > target {
				k -= 1
			} else {
				j += 1
			}
		}
	}
	return best
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
