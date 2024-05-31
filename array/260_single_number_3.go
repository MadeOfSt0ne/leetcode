package array

/* Given an integer array nums, in which exactly two elements appear only once and all the other elements appear exactly
twice. Find the two elements that appear only once. You can return the answer in any order.
You must write an algorithm that runs in linear runtime complexity and uses only constant extra space. */
func singleNumber(nums []int) []int {
	m := make(map[int]int)
	res := []int{}
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
