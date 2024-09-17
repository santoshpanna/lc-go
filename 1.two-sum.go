/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m:= make(map[int]int)
	count:= len(nums)

	for i := 0; i < count; i++ {
		complement:= target - nums[i]

		if j, ok := m[complement]; ok {
			return []int{j, i}
		}

		m[nums[i]] = i
	}

	return nil
}
// @lc code=end

