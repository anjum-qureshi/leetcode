/*
https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
difficulty: easy

PROBLEM
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

EXAMPLE
Input: head = [1,1,2]
Output: [1,2]

SOLUTION
* iterate through linked list curr.Next != nil
* if curr.Next.Val is curr.Val set curr.next to next.next
* else curr = curr.Next
*/

package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    curr := head
    for curr.Next != nil {
        if curr.Val == curr.Next.Val {
            curr.Next = curr.Next.Next
        } else {
            curr = curr.Next
        }
    }
    return head
}

