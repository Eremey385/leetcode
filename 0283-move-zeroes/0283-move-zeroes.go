func moveZeroes(nums []int)  {
  fast := 0
  slow := 0
 
 for fast < len(nums){
    if nums[fast] != 0{
        nums[fast],nums[slow]=nums[slow],nums[fast]
        fast++
        slow++
    }else {
        fast++
        }
 }
}