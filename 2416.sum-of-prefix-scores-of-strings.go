/*
 * @lc app=leetcode id=2416 lang=golang
 *
 * [2416] Sum of Prefix Scores of Strings
 *
 * https://leetcode.com/problems/sum-of-prefix-scores-of-strings/description/
 *
 * algorithms
 * Hard (46.48%)
 * Likes:    1005
 * Dislikes: 88
 * Total Accepted:    77.4K
 * Total Submissions: 135.1K
 * Testcase Example:  '["abc","ab","bc","b"]'
 *
 * You are given an array words of size n consisting of non-empty strings.
 *
 * We define the score of a string term as the number of strings words[i] such
 * that term is a prefix of words[i].
 *
 *
 * For example, if words = ["a", "ab", "abc", "cab"], then the score of "ab" is
 * 2, since "ab" is a prefix of both "ab" and "abc".
 *
 *
 * Return an array answer of size n where answer[i] is the sum of scores of
 * every non-empty prefix of words[i].
 *
 * Note that a string is considered as a prefix of itself.
 *
 *
 * Example 1:
 *
 *
 * Input: words = ["abc","ab","bc","b"]
 * Output: [5,4,3,2]
 * Explanation: The answer for each string is the following:
 * - "abc" has 3 prefixes: "a", "ab", and "abc".
 * - There are 2 strings with the prefix "a", 2 strings with the prefix "ab",
 * and 1 string with the prefix "abc".
 * The total is answer[0] = 2 + 2 + 1 = 5.
 * - "ab" has 2 prefixes: "a" and "ab".
 * - There are 2 strings with the prefix "a", and 2 strings with the prefix
 * "ab".
 * The total is answer[1] = 2 + 2 = 4.
 * - "bc" has 2 prefixes: "b" and "bc".
 * - There are 2 strings with the prefix "b", and 1 string with the prefix
 * "bc".
 * The total is answer[2] = 2 + 1 = 3.
 * - "b" has 1 prefix: "b".
 * - There are 2 strings with the prefix "b".
 * The total is answer[3] = 2.
 *
 *
 * Example 2:
 *
 *
 * Input: words = ["abcd"]
 * Output: [4]
 * Explanation:
 * "abcd" has 4 prefixes: "a", "ab", "abc", and "abcd".
 * Each prefix has a score of one, so the total is answer[0] = 1 + 1 + 1 + 1 =
 * 4.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= words.length <= 1000
 * 1 <= words[i].length <= 1000
 * words[i] consists of lowercase English letters.
 *
 *
 */

package main

import "fmt"

func main() {
	testCases := []struct {
		words  []string
		output []int
	}{
		{words: []string{"abc", "ab", "bc", "b"}, output: []int{5, 4, 3, 2}},
		{words: []string{"abcd"}, output: []int{4}},
	}

	for _, tc := range testCases {
		result := sumPrefixScores(tc.words)
		fmt.Printf("expected: %v, got: %v\n", tc.output, result)
	}
}

// @lc code=start
type Trie struct {
	Children map[rune]*Trie
	Count    int
}

func newTrie() *Trie {
	return &Trie{
		Children: make(map[rune]*Trie),
		Count:    0,
	}
}

func (t *Trie) insert(word string) {
	node := t
	for _, c := range word {
		if _, ok := node.Children[c]; !ok {
			node.Children[c] = newTrie()
		}
		node = node.Children[c]
		node.Count++
	}
}

func sumPrefixScores(words []string) []int {
	trie := newTrie()
	for _, word := range words {
		trie.insert(word)
	}

	result := make([]int, len(words))
	for i, word := range words {
		t := trie
		for _, c := range word {
			t = t.Children[c]
			result[i] += t.Count
		}
	}

	return result
}
// @lc code=end

