/*
https://leetcode.com/problems/range-sum-query-immutable/description/
difficulty: easy

PROBLEM
* Given an integer x, return true if x is a palindrome, and false otherwise.

SOLUTION
* Reverse the given number
* Return true if reverse is equal to given input
*/

package numeric

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	n := x
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n = n / 10
	}
	return reversed == x
}
