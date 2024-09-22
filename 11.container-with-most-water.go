/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 *
 * https://leetcode.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (56.13%)
 * Likes:    29579
 * Dislikes: 1838
 * Total Accepted:    3.3M
 * Total Submissions: 5.9M
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * You are given an integer array height of length n. There are n vertical
 * lines drawn such that the two endpoints of the i^th line are (i, 0) and (i,
 * height[i]).
 *
 * Find two lines that together with the x-axis form a container, such that the
 * container contains the most water.
 *
 * Return the maximum amount of water a container can store.
 *
 * Notice that you may not slant the container.
 *
 *
 * Example 1:
 *
 *
 * Input: height = [1,8,6,2,5,4,8,3,7]
 * Output: 49
 * Explanation: The above vertical lines are represented by array
 * [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the
 * container can contain is 49.
 *
 *
 * Example 2:
 *
 *
 * Input: height = [1,1]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * n == height.length
 * 2 <= n <= 10^5
 * 0 <= height[i] <= 10^4
 *
 *
 */
package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		input []int
		expected int
	}{
		{input: []int{1,8,6,2,5,4,8,3,7}, expected: 49},
		{input: []int{1,1}, expected: 1},
	}

	for _, testCase := range testCases {
		testCase := testCase
		result := maxArea(testCase.input)
		fmt.Printf("expected: %v, result: %v\n", testCase.expected, result)
	}
}

// @lc code=start
func maxArea(height []int) int {
	res := 0

	for i, j := 0, len(height) - 1; i < j; {
		res = max(min(height[i], height[j]) * (j - i), res)

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return res
}
// @lc code=end

