/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	lens := len(s)
	lent := len(t)

	if lens != lent {
		return false
	}
	
	freq := make(map[rune]int)

	for _, c := range s {
		freq[c]++
	}

	for _, c := range t {
		freq[c]--
		if freq[c] < 0 {
			return false
		}
	}

	return true
}
// @lc code=end

