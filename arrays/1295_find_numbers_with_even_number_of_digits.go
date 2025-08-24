package arrays

func findNumbers(nums []int) int {
	var counter int
	for _, val := range nums {
		currNumber := 1
		currVal := val
		for currVal/10 != 0 {
			currNumber++
			currVal = currVal / 10
		}
		if currNumber%2 == 0 {
			counter++
		}
	}
	return counter
}
