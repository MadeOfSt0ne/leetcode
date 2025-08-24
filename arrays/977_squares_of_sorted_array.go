package arrays

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left := 0
	right := len(nums) - 1
	index := len(nums) - 1
	for index >= 0 {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[index] = nums[left] * nums[left]
			left++
		} else {
			res[index] = nums[right] * nums[right]
			right--
		}
		index--
	}
	return res
}
