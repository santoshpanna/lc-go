/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	m:= make(map[int]int)

	for i, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}

		m[num] = i
	}

	return false
}
// @lc code=end

