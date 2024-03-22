package main

import (
	"algos/sliding_window"
	"fmt"
)

func main() {
<<<<<<< Updated upstream
<<<<<<< HEAD
	result := sliding_window.FindRepeatedDnaSequences("AAAAAAAAAAA")
=======
	result := sliding_window.MaxSlidingWindow([]int{7,2,4},2)
>>>>>>> Stashed changes
	fmt.Print(result)
=======
	repeated := sliding_window.FindRepeatedSequences("AGAGCTCCAGAGCTTG",6)
	fmt.Print(repeated)
>>>>>>> 4736469c001b27782897c622439e36268ec2ae53
}
