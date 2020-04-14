package DecompressRun_LengthEncodedList

func DecompressRLElist(nums []int) []int {
	encList := make([]int, 0)
	for i, val := range nums {
		if i % 2 == 0 {
			for n := 0; n < val; n++ {
				encList = append(encList, nums[i + 1])
			}
		}
	}
	
	return encList
}
