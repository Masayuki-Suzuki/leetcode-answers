package numberOfSteps

import "fmt"

func NumberOfSteps(num int) int {
	var steps int
	
	for num > 0 {
		fmt.Printf("Num: %d\n", num)
		fmt.Printf("Steps: %d\n", steps)
		if num % 2 == 0 {
			num = num / 2
		} else {
			num--
		}
		steps++
	}
	return steps
}
