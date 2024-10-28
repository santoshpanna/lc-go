/*
 * @lc app=leetcode id=2458 lang=golang
 *
 * [2458] Height of Binary Tree After Subtree Removal Queries
 *
 * https://leetcode.com/problems/height-of-binary-tree-after-subtree-removal-queries/description/
 *
 * algorithms
 * Hard (41.32%)
 * Likes:    1237
 * Dislikes: 27
 * Total Accepted:    60.8K
 * Total Submissions: 122.1K
 * Testcase Example:  '[1,3,4,2,null,6,5,null,null,null,null,null,7]\n[4]'
 *
 * You are given the root of a binary tree with n nodes. Each node is assigned
 * a unique value from 1 to n. You are also given an array queries of size m.
 *
 * You have to perform m independent queries on the tree where in the i^th
 * query you do the following:
 *
 *
 * Remove the subtree rooted at the node with the value queries[i] from the
 * tree. It is guaranteed that queries[i] will not be equal to the value of the
 * root.
 *
 *
 * Return an array answer of size m where answer[i] is the height of the tree
 * after performing the i^th query.
 *
 * Note:
 *
 *
 * The queries are independent, so the tree returns to its initial state after
 * each query.
 * The height of a tree is the number of edges in the longest simple path from
 * the root to some node in the tree.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,3,4,2,null,6,5,null,null,null,null,null,7], queries = [4]
 * Output: [2]
 * Explanation: The diagram above shows the tree after removing the subtree
 * rooted at node with value 4.
 * The height of the tree is 2 (The path 1 -> 3 -> 2).
 *
 *
 * Example 2:
 *
 *
 * Input: root = [5,8,9,2,1,3,7,4,6], queries = [3,2,4,8]
 * Output: [3,2,3,2]
 * Explanation: We have the following queries:
 * - Removing the subtree rooted at node with value 3. The height of the tree
 * becomes 3 (The path 5 -> 8 -> 2 -> 4).
 * - Removing the subtree rooted at node with value 2. The height of the tree
 * becomes 2 (The path 5 -> 8 -> 1).
 * - Removing the subtree rooted at node with value 4. The height of the tree
 * becomes 3 (The path 5 -> 8 -> 2 -> 6).
 * - Removing the subtree rooted at node with value 8. The height of the tree
 * becomes 2 (The path 5 -> 9 -> 3).
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is n.
 * 2 <= n <= 10^5
 * 1 <= Node.val <= n
 * All the values in the tree are unique.
 * m == queries.length
 * 1 <= m <= min(n, 10^4)
 * 1 <= queries[i] <= n
 * queries[i] != root.val
 *
 *
 */

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{Val: 2},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{
				Val: 5,
				Right : &TreeNode{Val: 7},
			},
		},
	}
	queries := []int{4}
	fmt.Println(treeQueries(root, queries))
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func treeQueries(root *TreeNode, queries []int) []int {
	heightMap := make(map[int]int)
	heightCache := make(map[*TreeNode]int)

	var computeHeight func(*TreeNode) int
	computeHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if height, ok := heightCache[node]; ok {
			return height
		}
		leftHeight := computeHeight(node.Left)
		rightHeight := computeHeight(node.Right)
		heightCache[node] = 1 + max(leftHeight, rightHeight)
		return heightCache[node]
	}
	computeHeight(root)

	var dfs func(*TreeNode, int, int, map[*TreeNode]int, map[int]int)

	dfs = func(node *TreeNode, depth int, maxDepth int, heightCache map[*TreeNode]int, heightMap map[int]int) {
		if node == nil {
			return
		}
		heightMap[node.Val] = maxDepth
		dfs(node.Left, depth+1, max(maxDepth, depth + 1 + computeHeight(node.Right)), heightCache, heightMap)
		dfs(node.Right, depth+1, max(maxDepth, depth + 1 + computeHeight(node.Left)), heightCache, heightMap)
	}

	dfs(root, 0, 0, heightCache, heightMap)

	var ans []int

	for _, query := range queries {
		ans = append(ans, heightMap[query] - 1)
	}
	return ans
}
// @lc code=end

