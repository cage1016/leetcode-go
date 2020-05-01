/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
 *
 * https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (64.78%)
 * Likes:    2186
 * Dislikes: 66
 * Total Accepted:    754.5K
 * Total Submissions: 1.2M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, find its maximum depth.
 * 
 * The maximum depth is the number of nodes along the longest path from the
 * root node down to the farthest leaf node.
 * 
 * Note: A leaf is a node with no children.
 * 
 * Example:
 * 
 * Given binary tree [3,9,20,null,null,15,7],
 * 
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * return its depth = 3.
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// max := math.MinInt64
	// dfs(root, 1, &max)
	// return max
	return bfs(root)
}

func dfs(node *TreeNode, level int, max *int) {
	if node.Left == nil && node.Right == nil {
		if level > *max {
			*max = level
		}
	}
	
	if node.Left != nil {
		dfs(node.Left, level + 1, max)
	}
	if node.Right != nil {
		dfs(node.Right, level + 1, max)
	}
}

func bfs(node *TreeNode) int {
	var queue = []*TreeNode{}
	queue = append(queue, node, nil)
	depth := 0

	for len(queue) > 0 {
		p := queue[0]
		if p == nil {
			depth ++
			queue = append(queue, nil)
			if queue[0] == nil && queue[1] == nil {
				break
			}
		}else{
			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
		}
		queue = queue[1:]
	}
	return depth
}
// @lc code=end

