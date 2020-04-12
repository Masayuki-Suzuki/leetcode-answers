package main

import (
	"fmt"
	"leetcode/jewelsAndStones"
)

func main() {
	//fmt.Println( "Total Steps: ", numberOfSteps.NumberOfSteps(14))
	fmt.Printf("Jewels: %d \n", jewelsAndStones.NumJewelsIsInStones("aA", "aAAbbbb"))
	fmt.Printf("Jewels: %d \n", jewelsAndStones.NumJewelsIsInStones("z", "ZZ"))
}
