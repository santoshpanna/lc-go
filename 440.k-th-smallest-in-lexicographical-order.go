/*
 * @lc app=leetcode id=440 lang=golang
 *
 * [440] K-th Smallest in Lexicographical Order
 *
 * https://leetcode.com/problems/k-th-smallest-in-lexicographical-order/description/
 *
 * algorithms
 * Hard (32.72%)
 * Likes:    954
 * Dislikes: 104
 * Total Accepted:    42.4K
 * Total Submissions: 121.2K
 * Testcase Example:  '13\n2'
 *
 * Given two integers n and k, return the k^th lexicographically smallest
 * integer in the range [1, n].
 *
 *
 * Example 1:
 *
 *
 * Input: n = 13, k = 2
 * Output: 10
 * Explanation: The lexicographical order is [1, 10, 11, 12, 13, 2, 3, 4, 5, 6,
 * 7, 8, 9], so the second smallest number is 10.
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1, k = 1
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= k <= n <= 10^9
 *
 *
 */

package main

import "fmt"

func main() {
	testCases := []struct {
		n int
		k int
		want int
	}{
		{13, 2, 10},
		{1, 1, 1},
		{10, 3, 2},
		{100, 10, 17},
		{804289384, 42641503, 138377349},
	}

	for _, tc := range testCases {
		got := findKthNumber(tc.n, tc.k)
		fmt.Printf("got: %v, want: %v\n", got, tc.want)
	}
}

// @lc code=start
func calSteps(n, n1, n2 int) int {
	steps := 0
	for n1 <= n {
		steps += min(n+1, n2) - n1
		n1 *= 10
		n2 *= 10
	}
	return steps
}

func findKthNumber(n int, k int) int {
	cur := 1
	k--

	for k > 0 {
		steps := calSteps(n, cur, cur+1)
		if steps <= k {
			cur++
			k -= steps
		} else {
			cur *= 10
			k--
		}
	}

	return cur
}
// @lc code=end

