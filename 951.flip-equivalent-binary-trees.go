/*
 * @lc app=leetcode id=951 lang=golang
 *
 * [951] Flip Equivalent Binary Trees
 *
 * https://leetcode.com/problems/flip-equivalent-binary-trees/description/
 *
 * algorithms
 * Medium (67.12%)
 * Likes:    2650
 * Dislikes: 110
 * Total Accepted:    209.4K
 * Total Submissions: 304.5K
 * Testcase Example:  '[1,2,3,4,5,6,null,null,null,7,8]\n[1,3,2,null,6,4,5,null,null,null,null,8,7]'
 *
 * For a binary tree T, we can define a flip operation as follows: choose any
 * node, and swap the left and right child subtrees.
 * 
 * A binary tree X is flip equivalent to a binary tree Y if and only if we can
 * make X equal to Y after some number of flip operations.
 * 
 * Given the roots of two binary trees root1 and root2, return true if the two
 * trees are flip equivalent or false otherwise.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: root1 = [1,2,3,4,5,6,null,null,null,7,8], root2 =
 * [1,3,2,null,6,4,5,null,null,null,null,8,7]
 * Output: true
 * Explanation: We flipped at nodes with values 1, 3, and 5.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: root1 = [], root2 = []
 * Output: true
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: root1 = [], root2 = [1]
 * Output: false
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The number of nodes in each tree is in the range [0, 100].
 * Each tree will have unique node values in the range [0, 99].
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
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		return (dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left)) || (dfs(node1.Left, node2.Left) && dfs(node1.Right, node2.Right))
	}
	return dfs(root1, root2)
}
// @lc code=end

