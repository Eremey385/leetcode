func sortArrayByParity(nums []int) []int {
 i := 0
 j := len(nums)-1 
 for i < j {
    if nums[i] % 2 == 1 {
        nums[i],nums[j] = nums[j],nums[i]
        j--
    }else {
        i++
    }
 }  
 return nums
}