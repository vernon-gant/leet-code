package merge_intervals

import "sort"

func findSets(intervals [][]int) int{
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	meetingEndTimes := []int{intervals[0][1] }
	for _, interval := range intervals[1:] {
		roomFound := false
		for i := 0; i < len(meetingEndTimes); i++ {
			if interval[0] >= meetingEndTimes[i] {
				meetingEndTimes[i] = interval[0]
				roomFound = true
				break
			}
		}
		if roomFound {
			continue
		}
		meetingEndTimes = append(meetingEndTimes, interval[1])
	}
	return len(meetingEndTimes)
}