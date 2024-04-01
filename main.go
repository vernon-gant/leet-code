package main

import (
	"algos/merge_intervals"
	"fmt"
)

func main() {
	leastInterval := merge_intervals.LeastInterval([]byte{'A','C','A','B','D','B'},1)
	fmt.Print(leastInterval)
}
