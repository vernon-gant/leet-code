package two_pointers

func RemoveDuplicates(nums []int) int {
	nextInsertIdx := 1
	currentIndex := 0
	for currentIndex < len(nums) {
		currentElement := nums[currentIndex]
		for currentIndex < len(nums) && currentElement == nums[currentIndex] {
			currentIndex++
		}
		if currentIndex == len(nums) {
			break
		}
		nums[nextInsertIdx] = nums[currentIndex]
		nextInsertIdx++
	}
	return nextInsertIdx
}
