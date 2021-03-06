package main

import "github.com/butuzov/leetcode.go/pkg/linkedlist"

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func addTwoNumbers(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode) *linkedlist.ListNode {
	var numbers []int
	var extra int
	var number int

	for {
		number = 0
		var done bool

		if l1 != nil {
			number += l1.Val
			l1 = l1.Next
			done = true
		}

		if l2 != nil {
			number += l2.Val
			l2 = l2.Next
			done = true
		}

		if extra > 0 {
			number += extra
			done = true
		}

		if done == false {
			break
		}

		numbers = append(numbers, number%10)

		if number >= 10 {
			extra = 1
		} else {
			extra = 0
		}
	}

	return linkedlist.MakeList(numbers...)
}
