/*
https://leetcode.com/problems/middle-of-the-linked-list/description/
difficulty: easy

PROBLEM
Given the head of a singly linked list, return the middle node of the linked list.
If there are two middle nodes, return the second middle node.

EXAMPLE
Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.

SOLUTION
* use fast and slow pointer
* traverse until fast.Next is nil, fast = fast.Next.Next,slow = slow.Next
* return slow pointer node
*/

package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
    fast,slow := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}
