/*
 * @lc app=leetcode id=1189 lang=golang
 *
 * [1189] Maximum Number of Balloons
 *
 * https://leetcode.com/problems/maximum-number-of-balloons/description/
 *
 * algorithms
 * Easy (59.71%)
 * Likes:    1724
 * Dislikes: 107
 * Total Accepted:    226.3K
 * Total Submissions: 379K
 * Testcase Example:  '"nlaebolko"'
 *
 * Given a string text, you want to use the characters of text to form as many
 * instances of the word "balloon" as possible.
 * 
 * You can use each character in text at most once. Return the maximum number
 * of instances that can be formed.
 * 
 * 
 * Example 1:
 * 
 * 
 * 
 * 
 * Input: text = "nlaebolko"
 * Output: 1
 * 
 * 
 * Example 2:
 * 
 * 
 * 
 * 
 * Input: text = "loonbalxballpoon"
 * Output: 2
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: text = "leetcode"
 * Output: 0
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= text.length <= 10^4
 * text consists of lower case English letters only.
 * 
 * 
 * 
 * Note: This question is the same as  2287: Rearrange Characters to Make
 * Target String.
 * 
 */

// @lc code=start
func maxNumberOfBalloons(text string) int {
	lookup := make(map [rune]int)

	for _, r := range text {
		lookup[r]++
	}

	lookup['l'] = lookup['l']/2
	lookup['o'] = lookup['o']/2

	min := 1000000000

	for _, r := range "balon" {
		if lookup[r] < min {
			min = lookup[r]
		}
	}

	return min
}
// @lc code=end

