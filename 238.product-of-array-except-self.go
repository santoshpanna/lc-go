/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 *
 * https://leetcode.com/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (66.69%)
 * Likes:    22913
 * Dislikes: 1446
 * Total Accepted:    3M
 * Total Submissions: 4.4M
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given an integer array nums, return an array answer such that answer[i] is
 * equal to the product of all the elements of nums except nums[i].
 *
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 *
 * You must write an algorithm that runs in O(n) time and without using the
 * division operation.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3,4]
 * Output: [24,12,8,6]
 * Example 2:
 * Input: nums = [-1,1,0,-3,3]
 * Output: [0,0,9,0,0]
 *
 *
 * Constraints:
 *
 *
 * 2 <= nums.length <= 10^5
 * -30 <= nums[i] <= 30
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
 * integer.
 *
 *
 *
 * Follow up: Can you solve the problem in O(1) extra space complexity? (The
 * output array does not count as extra space for space complexity analysis.)
 *
 */
package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{nums: []int{1,2,3,4}, expected: []int{24,12,8,6}},
		{nums: []int{-1,1,0,-3,3}, expected: []int{0,0,9,0,0}},
	}

	for _, tc := range testCases {
		result := productExceptSelf(tc.nums)
		fmt.Printf("nums: %v, expected: %v, got: %v\n", tc.nums, tc.expected, result)
	}
}

// @lc code=start
func productExceptSelf(nums []int) []int {
	product := 1
	zeroCount := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			product *= nums[i]
		}

		if nums[i] == 0 {
			zeroCount++
		}
	}

	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && zeroCount == 1 {
			res[i] = product
		} else if zeroCount >= 1 {
			res[i] = 0
		} else {
			res[i] = product / nums[i]
		}
	}

	return res
}
// @lc code=end

