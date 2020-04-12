package subtractTheProductAndSum

import "strconv"

/* My answer
func digit(i int, list []int) []int {
	if i > 0 {
		return digit(i/10, append(list, i%10))
	}
	return list
}

func SubtractProductAndSum(n int) int {
	list := make([]int, 0)
	digits := digit(n, list)
	fmt.Println(digits)
	var product int
	var sum int
	for i , num := range digits {
		if i == 0 {
			product, sum = num, num
		} else {
			product *= num
			sum += num
		}
	}
	return product - sum
}
*/

// More efficient code
func SubtractProductAndSum(n int) int {
	var product = 1
	var sum int
	
	strN := strconv.Itoa(n)
	for _, v := range strN {
		num, _ := strconv.Atoi(string(v))
		product *= num
		sum += num
	}
	
	return product - sum
}
