/*
https://leetcode.com/problems/product-of-array-except-self/description/
difficulty: hard

PROBLEM
* Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation

EXAMPLE
Input: nums = [1,2,3,4]
Output: [24,12,8,6]


[]

SOLUTION
* Traverse array from start find prefix product without including the curr value
* Traverse from the end of the array and at each index
* val = postFix of index * prefix of index
* complexity time: O(n), space O(n)
*/

package arraysandhashing

func productExceptSelf(nums []int) []int {
	left := []int{1}
	curr := 1
	for i := 0; i < len(nums); i++ {
		left = append(left, curr*left[i])
		curr = left[i]
	}
	return []int{}
}



// [1,2,3,4]
// [1,2,6,24]
// [24,24,12,4]


// [1,1,2,6]
// [,,,1*6]


//  [24,12,8,6]