package fast_slow

func CircularArrayLoop(nums []int) bool {
	for i, _ := range nums {
		if nums[i] == 0 {
			continue
		}
		slowIdx, fastIdx := i, i
		for nextFastIdx := getNextIdx(nums, fastIdx); nums[fastIdx]*nums[nextFastIdx] > 0 && nums[fastIdx]*nums[getNextIdx(nums, nextFastIdx)] > 0 && fastIdx != nextFastIdx; nextFastIdx = getNextIdx(nums, fastIdx) {
			slowIdx = getNextIdx(nums, slowIdx)
			fastIdx = getNextIdx(nums, getNextIdx(nums, fastIdx))
			if slowIdx == fastIdx && fastIdx != getNextIdx(nums, fastIdx) {
				return true
			}
		}
	}
	return false
}

func getNextIdx(nums []int, idx int) int {
	return nonNegativeModulo(idx+nums[idx], len(nums))
}

func nonNegativeModulo(num, modulo int) int {
	return ((num % modulo) + modulo) % modulo
}
