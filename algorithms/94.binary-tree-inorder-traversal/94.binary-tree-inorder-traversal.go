/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (61.70%)
 * Likes:    2712
 * Dislikes: 117
 * Total Accepted:    674.5K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the inorder traversal of its nodes' values.
 * 
 * Example:
 * 
 * 
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 * 
 * Output: [1,3,2]
 * 
 * Follow up: Recursive solution is trivial, could you do it iteratively?
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
 func inorderTraversal(root *TreeNode) []int {
	ret := make([]int,0)
    if root == nil {
        return ret
    }
    dfs(root, &ret)
    return ret
}

func dfs(node *TreeNode, ret *[]int) {
    if node == nil {
        return
    }
    dfs(node.Left, ret)
    (*ret) = append(*ret, node.Val)
    dfs(node.Right, ret)
}
// @lc code=end

