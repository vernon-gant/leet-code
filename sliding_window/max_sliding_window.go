package sliding_window

import (
	"math"
	"slices"
)

func MaxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	var maxWindow, currentWindow []int
	currentWindow = nums[0:k]
	slices.Sort(currentWindow)
	maxWindow = append(maxWindow,currentWindow[k - 1])
	for i := k; i < len(nums); i++ {
		currentWindow[k-1] = 0;
		
	}
	return maxWindow
}

func GetMaxFromSlice(nums []int) int {
	tempMax := math.MinInt
	for _, value := range nums {
		if value > tempMax {
			tempMax = value
		}
	}
	return tempMax
}

func maxNumber(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}