/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (35.32%)
 * Likes:    31313
 * Dislikes: 2921
 * Total Accepted:    3.9M
 * Total Submissions: 11.1M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an integer array nums, return all the triplets [nums[i], nums[j],
 * nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] +
 * nums[k] == 0.
 *
 * Notice that the solution set must not contain duplicate triplets.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [-1,0,1,2,-1,-4]
 * Output: [[-1,-1,2],[-1,0,1]]
 * Explanation:
 * nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
 * nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
 * nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
 * The distinct triplets are [-1,0,1] and [-1,-1,2].
 * Notice that the order of the output and the order of the triplets does not
 * matter.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,1,1]
 * Output: []
 * Explanation: The only possible triplet does not sum up to 0.
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [0,0,0]
 * Output: [[0,0,0]]
 * Explanation: The only possible triplet sums up to 0.
 *
 *
 *
 * Constraints:
 *
 *
 * 3 <= nums.length <= 3000
 * -10^5 <= nums[i] <= 10^5
 *
 *
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	testCases := []struct {
		nums []int
		expected [][]int
	} {
		// {[]int{-1,0,1,2,-1,-4}, [][]int{{-1,-1,2},{-1,0,1}}},
		// {[]int{0,1,1}, [][]int{}},
		{[]int{0,0,0,0}, [][]int{{0,0,0}}},
		// {[]int{-1,0,1,2,-1,-4,-2,-3,3,0,4}, [][]int{{-4,0,4},{-4,1,3},{-3,-1,4},{-3,0,3},{-3,1,2},{-2,-1,3},{-2,0,2},{-1,-1,2},{-1,0,1}}},
	}

	for _, testCase := range testCases {
		result := threeSum(testCase.nums)
		fmt.Printf("expected: %v, returned: %v\n", testCase.expected, result)
	}
}

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int{}

	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j, k := i+1, len(nums) - 1; j < k; {
			if nums[i] + nums[j] + nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})

				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}

				j++
				k--
			}	else if nums[i] + nums[j] + nums[k] > 0 {
				k--
			} else if nums[i] + nums[j] + nums[k] < 0 {
				j++
			}
		}
	}

	return res
}
// @lc code=end

