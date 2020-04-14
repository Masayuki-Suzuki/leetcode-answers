package SearchInsertPosition

import "fmt"

func SearchInsert(nums []int, target int) int {
	fmt.Println("================")
	fmt.Printf("traget: %d, array: %v\n", target, nums)
	var result, left int
	right := nums[len(nums) - 1]
	
	if right < target {
		return len(nums)
	} else if nums[0] > target {
		return 0
	} else {
		for i := 0; i < len(nums); i++ {
			switch {
			case nums[i] == target:
				return i
			case nums[i] > left && nums[i] <= target:
				left = nums[i]
				result = i + 1
			case nums[i] < right && nums[i] >= target:
				right = nums[i]
				result = i
			}
		}
	}
	
	fmt.Printf("result: %d, left: %d, right: %d \n", result, left, right)
	
	return result
}
