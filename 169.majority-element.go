/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (65.20%)
 * Likes:    19558
 * Dislikes: 653
 * Total Accepted:    3.4M
 * Total Submissions: 5.1M
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array nums of size n, return the majority element.
 * 
 * The majority element is the element that appears more than ⌊n / 2⌋ times.
 * You may assume that the majority element always exists in the array.
 * 
 * 
 * Example 1:
 * Input: nums = [3,2,3]
 * Output: 3
 * Example 2:
 * Input: nums = [2,2,1,1,1,2,2]
 * Output: 2
 * 
 * 
 * Constraints:
 * 
 * 
 * n == nums.length
 * 1 <= n <= 5 * 10^4
 * -10^9 <= nums[i] <= 10^9
 * 
 * 
 * 
 * Follow-up: Could you solve the problem in linear time and in O(1) space?
 */

// @lc code=start
func majorityElement(nums []int) int {
	count := 0
	num := 0

	for i := 0; i < len(nums); i++ {
		if (count == 0) {
			num = nums[i]
			count++
		} else if num == nums[i] {
			count++
		} else {
			count--
		}
	}

	return num
}
// @lc code=end

