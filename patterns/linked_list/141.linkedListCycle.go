/*
https://leetcode.com/problems/linked-list-cycle/description/
difficulty: easy

PROBLEM
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

SOLUTION 1 with hash set
* Maintain a map with node and visited as bool
* traverse through linked list if curr node is already then return true
* save the curr node to visited map
* else continue the loop and when curr is nil return false
Time complexity: O(n), Space complexity: O(n)

SOLUTION 2
* add a slow and fast both point to head
* traverse through linked list, slow = slow.Next and fast = fast.Next
* save the curr node to visited map
* else continue the loop and when curr is nil return false
*/

package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(curr *ListNode) bool {
	fast, slow := curr, curr
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func hasCycleWithMap(curr *ListNode) bool {
	visited := map[*ListNode]bool{}
	for curr != nil {
		if visited[curr] {
			return true
		}
		visited[curr] = true
		curr = curr.Next
	}
	return false
}
