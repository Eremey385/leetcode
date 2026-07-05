func removeElement(nums []int, val int) int {
    z := len(nums)
    i := 0
    for i < z {
        if nums[i] == val {
            nums[i] = nums[z-1]
            z--
        } else {
            i++
        }
    }
    return z
}