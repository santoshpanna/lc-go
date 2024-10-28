/*
 * @lc app=leetcode id=2501 lang=golang
 *
 * [2501] Longest Square Streak in an Array
 *
 * https://leetcode.com/problems/longest-square-streak-in-an-array/description/
 *
 * algorithms
 * Medium (40.21%)
 * Likes:    762
 * Dislikes: 25
 * Total Accepted:    87.9K
 * Total Submissions: 175.1K
 * Testcase Example:  '[4,3,6,16,8,2]'
 *
 * You are given an integer array nums. A subsequence of nums is called a
 * square streak if:
 * 
 * 
 * The length of the subsequence is at least 2, and
 * after sorting the subsequence, each element (except the first element) is
 * the square of the previous number.
 * 
 * 
 * Return the length of the longest square streak in nums, or return -1 if
 * there is no square streak.
 * 
 * A subsequence is an array that can be derived from another array by deleting
 * some or no elements without changing the order of the remaining elements.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [4,3,6,16,8,2]
 * Output: 3
 * Explanation: Choose the subsequence [4,16,2]. After sorting it, it becomes
 * [2,4,16].
 * - 4 = 2 * 2.
 * - 16 = 4 * 4.
 * Therefore, [4,16,2] is a square streak.
 * It can be shown that every subsequence of length 4 is not a square streak.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [2,3,5,6,7]
 * Output: -1
 * Explanation: There is no square streak in nums so return -1.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 2 <= nums.length <= 10^5
 * 2 <= nums[i] <= 10^5
 * 
 * 
 */

// @lc code=start
func longestSquareStreak(nums []int) int {
	lookup := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		lookup[nums[i]] = i
	} 
	
	sort.Slice(nums, func(a, b int) bool{
		return nums[a] < nums[b]
	})

	i, largestStreak := 0, -1

	for i < len(nums) {
		streak := 0
		sq := nums[i] * nums[i]

		for lookup[sq] > i {
			streak++
			sq *= sq
		}

		largestStreak = max(largestStreak, streak)

		i++
	}

	if largestStreak < 1 {
		return -1
	}

	return largestStreak
}
// @lc code=end

