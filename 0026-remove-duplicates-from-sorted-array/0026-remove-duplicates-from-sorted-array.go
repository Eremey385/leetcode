func removeDuplicates(nums []int) int {
    slow := 0
    k := 1
    for fast := 1; fast < len(nums); fast++ {
      if  (nums[slow] < nums[fast]){
        slow ++
        nums[slow],nums [fast] = nums[fast],nums[slow]
        k ++ 
      }
    }
    return k
}