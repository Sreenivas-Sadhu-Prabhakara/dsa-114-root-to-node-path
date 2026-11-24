package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// PathToNode returns the values on the path from the root to the first node
// holding target (nil if absent). DFS with backtracking: the path grows on
// the way down and sheds dead branches on the way back.
func PathToNode(root *TreeNode, target int) []int {
	path := []int{}
	var walk func(n *TreeNode) bool
	walk = func(n *TreeNode) bool {
		if n == nil {
			return false
		}
		path = append(path, n.Val)
		if n.Val == target {
			return true
		}
		if walk(n.Left) || walk(n.Right) {
			return true
		}
		path = path[:len(path)-1] // this subtree doesn't contain the target
		return false
	}
	if walk(root) {
		return path
	}
	return nil
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(1, tn(2, tn(4, nil, nil), tn(5, nil, nil)), tn(3, nil, nil))
	fmt.Println(PathToNode(root, 5))
}
