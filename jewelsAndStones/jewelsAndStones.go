package jewelsAndStones

func NumJewelsIsInStones(J, S string) int {
	var count int
	
	for _, chr := range J {
		for _, s := range S {
			if chr == s {
				count++
			}
		}
	}
	return count
}
