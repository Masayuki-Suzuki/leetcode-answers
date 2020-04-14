package DecompressRun_LengthEncodedList

func DecompressRLElist(nums []int) []int {
	encList := make([]int, 0)
	
	for i, val := range nums {
		if i % 2 == 0 {
			temp := make([]int, val)
			for n := 0; n < val; n++ {
				temp[n] = nums[i + 1]
			}
			encList = append(encList, temp...)
		}
	}
	
	return encList
}
