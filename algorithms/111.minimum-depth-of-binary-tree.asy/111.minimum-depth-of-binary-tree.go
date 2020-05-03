/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
 *
 * https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (36.87%)
 * Likes:    1181
 * Dislikes: 636
 * Total Accepted:    388.1K
 * Total Submissions: 1.1M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, find its minimum depth.
 * 
 * The minimum depth is the number of nodes along the shortest path from the
 * root node down to the nearest leaf node.
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
 * return its minimum depth = 2.
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// min := math.MaxInt64
	// dfs(root, 1, &min)
	// return min
	return bfs(root)
}

func dfs(node *TreeNode, level int, min *int) {
	if node.Left == nil && node.Right == nil {
		if level < *min {
			*min = level
		}
	}
	
	if node.Left != nil {
		dfs(node.Left, level + 1, min)
	}
	if node.Right != nil {
		dfs(node.Right, level + 1, min)
	}
}

func bfs(node *TreeNode) int {
	var queue = []*TreeNode{}
	queue = append(queue, node, nil)
	depth := 0
	min := math.MaxInt64

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
			if p.Left == nil && p.Right == nil {
				if (depth + 1) < min {
					min = (depth + 1)
				}
			}
		}
		queue = queue[1:]
	}
	return min
}
// @lc code=end

