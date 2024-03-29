package merge_intervals

import (
	"sort"
	"strconv"
)

type Interval struct {
	Start int
	End   int
}

func (i *Interval) IntervalInit(start int, end int) {
	i.Start = start
	i.End = end
}

func (i *Interval) Str() string {
	out := "(" + strconv.Itoa(i.Start) + ", " + strconv.Itoa(i.End) + ")"
	return out
}

func EmployeeFreeTime(schedule [][]*Interval) []*Interval {
	var allIntervals, result []*Interval
	for _, singleSchedule := range schedule {
		allIntervals = append(allIntervals, singleSchedule...)
	}
	sort.Slice(allIntervals, func(i, j int) bool {
		if allIntervals[i].Start == allIntervals[j].Start {
			return allIntervals[i].End < allIntervals[j].End
		}
		return allIntervals[i].Start < allIntervals[j].Start
	})
	for i := 0; i < len(allIntervals) - 1; i++ {
		if allIntervals[i+1].Start <= allIntervals[i].End {
			allIntervals[i + 1].Start = min(allIntervals[i].Start,allIntervals[i+1].Start)
			allIntervals[i + 1].End = max(allIntervals[i].End,allIntervals[i+1].End)
			continue
		}
		result = append(result, &Interval{
			Start: allIntervals[i].End,
			End:   allIntervals[i+1].Start,
		})
	}
	return result
}