package backtrack

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)

	backtrack(&result, []int{}, candidates, target, 0)
	return result
}

func backtrack(result *[][]int, combination []int, candidates []int, target, startIdx int) {
	if target == 0 {
		validCombo := make([]int, len(combination))
		copy(validCombo, combination)
		*result = append(*result, validCombo)
		return
	} else if target < 0 {
		return
	} else {
		for i := startIdx; i < len(candidates); i++ {
			backtrack(result, append(combination, candidates[i]), candidates, target-candidates[i], i)
		}
	}
}
