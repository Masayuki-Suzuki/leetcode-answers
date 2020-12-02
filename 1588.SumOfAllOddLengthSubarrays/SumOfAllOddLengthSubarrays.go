package _588_SumOfAllOddLengthSubarrays

// This is bad idea.
func SumOddLengthSubarrays(arr []int) int {
  result, length := 0, len(arr)

  for step := 1; step <= length; step += 2 {
    if step <= length {
      for i := 0; i < length; i++ {
        if i + step <= length {
          for p := i; p < i + step; p++ {
            result = result + arr[p]
          }
        }
      }
    }
  }

  return result
}


//func sumOddLengthSubarrays(arr []int) int {
//  ans := 0
//
//  for left, right := 1, len(arr); left - 1 < len(arr); left, right = left + 1, right - 1 {
//    ans += (left * right + 1) / 2 * arr[left - 1]
//  }
//
//  return ans
//}
