package _486_XOROperationInAnArray

//func XorOperation(n int, start int) int {
//  nums := make([]int, 0)
//  sum := 0
//
//  for i := start; len(nums) < n; i += 2 {
//    nums = append(nums, i)
//    sum ^= i
//  }
//
//  return sum
//}

// Improved
func XorOperation(n int, start int) int {
  res := 0

  for i := 0; i < n; i++ {
    res ^= start + 2 * i
  }

  return res
}
