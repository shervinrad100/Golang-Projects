// stats:
// Runtime: 0ms - Beats 100.00%
//Memory: 4.16MB - Beats 44.63%

package main

import (
	"fmt"
)

func main() {
	fmt.Println(sameTree(buildTree([]any{1, 2, 3}), buildTree([]any{1, 2, 3}))) // true
	fmt.Println(sameTree(buildTree([]any{1, 2}), buildTree([]any{1, nil, 2})))  // false
	fmt.Println(sameTree(buildTree([]any{1, 2, 1}), buildTree([]any{1, 1, 2}))) // false
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

func sameTree(root1, root2 *TreeNode) bool {

	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	q1 := []*TreeNode{root1}
	q2 := []*TreeNode{root2}

	for len(q1) > 0 && len(q2) > 0 {
		// since we are not pointing at the same memory location we have to track 2 queueus
		// what makes two nodes the same is if they have the same values and have children with the same values
		cur1 := q1[0]
		cur2 := q2[0]

		if cur1.Val != cur2.Val {
			return false
		}

		// if they don't have the same number of children break
		// they have the same number of children but in the wrong order, break
		if (cur1.Left != nil) == (cur2.Left != nil) {
			if cur1.Left != nil {
				if cur1.Left.Val != cur2.Left.Val {
					return false
				}
				q1 = append(q1, cur1.Left)
				q2 = append(q2, cur2.Left)
			}
		} else {
			return false
		}

		if (cur1.Right != nil) == (cur2.Right != nil) {
			if cur1.Right != nil {
				if cur1.Right.Val != cur2.Right.Val {
					return false
				}
				q1 = append(q1, cur1.Right)
				q2 = append(q2, cur2.Right)
			}
		} else {
			return false
		}

		q1 = q1[1:]
		q2 = q2[1:]
	}

	return true
}
