package main

import (
	"algos/merge_intervals"
	"fmt"
)

func main() {
	result := merge_intervals.IntervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}})
	fmt.Print(result)
}
