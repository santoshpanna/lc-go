/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 *
 * https://leetcode.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (51.00%)
 * Likes:    16019
 * Dislikes: 423
 * Total Accepted:    2M
 * Total Submissions: 3.9M
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3'
 *
 * You are given an m x n integer matrix matrix with the following two
 * properties:
 *
 *
 * Each row is sorted in non-decreasing order.
 * The first integer of each row is greater than the last integer of the
 * previous row.
 *
 *
 * Given an integer target, return true if target is in matrix or false
 * otherwise.
 *
 * You must write a solution in O(log(m * n)) time complexity.
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 100
 * -10^4 <= matrix[i][j], target <= 10^4
 *
 *
 */
package main

import "fmt"

func main() {
	testCases := []struct {
		matrix [][]int
		target int
		output bool
	}{
		// {
		// 	matrix: [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}},
		// 	target: 3,
		// 	output: true,
		// },
		// {
		// 	matrix: [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}},
		// 	target: 13,
		// 	output: false,
		// },
		// {
		// 	matrix: [][]int{{1},{3}},
		// 	target: 0,
		// 	output: false,
		// },
		// {
		// 	matrix: [][]int{{1},{3}},
		// 	target: 2,
		// 	output: false,
		// },
		// {
		// 	matrix: [][]int{{1}},
		// 	target: 1,
		// 	output: true,
		// },
		// {
		// 	matrix: [][]int{{1},{3},{5}},
		// 	target: 3,
		// 	output: true,
		// },
		{
			matrix: [][]int{{1}, {3}, {5}, {7}, {9}, {11}, {13}, {15}, {17}, {19},    {21}, {23}, {25}, {27}, {29}, {31}, {33}, {35}, {37}, {39},    {41}, {43}, {45}, {47}, {49}, {51}, {53}, {55}, {57}, {59},    {61}, {63}, {65}, {67}, {69}, {71}, {73}, {75}, {77}, {79},    {81}, {83}, {85}, {87}, {89}, {91}, {93}, {95}, {97}, {99}},
			target: 3,
			output: true,
		},
	}

	for _, testCase := range testCases {
		result := searchMatrix(testCase.matrix, testCase.target)
		fmt.Println("Expected:", testCase.output)
		fmt.Println("Output:", result)
	}
}

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix) - 1

	mid := i + (j - i) / 2
	
	for i <= j {
		mid = i + (j - i) / 2
		if matrix[mid][0] <= target && matrix[mid][len(matrix[mid]) - 1] >= target {
			break
		} else if matrix[mid][0] < target {
			i = mid + 1
		} else{
			j = mid - 1
		}
	}

	if mid < 0 || mid >= len(matrix) {
		return false
	}

	i, j = 0, len(matrix[mid]) - 1

	for i <= j {
		midInner := i + (j - i) / 2

		if matrix[mid][midInner] == target {
			return true
		} else if matrix[mid][midInner] < target {
			i = midInner + 1
		} else {
			j = midInner - 1
		}
	}

	return false
}
// @lc code=end

