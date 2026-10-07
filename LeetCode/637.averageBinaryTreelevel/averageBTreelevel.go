//stats
// Runtime: 0ms - Beats 100.00%
// Memory: 8.06MB - Beats 93.79%

package main

import (
	"fmt"
)

func main() {
	fmt.Println(averageBTreeLevel(makeBinaryTree([]any{3, 9, 20, nil, nil, 15, 7}))) // [3.00000,14.50000,11.00000]
	fmt.Println(averageBTreeLevel(makeBinaryTree([]any{3, 9, 20, 15, 7})))           // [3.00000,14.50000,11.00000]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (N *TreeNode) String() string {
	treeList := []any{}
	q := []*TreeNode{N}
	for len(q) > 0 {
		currentNode := q[0]
		if currentNode != nil {
			treeList = append(treeList, currentNode.Val)
			q = append(q, currentNode.Left)
			q = append(q, currentNode.Right)
		} else {
			treeList = append(treeList, nil)
		}
		q = q[1:]

	}

	lastNum := len(treeList) - 1
	for lastNum > 0 && treeList[lastNum-1] == nil {
		lastNum--
	}

	return fmt.Sprint(treeList[:lastNum])
}

func makeBinaryTree(lst []any) *TreeNode {
	i := 0
	root := &TreeNode{
		Val: lst[i].(int),
	}
	treeList := map[int]*TreeNode{i: root}

	for i < len(lst) {

		leftChildIndex := 2*i + 1
		if leftChildIndex < len(lst) && lst[leftChildIndex] != nil {
			treeList[i].Left = &TreeNode{
				Val: lst[leftChildIndex].(int),
			}
			treeList[leftChildIndex] = treeList[i].Left
		}

		rightChildIndex := 2*i + 2
		if rightChildIndex < len(lst) && lst[rightChildIndex] != nil {
			treeList[i].Right = &TreeNode{
				Val: lst[rightChildIndex].(int),
			}
			treeList[rightChildIndex] = treeList[i].Right
		}

		i++
	}

	return root
}

func averageBTreeLevel(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	levelAverages := []float64{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		nodesInLevel := len(q)
		levelSum := 0

		for range nodesInLevel {
			currentNode := q[0]
			q = q[1:]

			levelSum += currentNode.Val

			if currentNode.Left != nil {
				q = append(q, currentNode.Left)
			}
			if currentNode.Right != nil {
				q = append(q, currentNode.Right)
			}
		}

		levelAverages = append(levelAverages, float64(levelSum)/float64(nodesInLevel))

	}

	return levelAverages
}
