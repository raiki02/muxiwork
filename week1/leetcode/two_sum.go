func twoSum(nums []int, target int) []int {
  arr := [2]int{}
  for i := 0; i < len(nums); i++ {
    for j := i + 1; j < len(nums); j++ {
       if nums[i]+nums[j] == target {
         arr[0] = i
         arr[1] = j
         break
    }
   }
  }
  return arr[:]
}
