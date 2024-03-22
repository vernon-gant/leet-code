package main

import (
	"algos/merge_intervals"
	"fmt"
)

func main() {
	result := merge_intervals.MergeIntervals([][]int{{1,4}, {0,0}})
	fmt.Print(result)
}
