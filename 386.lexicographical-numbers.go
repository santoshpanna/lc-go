/*
 * @lc app=leetcode id=386 lang=golang
 *
 * [386] Lexicographical Numbers
 *
 * https://leetcode.com/problems/lexicographical-numbers/description/
 *
 * algorithms
 * Medium (64.72%)
 * Likes:    2272
 * Dislikes: 166
 * Total Accepted:    167.5K
 * Total Submissions: 245K
 * Testcase Example:  '13'
 *
 * Given an integer n, return all the numbers in the range [1, n] sorted in
 * lexicographical order.
 *
 * You must write an algorithm that runs in O(n) time and uses O(1) extra
 * space.
 *
 *
 * Example 1:
 * Input: n = 13
 * Output: [1,10,11,12,13,2,3,4,5,6,7,8,9]
 * Example 2:
 * Input: n = 2
 * Output: [1,2]
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 5 * 10^4
 *
 *
 */
package main

import "fmt"

func main() {
	testCases := []struct {
		nums     int
		expected []int
	}{
		{nums: 13, expected: []int{1,10,11,12,13,2,3,4,5,6,7,8,9}},
		{nums: 2, expected: []int{1,2}},
		{nums: 911, expected: []int{1,2}},
		{nums: 50, expected: []int{1,2}},
	}

	for _, tc := range testCases {
		result := lexicalOrder(tc.nums)
		fmt.Printf("nums: %v, expected: %v, got: %v\n", tc.nums, tc.expected, result)
	}
}

// @lc code=start
// func populateRes(i int, n int, res *[]int) {
// 	if i > n {
// 		return
// 	}

// 	*res = append(*res, i)

// 	for j := 0; j < 10; j++ {
// 		num := (i * 10) + j
// 		if num > n {
// 			return
// 		}
// 		populateRes(num, n, res)
// 	}
// }

// func lexicalOrder(n int) []int {
// 	res := []int{}

// 	for i := 1; i<10; i++ {
// 		populateRes(i, n, &res)
// 	}

// 	return res
// }

func lexicalOrder(n int) []int {
	res := []int{}
	val := 1

	for i := 1; i <= n; i++ {
		res = append(res, val)
		if (val * 10) <= n {
			val = val * 10
		} else {
			for val % 10 == 9 || val >= n {
				val = val / 10
			}
			val++
		}
	}

	return res
}
// @lc code=end

