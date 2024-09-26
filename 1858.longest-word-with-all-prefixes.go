/*
 * @lc app=leetcode id=1858 lang=golang
 *
 * [1858] Longest Word With All Prefixes
 *
 * https://leetcode.com/problems/longest-word-with-all-prefixes/description/
 *
 * algorithms
 * Medium (68.76%)
 * Likes:    194
 * Dislikes: 6
 * Total Accepted:    11.6K
 * Total Submissions: 16.4K
 * Testcase Example:  '["k","ki","kir","kira","kiran"]'
 *
 * Given an array of strings words, find the longest string in words such that
 * every prefix of it is also in words.
 *
 *
 * For example, let words = ["a", "app", "ap"]. The string "app" has prefixes
 * "ap" and "a", all of which are in words.
 *
 *
 * Return the string described above. If there is more than one string with the
 * same length, return the lexicographically smallest one, and if no string
 * exists, return "".
 *
 *
 * Example 1:
 *
 *
 * Input: words = ["k","ki","kir","kira", "kiran"]
 * Output: "kiran"
 * Explanation: "kiran" has prefixes "kira", "kir", "ki", and "k", and all of
 * them appear in words.
 *
 *
 * Example 2:
 *
 *
 * Input: words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
 * Output: "apple"
 * Explanation: Both "apple" and "apply" have all their prefixes in words.
 * However, "apple" is lexicographically smaller, so we return that.
 *
 *
 * Example 3:
 *
 *
 * Input: words = ["abc", "bc", "ab", "qwe"]
 * Output: ""
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= words.length <= 10^5
 * 1 <= words[i].length <= 10^5
 * 1 <= sum(words[i].length) <= 10^5
 * words[i] consists only of lowercase English letters.
 *
 *
 */

package main

import "fmt"

func main() {
	testCases := []struct {
		input []string
		output string
	}{
		{ input: []string{"k","ki","kir","kira", "kiran"}, output: "kiran" },
		{ input: []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}, output: "apple" },
		{ input: []string{"abc", "bc", "ab", "qwe"}, output: "" },
	}

	for _, testCase := range testCases {
		res := longestWord(testCase.input)
		fmt.Println("Test case:", testCase.input, res, testCase.output)
	}
}

// @lc code=start
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd bool
}

func insert(root *TrieNode, word string) {
	node := root
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func search(root *TrieNode, word string) bool {
	node := root
	for _, ch := range word {
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
		if !node.isEnd {
			return false
		}
	}
	return true
}

func longestWord(words []string) string {
	root := &TrieNode{children: make(map[rune]*TrieNode)}
	for _, word := range words {
		insert(root, word)
	}

	res := ""
	for _, word := range words {
		if search(root, word) {
			if len(word) > len(res) || (len(word) == len(res) && word < res) {
				res = word
			}
		}
	}
	return res	
}
// @lc code=end

