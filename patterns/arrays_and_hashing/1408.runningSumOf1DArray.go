/*
https://leetcode.com/problems/running-sum-of-1d-array/description
difficulty: easy

PROBLEM
* Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
* Return the running sum of nums.

SOLUTION
* Implement prefix sum query where you save the sum of array till each index and the currSum = currEle+prevIdxSum
* complexity time: O(n), space O(n)
*/

package arrays_and_hashing

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
