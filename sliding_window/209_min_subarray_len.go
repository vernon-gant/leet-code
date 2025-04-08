package sliding_window

func MinSubArrayLen(target int, nums []int) int {
    left, right, tempSum := 0, 0, 0

    for ; right < len(nums) && tempSum < target; right++ {
        tempSum += nums[right]
    }

    if tempSum < target {
        return 0
    }

    if right == len(nums) {
        left = getNewLeft(&tempSum, left, target, nums)
    }

    result := right - left

    for ; right != len(nums); right++ {
        tempSum += nums[right]
        left = getNewLeft(&tempSum, left, target, nums)
        result = min(result, right-left+1)
    }

    return result
}

func getNewLeft(tempSum *int, left int, target int, nums []int) int {
    for *tempSum-nums[left] >= target {
        *tempSum -= nums[left]
        left++
    }

    return left
}
