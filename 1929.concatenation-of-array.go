/*
 * @lc app=leetcode id=1929 lang=golang
 *
 * [1929] Concatenation of Array
 */

// @lc code=start
func getConcatenation(nums []int) []int {
	arr:= make([]int, len(nums) * 2)

	for i, num := range nums {
		arr[i] = num
		arr[i + len(nums)] = num
	}

	return arr
}
// @lc code=end

