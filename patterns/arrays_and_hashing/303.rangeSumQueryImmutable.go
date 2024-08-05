/*
https://leetcode.com/problems/range-sum-query-immutable/description/
difficulty: easy

PROBLEM
* Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.

EXAMPLE
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
Prefix Sum = [-2, -2, 1, -4, -2, -3]
numArray.sumRange(0, 2); return (-2) + 0 + 3 = 1
numArray.sumRange(2, 5); return 3 + (-5) + 2 + (-1) = -1
numArray.sumRange(0, 5); return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3

SOLUTION
* Calculate prefix sum at each index and store in the NumArray
* When left is 0 return the sum at right index
* Else return right sum - left sum
* complexity time: O(n+q), space O(n)
*/

package arraysandhashing

type NumArray struct {
	elements []int
}

func Constructor(nums []int) NumArray {
	return NumArray{elements: runningSum(nums)}
}

func runningSum(array []int) []int {
	result := []int{}
	if len(array) == 0 {
		return result
	}
	result = append(result, array[0])
	for idx := 1; idx < len(array); idx++ {
		result = append(result, result[idx-1]+array[idx])
	}
	return result
}

func (na *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return na.elements[right]
	}
	return na.elements[right] - na.elements[left-1]
}
