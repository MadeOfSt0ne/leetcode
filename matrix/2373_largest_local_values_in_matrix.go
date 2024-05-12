package matrix

/* You are given an n x n integer matrix grid.
Generate an integer matrix maxLocal of size (n - 2) x (n - 2) such that:
maxLocal[i][j] is equal to the largest value of the 3 x 3 matrix in grid centered around row i + 1 and column j + 1.
In other words, we want to find the largest value in every contiguous 3 x 3 matrix in grid.
Return the generated matrix. */

func largestLocal(grid [][]int) [][]int {
	res := make([][]int, 0)
	for row := range len(grid) - 2 {
		ansRow := make([]int, 0)
		for column := range len(grid[0]) - 2 {
			m := grid[row][column]
			for i := range 3 {
				for j := range 3 {
					m = max(m, grid[row+i][column+j])
				}
			}
			ansRow = append(ansRow, m)
		}
		res = append(res, ansRow)
	}
	return res
}
