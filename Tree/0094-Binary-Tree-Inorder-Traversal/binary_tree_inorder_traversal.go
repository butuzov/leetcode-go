package main

import "github.com/butuzov/leetcode.go/pkg/binarytree"

/*******************************************************************************
  Problem Solution
*******************************************************************************/

func inorderTraversalRecursive(root *binarytree.TreeNode) []int {
	var results = []int{}
	if root == nil {
		return results
	}

	var inOrder func(*binarytree.TreeNode)
	inOrder = func(t *binarytree.TreeNode) {
		if t == nil {
			return
		}

		inOrder(t.Left)
		results = append(results, t.Val)
		inOrder(t.Right)
	}

	inOrder(root)
	return results
}

func inorderTraversal(root *binarytree.TreeNode) []int {
	var results = []int{}
	var stack = []*binarytree.TreeNode{}
	var node = root

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		if len(stack) > 0 {
			node = stack[len(stack)-1]
			results = append(results, node.Val)
			stack = stack[:len(stack)-1]
			node = node.Right
		}

	}

	return results
}
