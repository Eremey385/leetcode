func sortColors(nums []int) {
   search := 0
   start := 0
for search < 3{
j := start
count := 0
   for k := start; k < len(nums);k++{
      if  nums[k] == search {
            nums[j],nums[k] = nums[k],nums[j]
            count ++
                j++
            }
   }
    start += count
    search ++
   }
}

