/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 *
 * https://leetcode.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (75.35%)
 * Likes:    13063
 * Dislikes: 468
 * Total Accepted:    1.8M
 * Total Submissions: 2.4M
 * Testcase Example:  '5'
 *
 * Given an integer numRows, return the first numRows of Pascal's triangle.
 *
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it as shown:
 *
 *
 * Example 1:
 * Input: numRows = 5
 * Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
 * Example 2:
 * Input: numRows = 1
 * Output: [[1]]
 *
 *
 * Constraints:
 *
 *
 * 1 <= numRows <= 30
 *
 *
 */
package main

import "fmt"
 
 func main() {
	 testCases := []struct {
		 nums     int
		 expected [][]int
	 }{
		 {nums: 5, expected: [][]int{{1}, {1,1}, {1,2,1}, {1,3,3,1}, {1,4,6,4,1}}},
		 {nums: 1, expected: [][]int{{1}}},
	 }
 
	 for _, tc := range testCases {
		 result := generate(tc.nums)
		 fmt.Printf("nums: %v, expected: %v, got: %v\n", tc.nums, tc.expected, result)
	 }
 }

// @lc code=start
func generate(numRows int) [][]int {
	res := [][]int{}

	res = append(res, []int{1})

	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1

		for j:=1; j<i; j++ {
			l := res[i-1][j-1]
			r := res[i-1][j]

			row[j] = l + r
		}

		res = append(res, row)
	}

	return res
}
// @lc code=end

