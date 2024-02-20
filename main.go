package main

import (
	"algos/sliding_window"
	"fmt"
)

func main() {
	replaced := sliding_window.MinSubArrayLen(4,[]int{1,4,4})
	fmt.Print(replaced)
}
