/* 
https://leetcode.com/problems/find-pivot-index/description/
difficulty: easy

PROBLEM
* Given an array of integers nums, calculate the pivot index of this array.
* The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.
* If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.
* Return the leftmost pivot index. If no such index exists, return -1.

EXAMPLE
Input: nums = [1,7,3,6,5,6]
Output: 3
Explanation:
The pivot index is 3.
Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
Right sum = nums[4] + nums[5] = 5 + 6 = 11

SOLUTION1
* Calculate prefix sum sum at each index
* Check if prefix and postfix are same return index, else return -1 is not exist
* complexity time: O(2n), space O(n)

SOLUTION2
* Calculate total totalSum of the array
* Now iterate through array, leftSum += curr, rightSum = totalSum - curr - leftSum
* complexity time: O(2n), space O(1)
*/


package arraysandhashing


func pivotIndex(nums []int) int {
    n := len(nums)
    prefixSums := []int{}
	if n == 0 {
		return -1
	}
	prefixSums = append(prefixSums, nums[0])
	for idx := 1; idx < n; idx++ {
		prefixSums = append(prefixSums, prefixSums[idx-1]+nums[idx])
	}
    if prefixSums[n-1]==prefixSums[0] {
        return 0
    }
    for i := 1; i < n; i++ {
        if prefixSums[i-1] == prefixSums[n-1] - prefixSums[i] {
            return i
        }
    }
    return -1
}

func pivotIndexWithoutPrefixSum(nums []int) int {
    sum := 0
	for idx := 0; idx < len(nums); idx++ {
        sum += nums[idx]
	}
    left := 0
    for i := 0; i < len(nums); i++ {
       if left == sum - left - nums[i] {
        return i
       }
       left += nums[i]
    }
    return -1
}
