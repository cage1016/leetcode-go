/*
 * @lc app=leetcode id=297 lang=golang
 *
 * [297] Serialize and Deserialize Binary Tree
 *
 * https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/
 *
 * algorithms
 * Hard (45.94%)
 * Likes:    2650
 * Dislikes: 131
 * Total Accepted:    293.1K
 * Total Submissions: 638K
 * Testcase Example:  '[1,2,3,null,null,4,5]'
 *
 * Serialization is the process of converting a data structure or object into a
 * sequence of bits so that it can be stored in a file or memory buffer, or
 * transmitted across a network connection link to be reconstructed later in
 * the same or another computer environment.
 * 
 * Design an algorithm to serialize and deserialize a binary tree. There is no
 * restriction on how your serialization/deserialization algorithm should work.
 * You just need to ensure that a binary tree can be serialized to a string and
 * this string can be deserialized to the original tree structure.
 * 
 * Example: 
 * 
 * 
 * You may serialize the following tree:
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * ⁠    / \
 * ⁠   4   5
 * 
 * as "[1,2,3,null,null,4,5]"
 * 
 * 
 * Clarification: The above format is the same as how LeetCode serializes a
 * binary tree. You do not necessarily need to follow this format, so please be
 * creative and come up with different approaches yourself.
 * 
 * Note: Do not use class member/global/static variables to store states. Your
 * serialize and deserialize algorithms should be stateless.
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

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
		return ""
	}
	
	s := []string{}
	var queue = []*TreeNode{}
	var node = &TreeNode{}
	queue = append(queue, root)
	
	for len(queue) > 0 {
		node, queue = queue[0], queue[1:]
		if node != nil {
			s = append(s, strconv.Itoa(node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}else{
			s = append(s, "null")
		}
	}

	return strings.Join(s, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    if len(data) == 0 {
		return nil
	}

	c := func(s string) int {
		v, _ := strconv.Atoi(s)
		return v
	}
	
	values := strings.Split(data,",")
	var queue = []*TreeNode{}
	root := &TreeNode{Val:c(values[0])}
	queue = append(queue, root)

	values = values[1:]
	for len(queue) > 0 && len(values) > 0{
		if values[0] != "null" {
			queue[0].Left = &TreeNode{Val:c(values[0])}
			queue = append(queue, queue[0].Left)
		}else{
			queue[0].Left = nil
		}

		if values[1] != "null" {
			queue[0].Right = &TreeNode{Val:c(values[1])}
			queue = append(queue, queue[0].Right)
		}else{
			queue[0].Right = nil
		}
		
		queue = queue[1:]
		values = values[2:]		
	}
    return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end

