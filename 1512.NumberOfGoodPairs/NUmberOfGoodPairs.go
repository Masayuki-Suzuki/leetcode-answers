package _512_NumberOfGoodPairs

func NumIdenticalPairs(nums []int) int {
  pairs := 0
  length := len(nums)

  for i := 0; i < length; i++ {
    for j := 0; j < length; j++{
      if nums[i] == nums[j] && i < j {
        pairs++
      }
    }
  }

  return pairs
}
