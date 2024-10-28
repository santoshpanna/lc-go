/*
 * @lc app=leetcode id=1277 lang=golang
 *
 * [1277] Count Square Submatrices with All Ones
 *
 * https://leetcode.com/problems/count-square-submatrices-with-all-ones/description/
 *
 * algorithms
 * Medium (75.59%)
 * Likes:    5051
 * Dislikes: 87
 * Total Accepted:    257.4K
 * Total Submissions: 336.8K
 * Testcase Example:  '[[0,1,1,1],[1,1,1,1],[0,1,1,1]]'
 *
 * Given a m * n matrix of ones and zeros, return how many square submatrices
 * have all ones.
 *
 *
 * Example 1:
 *
 *
 * Input: matrix =
 * [
 * [0,1,1,1],
 * [1,1,1,1],
 * [0,1,1,1]
 * ]
 * Output: 15
 * Explanation:
 * There are 10 squares of side 1.
 * There are 4 squares of side 2.
 * There is  1 square of side 3.
 * Total number of squares = 10 + 4 + 1 = 15.
 *
 *
 * Example 2:
 *
 *
 * Input: matrix =
 * [
 * ⁠ [1,0,1],
 * ⁠ [1,1,0],
 * ⁠ [1,1,0]
 * ]
 * Output: 7
 * Explanation:
 * There are 6 squares of side 1.
 * There is 1 square of side 2.
 * Total number of squares = 6 + 1 = 7.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr.length <= 300
 * 1 <= arr[0].length <= 300
 * 0 <= arr[i][j] <= 1
 *
 *
 */
package main

import "fmt"

func main() {
	mat := [][]int{{0,1,1,1},{1,1,1,1},{0,1,1,1}}
	fmt.Println(countSquares(mat))
}

// @lc code=start
func countSquares(matrix [][]int) int {
	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row + 1)

	for i := 0; i <= row; i++ {
			dp[i] = make([]int, col + 1)
	}

	ans := 0

	for i := 0; i < row; i++ {
			for j := 0; j < col; j++ {
					if i < row && j < col && matrix[i][j] == 1 {
							dp[i+1][j+1] = min(dp[i][j+1], dp[i+1][j], dp[i][j]) + 1
					}
					ans += dp[i+1][j+1]
			}
	}

	return ans
}
// @lc code=end

