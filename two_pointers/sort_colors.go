package two_pointers

func SortColors(nums []int) {
    start, end := 0, len(nums)-1
    for start < end {
        if nums[start] == 2 {
            nums[start], nums[end] = nums[end], nums[start]
            end--
        }
        if nums[end] == 0 {
            nums[start], nums[end] = nums[end], nums[start]
            start++
        }
        if nums[start] == 0 {
            start++
        }
        if nums[end] == 2 {
            end--
        }
    }
}
