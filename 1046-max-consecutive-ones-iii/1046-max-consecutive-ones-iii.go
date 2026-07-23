func longestOnes(nums []int, k int) int {
  l := 0
  r := 0
  max := 0
  zeros := 0
  for r < len(nums){
   if nums[r] == 0 {
    zeros++
   }
    for zeros > k {
        l++
        if nums[l-1] == 0{
        zeros --
        }
    }
    r++
    if r-l > max{
        max = r-l
    }

}
return max
}