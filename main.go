package main

import (
	"algos/two_heaps"
	"fmt"
)

func main() {
	medianFinder := two_heaps.Constructor()
	medianFinder.AddNum(1);    // arr = [1]
	medianFinder.AddNum(2);    // arr = [1, 2]
	fmt.Println(medianFinder.FindMedian()) // return 1.5 (i.e., (1 + 2) / 2)
	medianFinder.AddNum(3);    // arr[1, 2, 3]
	fmt.Println(medianFinder.FindMedian())
}
