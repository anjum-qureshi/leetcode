/*
https://leetcode.com/problems/remove-linked-list-elements/
difficulty: easy

PROBLEM
Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

EXAMPLE
Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]

SOLUTION
* keep a reduced head as dummy node and next as given head
* iterate through linked list curr.Next != nil
* if curr.Next.Val is val set curr.next to next.next
* else curr = curr.Next
*/

package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	reduced := &ListNode{Next: head}
	curr := reduced
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return reduced.Next
}

