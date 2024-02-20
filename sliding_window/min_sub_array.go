package sliding_window

import "math"

func MinSubArrayLen(target int, nums []int) int {
	start, end := 0, 0
	currentSum, minSum := 0, math.MaxInt
	minSubSize := math.MaxInt
	inBounds := func() bool {
		return end < len(nums)
	}
	increased := true
	for inBounds() {
		for inBounds() && currentSum < target {
			currentSum += nums[end]
			end++
			increased = true
		}
		if increased {end--}
		for currentSum - nums[start] >= target {
			currentSum -= nums[start]
			start++
		}
		minSubSize, minSum = min(end - start + 1, minSubSize), min(minSum, currentSum)
		end++
		if !inBounds() {break}
		currentSum += nums[end]
		increased = false
	}
	if minSum >= target {
		return minSubSize
	}
	return 0
}

/*
func minSubArrayLen(target int, nums []int) int {
	head, sum := 0, 0
	length := len(nums)+1

	for tail := 0; tail < len(nums); tail++ {
		sum += nums[tail]
		for sum >= target {
			if tail-head+1 < length {
				length = tail-head+1
			}
			sum -= nums[head]
			head++
		}
	}
	if length == len(nums)+1 {
		return 0
	}
	return length
}
*/