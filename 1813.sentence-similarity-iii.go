/*
 * @lc app=leetcode id=1813 lang=golang
 *
 * [1813] Sentence Similarity III
 *
 * https://leetcode.com/problems/sentence-similarity-iii/description/
 *
 * algorithms
 * Medium (33.68%)
 * Likes:    663
 * Dislikes: 109
 * Total Accepted:    52.9K
 * Total Submissions: 121.9K
 * Testcase Example:  '"My name is Haley"\n"My Haley"'
 *
 * You are given two strings sentence1 and sentence2, each representing a
 * sentence composed of words. A sentence is a list of words that are separated
 * by a single space with no leading or trailing spaces. Each word consists of
 * only uppercase and lowercase English characters.
 *
 * Two sentences s1 and s2 are considered similar if it is possible to insert
 * an arbitrary sentence (possibly empty) inside one of these sentences such
 * that the two sentences become equal. Note that the inserted sentence must be
 * separated from existing words by spaces.
 *
 * For example,
 *
 *
 * s1 = "Hello Jane" and s2 = "Hello my name is Jane" can be made equal by
 * inserting "my name is" between "Hello" and "Jane" in s1.
 * s1 = "Frog cool" and s2 = "Frogs are cool" are not similar, since although
 * there is a sentence "s are" inserted into s1, it is not separated from
 * "Frog" by a space.
 *
 *
 * Given two sentences sentence1 and sentence2, return true if sentence1 and
 * sentence2 are similar. Otherwise, return false.
 *
 *
 * Example 1:
 *
 *
 * Input: sentence1 = "My name is Haley", sentence2 = "My Haley"
 *
 * Output: true
 *
 * Explanation:
 *
 * sentence2 can be turned to sentence1 by inserting "name is" between "My" and
 * "Haley".
 *
 *
 * Example 2:
 *
 *
 * Input: sentence1 = "of", sentence2 = "A lot of words"
 *
 * Output: false
 *
 * Explanation:
 *
 * No single sentence can be inserted inside one of the sentences to make it
 * equal to the other.
 *
 *
 * Example 3:
 *
 *
 * Input: sentence1 = "Eating right now", sentence2 = "Eating"
 *
 * Output: true
 *
 * Explanation:
 *
 * sentence2 can be turned to sentence1 by inserting "right now" at the end of
 * the sentence.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= sentence1.length, sentence2.length <= 100
 * sentence1 and sentence2 consist of lowercase and uppercase English letters
 * and spaces.
 * The words in sentence1 and sentence2 are separated by a single space.
 *
 *
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	testCases := []struct {
		sentence1 string
		sentence2 string
		expected bool
	}{
		{"My name is Haley", "My Haley", true},
		{"of", "A lot of words", false},
		{"Eating right now", "Eating", true},
		{"hello world python", "python", true},
		{"hello world python", "hello", true},
		{"hello world python", "world", false},
		{"hello world python", "hello python", true},
		{"abc def cdi", "def ghi", false},
		{"Game is ON", "Game are ON", false},
	}

	for _, testCase := range testCases {
		result := areSentencesSimilar(testCase.sentence1, testCase.sentence2)
		fmt.Printf("expected: %v, got: %v\n", testCase.expected, result)
	}
}

// @lc code=start
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	splitS1 := strings.Split(sentence1, " ")
	splitS2 := strings.Split(sentence2, " ")

	sizes1 := len(splitS1)
	sizes2 := len(splitS2)

	if sizes1 > sizes2 {
		splitS1, splitS2 = splitS2, splitS1
		sizes1, sizes2 = sizes2, sizes1
	}

	i, j := 0, 0

	for i < sizes1 && splitS1[i] == splitS2[i] {
		i++
	}

	for j < sizes1 - i && splitS1[sizes1 - 1 - j] == splitS2[sizes2 - 1 - j] {
		j++
	}

	return i+j >= sizes1
}
// @lc code=end

