/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 *
 * https://leetcode.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Medium (47.44%)
 * Likes:    20383
 * Dislikes: 1050
 * Total Accepted:    2.1M
 * Total Submissions: 4.3M
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * Given an unsorted array of integers nums, return the length of the longest
 * consecutive elements sequence.
 * 
 * You must write an algorithm that runs in O(n) time.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [100,4,200,1,3,2]
 * Output: 4
 * Explanation: The longest consecutive elements sequence is [1, 2, 3, 4].
 * Therefore its length is 4.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [0,3,7,2,5,8,4,6,0,1]
 * Output: 9
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 * 
 * 
 */

package main

import "fmt"

func main() {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{100, 4, 200, 1, 3, 2},
			output: 4,
		},
		{
			input:  []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			output: 9,
		},
		{
			input:  []int{1, 2, 0, 1},
			output: 3,
		},
		{
			input:  []int{1, 2, 0, 1, 3},
			output: 4,
		},
	}

	for _, testCase := range testCases {
		res := longestConsecutive(testCase.input)
		fmt.Println("Input: ", testCase.input, " expected output: ", testCase.output, " output: ", res)
	}
}

// @lc code=start
func longestConsecutive(nums []int) int {
	lookup := make(map [int]bool)

	res := 0

	for _, num := range nums {
		lookup[num] = true
	}
	
	for num := range lookup {
		currentMax := 1
		nextNum := num + 1
		prevNum := num - 1

		if lookup[prevNum] == false {
			for lookup[nextNum] == true {
				currentMax = currentMax + 1
				nextNum = nextNum + 1
			}

			res = max(res, currentMax)
		}
	}

	return res
}
// @lc code=end

