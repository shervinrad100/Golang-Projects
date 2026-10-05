// stats:
// Runtime: 0ms - Beats 100.00%
// Memory: 6.81MB - Beats 6.92%

package main

import (
	"fmt"
)

func main() {
	fmt.Println(pathSum(buildTree([]any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}), 22)) // true
	fmt.Println(pathSum(buildTree([]any{1, 2, 3}), 5))                                          // false
	fmt.Println(pathSum(buildTree([]any{}), 0))                                                 // false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (N *TreeNode) String() string {
	if N == nil {
		return "[]"
	}

	tree := []any{}
	q := []*TreeNode{N}

	for len(q) > 0 {
		current := q[0]

		if current != nil {
			tree = append(tree, current.Val)
			q = append(q, current.Left)
			q = append(q, current.Right)
		} else {
			tree = append(tree, nil)
		}

		q = q[1:]
	}

	for len(tree) > 0 && tree[len(tree)-1] == nil {
		tree = tree[:len(tree)-1]
	}

	return fmt.Sprint(tree)
}

func buildTree(lst []any) *TreeNode {
	if len(lst) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: lst[0].(int),
	}

	q := map[int]*TreeNode{0: root}
	i := 0

	for len(q) > 0 {
		current := q[i]

		leftChildIndex := 2*i + 1
		if leftChildIndex < len(lst) {
			if lst[leftChildIndex] != nil {
				current.Left = &TreeNode{
					Val: lst[leftChildIndex].(int),
				}
				q[leftChildIndex] = current.Left
			}
		}

		rightChildIndex := 2*i + 2
		if rightChildIndex < len(lst) {
			if lst[rightChildIndex] != nil {
				current.Right = &TreeNode{
					Val: lst[rightChildIndex].(int),
				}
				q[rightChildIndex] = current.Right
			}
		}

		delete(q, i)
		i++
	}

	return root
}

type nodeRunningSum struct {
	node *TreeNode
	sum  int
}

func pathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := []nodeRunningSum{{root, root.Val}}

	// DFS
	for len(stack) > 0 {
		currentNodeSum := stack[len(stack)-1]
		currentNode := currentNodeSum.node
		stack = stack[:len(stack)-1]

		if currentNode.Right != nil {
			stack = append(stack, nodeRunningSum{currentNode.Right, currentNodeSum.sum + currentNode.Right.Val})
		}
		if currentNode.Left != nil {
			stack = append(stack, nodeRunningSum{currentNode.Left, currentNodeSum.sum + currentNode.Left.Val})
		}
		if currentNode.Right == nil && currentNode.Left == nil {
			if currentNodeSum.sum == targetSum {
				return true
			}
		}

	}

	return false
}
