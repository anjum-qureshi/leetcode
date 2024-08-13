/*
https://leetcode.com/problems/merge-two-sorted-lists/description/
difficulty: easy

PROBLEM
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.

EXAMPLE
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

SOLUTION
* keep a merged head with empty nil as insertAt
* take each node from list 1 and 2 , c1,c2
* if c1 < c2 then add c1 to the insertAt.Next
* else add c2 to the insertAt.Next
*/

package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(curr1 *ListNode, curr2 *ListNode) *ListNode {
	merged := &ListNode{}
	insertAt := merged

	for curr1 != nil && curr2 != nil {
		if curr1.Val < curr2.Val {
			insertAt.Next = curr1
			insertAt = insertAt.Next
			curr1 = curr1.Next
		} else {
			insertAt.Next = curr2
			insertAt = insertAt.Next
			curr2 = curr2.Next
		}
	}
	if curr1 != nil {
		insertAt.Next = curr1
	} else {
		insertAt.Next = curr2
	}
	return merged.Next
}
