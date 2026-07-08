func searchInsert(nums []int, target int) int {
    k := 1
  for i := 0; i < len(nums); i++{
    if nums[k-1] >= target{
        return k-1
    }
    k ++
  } 
  return len(nums) 
}