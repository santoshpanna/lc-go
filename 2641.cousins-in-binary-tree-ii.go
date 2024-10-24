/*
 * @lc app=leetcode id=2641 lang=golang
 *
 * [2641] Cousins in Binary Tree II
 *
 * https://leetcode.com/problems/cousins-in-binary-tree-ii/description/
 *
 * algorithms
 * Medium (67.39%)
 * Likes:    720
 * Dislikes: 23
 * Total Accepted:    40.4K
 * Total Submissions: 56.2K
 * Testcase Example:  '[5,4,9,1,10,null,7]'
 *
 * Given the root of a binary tree, replace the value of each node in the tree
 * with the sum of all its cousins' values.
 * 
 * Two nodes of a binary tree are cousins if they have the same depth with
 * different parents.
 * 
 * Return the root of the modified tree.
 * 
 * Note that the depth of a node is the number of edges in the path from the
 * root node to it.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: root = [5,4,9,1,10,null,7]
 * Output: [0,0,0,7,7,null,11]
 * Explanation: The diagram above shows the initial binary tree and the binary
 * tree after changing the value of each node.
 * - Node with value 5 does not have any cousins so its sum is 0.
 * - Node with value 4 does not have any cousins so its sum is 0.
 * - Node with value 9 does not have any cousins so its sum is 0.
 * - Node with value 1 has a cousin with value 7 so its sum is 7.
 * - Node with value 10 has a cousin with value 7 so its sum is 7.
 * - Node with value 7 has cousins with values 1 and 10 so its sum is 11.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: root = [3,1,2]
 * Output: [0,0,0]
 * Explanation: The diagram above shows the initial binary tree and the binary
 * tree after changing the value of each node.
 * - Node with value 3 does not have any cousins so its sum is 0.
 * - Node with value 1 does not have any cousins so its sum is 0.
 * - Node with value 2 does not have any cousins so its sum is 0.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The number of nodes in the tree is in the range [1, 10^5].
 * 1 <= Node.val <= 10^4
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
func replaceValueInTree(root *TreeNode) *TreeNode {
	levelSum := make(map[int]int)
	
	// First DFS to calculate the level sum
	var dfsLevelSum func(node *TreeNode, level int)
	dfsLevelSum = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		levelSum[level] += node.Val
		dfsLevelSum(node.Left, level+1)
		dfsLevelSum(node.Right, level+1)
	}
	dfsLevelSum(root, 0)
	
	// Second DFS to update node values with cousin sum
	var dfsUpdate func(node *TreeNode, parent *TreeNode, level int)
	dfsUpdate = func(node *TreeNode, val int, level int) {
		if node == nil {
			return
		}
		node.Val = val

		if level+1 < len(levelSum) {
			cousinSum := levelSum[level+1]
		}

		if node.Left != nil {
			cousinSum -= node.Left.Val
		}
		if node.Right != nil {
			cousinSum -= node.Right.Val
		}

		dfsUpdate(node.Left, cousinSum, level+1)
		dfsUpdate(node.Right, cousinSum, level+1)
	}
	dfsUpdate(root, nil, 0)
	
	return root
}
// @lc code=end

