package main

import (
	"algos/merge_intervals"
	"fmt"
)

func main() {
	scheduleData := [][][]int{
		{{1, 3}, {6, 7}},
		{{2, 4}},
		{{2, 5}, {9, 12}},
	}

	schedule := make([][]*merge_intervals.Interval, len(scheduleData))
	for i, empSchedule := range scheduleData {
		schedule[i] = make([]*merge_intervals.Interval, len(empSchedule))
		for j, iv := range empSchedule {
			interval := &merge_intervals.Interval{}
			interval.IntervalInit(iv[0], iv[1])
			schedule[i][j] = interval
		}
	}

	freeTimes := merge_intervals.EmployeeFreeTime(schedule)

	for _, iv := range freeTimes {
		fmt.Println(iv.Str())
	}
}
