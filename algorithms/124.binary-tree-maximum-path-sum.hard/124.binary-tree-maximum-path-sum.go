/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
 *
 * https://leetcode.com/problems/binary-tree-maximum-path-sum/description/
 *
 * algorithms
 * Hard (33.07%)
 * Likes:    3195
 * Dislikes: 258
 * Total Accepted:    336.9K
 * Total Submissions: 1M
 * Testcase Example:  '[1,2,3]'
 *
 * Given a non-empty binary tree, find the maximum path sum.
 * 
 * For this problem, a path is defined as any sequence of nodes from some
 * starting node to any node in the tree along the parent-child connections.
 * The path must contain at least one node and does not need to go through the
 * root.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,2,3]
 * 
 * ⁠      1
 * ⁠     / \
 * ⁠    2   3
 * 
 * Output: 6
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [-10,9,20,null,null,15,7]
 * 
 * -10
 * / \
 * 9  20
 * /  \
 * 15   7
 * 
 * Output: 42
 * 
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
func maxPathSum(root *TreeNode) int {
	max := math.MinInt64
	helper(root, &max)
	return max
}

func helper(node *TreeNode, max *int) int{
	if node == nil {
		return 0
	}

	left := helper(node.Left, max)
	right := helper(node.Right, max)

	if left < 0 {
		left = 0
	}
	if right < 0 {
		right = 0
	}

	sum := node.Val + left + right
	if sum > *max {
		*max = sum
	}

	if left > right {
		return left + node.Val
	}
	return right + node.Val
}
// @lc code=end

