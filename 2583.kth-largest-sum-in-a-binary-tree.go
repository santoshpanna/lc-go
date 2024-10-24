/*
 * @lc app=leetcode id=2583 lang=golang
 *
 * [2583] Kth Largest Sum in a Binary Tree
 *
 * https://leetcode.com/problems/kth-largest-sum-in-a-binary-tree/description/
 *
 * algorithms
 * Medium (47.98%)
 * Likes:    853
 * Dislikes: 35
 * Total Accepted:    117.5K
 * Total Submissions: 204.5K
 * Testcase Example:  '[5,8,9,2,1,3,7,4,6]\n2'
 *
 * You are given the root of a binary tree and a positive integer k.
 *
 * The level sum in the tree is the sum of the values of the nodes that are on
 * the same level.
 *
 * Return the k^th largest level sum in the tree (not necessarily distinct). If
 * there are fewer than k levels in the tree, return -1.
 *
 * Note that two nodes are on the same level if they have the same distance
 * from the root.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [5,8,9,2,1,3,7,4,6], k = 2
 * Output: 13
 * Explanation: The level sums are the following:
 * - Level 1: 5.
 * - Level 2: 8 + 9 = 17.
 * - Level 3: 2 + 1 + 3 + 7 = 13.
 * - Level 4: 4 + 6 = 10.
 * The 2^nd largest level sum is 13.
 *
 *
 * Example 2:
 *
 *
 * Input: root = [1,2,null,3], k = 1
 * Output: 3
 * Explanation: The largest level sum is 3.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is n.
 * 2 <= n <= 10^5
 * 1 <= Node.val <= 10^6
 * 1 <= k <= n
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
package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	testCases := []struct{
		root *TreeNode
		k int
		want int64
	} {
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			k: 2,
			want: 13,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			k: 1,
			want: 3,
		},
	}

	for _, tc := range testCases {
		got := kthLargestLevelSum(tc.root, tc.k)
		fmt.Println("got =", got, "want =",tc.want, got == tc.want)
	}
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	levelSums := []int64{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(levelSums) <= level {
			levelSums = append(levelSums, int64(node.Val))
		} else {
			levelSums[level] += int64(node.Val)
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)
	sort.Slice(levelSums, func(i, j int) bool {
		return levelSums[i] > levelSums[j]
	})

	if k > len(levelSums) {
		return -1
	}

	return levelSums[k-1]
}
// @lc code=end

