/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 *
 * https://leetcode.com/problems/binary-search/description/
 *
 * algorithms
 * Easy (58.37%)
 * Likes:    12003
 * Dislikes: 258
 * Total Accepted:    2.7M
 * Total Submissions: 4.6M
 * Testcase Example:  '[-1,0,3,5,9,12]\n9'
 *
 * Given an array of integers nums which is sorted in ascending order, and an
 * integer target, write a function to search target in nums. If target exists,
 * then return its index. Otherwise, return -1.
 * 
 * You must write an algorithm with O(log n) runtime complexity.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [-1,0,3,5,9,12], target = 9
 * Output: 4
 * Explanation: 9 exists in nums and its index is 4
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [-1,0,3,5,9,12], target = 2
 * Output: -1
 * Explanation: 2 does not exist in nums so return -1
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= nums.length <= 10^4
 * -10^4 < nums[i], target < 10^4
 * All the integers in nums are unique.
 * nums is sorted in ascending order.
 * 
 * 
 */
package main

import "fmt"

func main() {
	testCases := []struct {
		nums []int
		target int
		output int
	}{
		{
			nums: []int{-1,0,3,5,9,12},
			target: 9,
			output: 4,
		},
		{
			nums: []int{-1,0,3,5,9,12},
			target: 2,
			output: -1,
		},
	}

	for _, testCase := range testCases {
		result := search(testCase.nums, testCase.target)
		fmt.Println("Expected:", testCase.output)
		fmt.Println("Output:", result)
	}
}

// @lc code=start
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high - low) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
// @lc code=end

