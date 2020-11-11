package _389_CreateTargetArrayIntheGivenOrder

func CreateTargetArray(nums []int, index []int) []int {
  target := make([]int, 0)
  target = append(target, nums[0])

  for i :=  1; i < len(index); i++ {
    if index[i] > len(target) - 1 {
      target = append(target, nums[i])
    } else {
      target = append(target[:index[i]], append([]int{nums[i]},target[index[i]:]...)...)
      target[index[i]] = nums[i]
    }
  }

  return target
}

