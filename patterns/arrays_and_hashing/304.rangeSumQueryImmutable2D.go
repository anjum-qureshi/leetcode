/*
https://leetcode.com/problems/range-sum-query-2d-immutable/description/
difficulty: medium

PROBLEM
Given a 2D matrix matrix, handle multiple queries of the following type:

Calculate the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).
Implement the NumMatrix class:

NumMatrix(int[][] matrix) Initializes the object with the integer matrix matrix.
int sumRegion(int row1, int col1, int row2, int col2) Returns the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2)..

EXAMPLE
NumMatrix numArray = new NumMatrix([-2, 0, 3, -5, 2, -1]);
Prefix Sum = [-2, -2, 1, -4, -2, -3]
numArray.sumRange(0, 2); return (-2) + 0 + 3 = 1
numArray.sumRange(2, 5); return 3 + (-5) + 2 + (-1) = -1
numArray.sumRange(0, 5); return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3

SOLUTION
* Calculate prefix sum at each index and store in the NumMatrix at each index with following
* p[r-1][c] + p[r-1][c] + nums[r-1][c-1] - p[r-1][c-1]
* Calculate the sumRegin with below formula
* bottomRight,above := p[r2][c2], p[r1-1][c2]
* left,topLeft := p[r2][c1-1], p[r1-1][c1-1]
* complexity time: O(m*n+q), space O(m*n)
*/

package arraysandhashing

type NumMatrix struct {
	elements [][]int
}

func Constructor(nums [][]int) NumMatrix {
	p := [][]int{make([]int, len(nums[0])+1)}
	for r := 1; r < len(nums)+1; r++ {
        p = append(p, []int{0})
		for c := 1; c < len(nums[0])+1; c++ {
			ele := p[r-1][c] + p[r][c-1] + nums[r-1][c-1] - p[r-1][c-1]
			p[r] = append(p[r],ele)
		}
	}
	return NumMatrix{elements: p}
}

func (na *NumMatrix) SumRegion(r1 int, c1 int, r2 int, c2 int) int {
	r1, c1, r2, c2 = r1+1, c1+1, r2+1, c2+1
	return na.elements[r2+1][c2+1] - na.elements[r1-1][c2] - na.elements[r2][c1-1] + na.elements[r1-1][c1-1]
}

