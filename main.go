package main

import (
    "algos/subsets"
)

func main() {
    result := subsets.CanCompleteCircuit([]int{5,8,2,8}, []int{6,5,6,6})
    print(result)
}
