package _365_HowManyNumbersAreSmallerThanTheCurrentNumber

func SmallerNumbersThanCurrent(nums []int) []int {
  length := len(nums)
  newArray := make([]int, length)
  count := 0


  for i := 0; i < length; i++ {
    count = 0

    for j := 0; j < length; j++ {
      if nums[i] != nums[j] && nums[i] > nums[j] {
        count++
      }
    }

    newArray[i] = count

  }

  return newArray
}
