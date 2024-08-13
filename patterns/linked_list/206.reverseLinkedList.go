/*
https://leetcode.com/problems/reverse-linked-list/description/
difficulty: easy

PROBLEM
Given the head of a singly linked list, reverse the list, and return the reversed list.

EXAMPLE
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

SOLUTION
* keep prev = null, curr
* keep curr.Next in temp to rewire later
*/

package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(curr *ListNode) *ListNode {
    var prev,next *ListNode
    for curr != nil {
        next = curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}
