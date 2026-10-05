// stats:
// Runtime: 0ms - Beats 100.00%
// Memory: 7.80MB - Beats 99.42%

package main

import (
	"fmt"
)

func main() {
	fmt.Println(treeDepth(makeBinaryTree([]any{3, 9, 20, nil, nil, 15, 7}))) // 3
	fmt.Println(treeDepth(makeBinaryTree([]any{1, nil, 2})))                 // 2
	fmt.Println(treeDepth(makeBinaryTree([]any{})))                          // 0
	fmt.Println(treeDepth(makeBinaryTree([]any{0})))                         // 1
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
	queue := []*TreeNode{N}
	tree := []any{}

	for len(queue) > 0 {
		currentNode := queue[0]

		if currentNode != nil {
			tree = append(tree, currentNode.Val)
			queue = append(queue, currentNode.Left)
			queue = append(queue, currentNode.Right)

		} else {
			tree = append(tree, nil)
		}

		queue = queue[1:]
	}

	for len(tree) > 0 && tree[len(tree)-1] == nil {
		tree = tree[:len(tree)-1]
	}

	return fmt.Sprint(tree)
}

func makeBinaryTree(lst []any) *TreeNode {
	if len(lst) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: lst[0].(int),
	}
	queue := map[int]*TreeNode{0: root}
	i := 0

	for len(queue) > 0 {
		currentNode := queue[i]

		// add left child
		left := 2*i + 1
		if left < len(lst) && lst[left] != nil {
			currentNode.Left = &TreeNode{
				Val: lst[left].(int),
			}
			queue[left] = currentNode.Left
		}

		right := 2*i + 2
		if right < len(lst) && lst[right] != nil {
			currentNode.Right = &TreeNode{
				Val: lst[right].(int),
			}
			queue[right] = currentNode.Right
		}

		delete(queue, i)
		i++

	}

	return root
}

type nodeDepth struct {
	node  *TreeNode
	depth int
}

func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []nodeDepth{{root, 1}}
	maxDepth := 1

	for len(q) > 0 {
		current := q[0]
		if current.node.Left != nil {
			q = append(q, nodeDepth{current.node.Left, current.depth + 1})
			maxDepth = current.depth + 1
		}
		if current.node.Right != nil {
			q = append(q, nodeDepth{current.node.Right, current.depth + 1})
			maxDepth = current.depth + 1
		}
		q = q[1:]

	}
	return maxDepth
}
