package modified_binary_search

import (
    "math"
    "sort"
)

func FindClosestElements(arr []int, k int, target int) []int {
    if target >= arr[len(arr) - 1] {
		return arr[len(arr) - k :]
	}
	if target <= arr[0] {
		return arr[0 : k]
	}
	start, end := 0, len(arr) - 1
	var closestElementIdx int
	for mid := (start + end) / 2; start <= end; mid = (start + end) / 2 {
		if arr[mid] == target {
			closestElementIdx = mid
			break
		} else if target > arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	if !(start <= end) {
		closestElementIdx = findClosest(arr, start, end, target)
	}
	left, right, result := closestElementIdx - 1, closestElementIdx + 1, make([]int, k)
	result[0] = arr[closestElementIdx]
	for i := 1; i < k; i++ {
		if right >= len(arr) || (left >= 0 && findClosest(arr, left, right, target) == left) {
			result[i] = arr[left]
			left--
		} else {
			result[i] = arr[right]
			right++
		}
	}
	sort.Slice(result, func(i, j int) bool {
        return result[i] < result[j]
    })
	return result
}

func findClosest(arr []int, start, end, target int) int {
	startDifference := math.Abs(float64(arr[start] - target))
	endDifference := math.Abs(float64(arr[end] - target))
	if startDifference < endDifference || (startDifference == endDifference && start < end) {
		return start
	}
	return end
}