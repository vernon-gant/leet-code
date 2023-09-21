package two_pointers

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closestSum := math.MaxInt32
	for i := 0; i < len(nums) - 2; i++ {
		start, end := i + 1, len(nums) - 1
		for start < end {
			currentSum := nums[i] + nums[start] + nums[end]
			closestSum = newClosetSum(closestSum,currentSum,target)
			if currentSum > target {
				end--
			} else if currentSum < target {
				start++		
			} else {
				return target
			}
		}
	}
	return closestSum
}

func newClosetSum(closestSum, currentSum, target int) int {
	closestSumTargetDiff := absDiff(closestSum, target)
	currentSumTargetDiff := absDiff(currentSum, target)
	if closestSumTargetDiff < currentSumTargetDiff {
		return closestSum
	}
	return currentSum
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}