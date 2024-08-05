/* https://leetcode.com/problems/count-submatrices-with-top-left-element-and-sum-less-than-k/description/
difficulty: medium

PROBLEM:
You are given a 0-indexed integer matrix grid and an integer k.
Return the number of submatrices that contain the top-left element of the grid, and have a sum less than or equal to k.

EXAMPLE
grid = [[7,6,3],[6,6,1]], k = 18, expected = 4

SOLUTION
* Calculate prefixSum for the given matrix
* Run through prefixSum matrix starting from row,0 to col (n-1)
* If element less than or equal to k increment count by col, row by 1
* Else decrement col by 1
*/

package arraysandhashing


func countSubmatrices(grid [][]int, k int) int {
    p := [][]int{make([]int, len(grid[0])+1)}
	for r := 1; r < len(grid)+1; r++ {
        p = append(p, []int{0})
		for c := 1; c < len(grid[0])+1; c++ {
			ele := p[r-1][c] + p[r][c-1] + grid[r-1][c-1] - p[r-1][c-1]
			p[r] = append(p[r],ele)
		}
	}


    count,row,col := 0, 1, len(p[0])-1
    for row < len(p) && col > 0 {
        sum := p[row][col] - p[0][col] - p[row][0] + p[0][0]
        if sum <= k {
            count += col
            row++
        } else {
            col--
        }
    }
    return count
}
