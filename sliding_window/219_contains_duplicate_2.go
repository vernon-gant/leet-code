package sliding_window

func containsNearbyDuplicate(nums []int, k int) bool {
    windowFreq := make(map[int]bool)
    for i := 0; i < len(nums) && i <= k; i++ {
        inWindow, found := windowFreq[nums[i]]
        if found && inWindow {
            return true
        }
        windowFreq[nums[i]] = true
    }
    for i := 1; i + k < len(nums); i++ {
        windowFreq[nums[i - 1]] = false
        inWindow, found := windowFreq[nums[i + k]]
        if found && inWindow {
            return true
        }
        windowFreq[nums[i + k]] = true
    }
    return false
}