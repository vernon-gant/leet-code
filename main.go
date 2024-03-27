package main

import (
	"algos/merge_intervals"
	"fmt"
)

func main() {
	result := merge_intervals.Insert([][]int{{1, 3},{6,9}}, []int{2, 5})
	fmt.Print(result)
}
