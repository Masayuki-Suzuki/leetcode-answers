package _470_ShuffleTheArray

func Shuffle(nums []int, n int) []int {
  newArray := make([]int, 0)

  for i, p := 0,0; i < len(nums) && p < n; i, p = i + 2, p + 1 {
    newArray = append(newArray, nums[p])
    newArray = append(newArray, nums[n + p])
  }

  return newArray
}
