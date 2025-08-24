package arrays

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	counter := 0
	for _, val := range nums {
		if val == 0 {
			if counter > max {
				max = counter
			}
			counter = 0
		} else {
			counter++
		}
	}
	if counter > max {
		max = counter
	}
	return max
}
