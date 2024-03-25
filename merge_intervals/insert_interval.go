package merge_intervals

func Insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return make([][]int, 0)
	}

	result, currentIdx := make([][]int, 0), 0
	for ; !isOverlapping(intervals[currentIdx], newInterval); currentIdx++ {
		result = append(result, intervals[currentIdx])
	}
	for ; isOverlapping(intervals[currentIdx], newInterval); currentIdx++ {
		
	}
	return result
}

func isOverlapping(existing, newInterval []int) bool {
	return existing[1] >= newInterval[0] || newInterval[1] <= existing[0]
}
