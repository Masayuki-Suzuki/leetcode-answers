package main

import (
	"fmt"
	"leetcode/SearchInsertPosition"
)

func main() {
	// Number of Steps
	//fmt.Println( "Total Steps: ", numberOfSteps.NumberOfSteps(14))
	
	// Jewels and Stones
	//fmt.Printf("Jewels: %d \n", jewelsAndStones.NumJewelsIsInStones("aA", "aAAbbbb"))
	//fmt.Printf("Jewels: %d \n", jewelsAndStones.NumJewelsIsInStones("z", "ZZ"))
	
	// Subtract the Product and Sum of Digits of an Integer
	//fmt.Println(subtractTheProductAndSum.SubtractProductAndSum(234))
	//fmt.Println(subtractTheProductAndSum.SubtractProductAndSum(4421))
	
	// 1313. Decompress Run-Length Encoded List
	//nums := []int{1,2,3,4}
	//fmt.Println(DecompressRun_LengthEncodedList.DecompressRLElist(nums))
	//nums = []int{1,1,2,3}
	//fmt.Println(DecompressRun_LengthEncodedList.DecompressRLElist(nums))
	//nums = []int{4,1,3,2}
	//fmt.Println(DecompressRun_LengthEncodedList.DecompressRLElist(nums))
	//nums = []int{4,3,2,1}
	//fmt.Println(DecompressRun_LengthEncodedList.DecompressRLElist(nums))
	
	// 35. Search Insert Position
	nums := []int{1,3,5,6}
	fmt.Println(SearchInsertPosition.SearchInsert(nums, 5))
	fmt.Println(SearchInsertPosition.SearchInsert(nums, 2))
	fmt.Println(SearchInsertPosition.SearchInsert(nums, 7))
	fmt.Println(SearchInsertPosition.SearchInsert(nums, 0))
	nums = []int{1}
	fmt.Println(SearchInsertPosition.SearchInsert(nums, 1))
	nums = []int{1,2}
	fmt.Println(SearchInsertPosition.SearchInsert(nums, 3))
	nums = []int{1,3,5}
	fmt.Println(SearchInsertPosition.SearchInsert(nums, 1))
}
