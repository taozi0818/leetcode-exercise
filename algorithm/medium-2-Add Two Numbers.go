//You are given two non-empty linked lists representing two non-negative
//integers. The digits are stored in reverse order, and each of their nodes contains a
//single digit. Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the
//number 0 itself.
//
//
// Example 1:
//
//
//Input: l1 = [2,4,3], l2 = [5,6,4]
//Output: [7,0,8]
//Explanation: 342 + 465 = 807.
//
//
// Example 2:
//
//
//Input: l1 = [0], l2 = [0]
//Output: [0]
//
//
// Example 3:
//
//
//Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//Output: [8,9,9,9,0,0,0,1]
//
//
//
// Constraints:
//
//
// The number of nodes in each linked list is in the range [1, 100].
// 0 <= Node.val <= 9
// It is guaranteed that the list represents a number that does not have
//leading zeros.
//
// Related Topics 递归 链表 数学 👍 8166 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	r := addTwoNumbers(l1, l2)

	for r != nil {
		fmt.Printf("%d -> ", r.Val)
		r = r.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return _addTwoNumbers(l1, l2, 0)
}

func _addTwoNumbers(l1 *ListNode, l2 *ListNode, i int) *ListNode {
	if l1 == nil && l2 == nil && i == 0 {
		return nil
	}

	if l1 == nil {
		l1 = &ListNode{
			Val: 0,
		}
	}

	if l2 == nil {
		l2 = &ListNode{
			Val: 0,
		}
	}

	v := l1.Val + l2.Val + i

	if v >= 10 {
		v = v - 10
		i = 1
	} else {
		i = 0
	}

	return &ListNode{
		Val:  v,
		Next: _addTwoNumbers(l1.Next, l2.Next, i),
	}
}

//leetcode submit region end(Prohibit modification and deletion)
