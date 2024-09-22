/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (54.13%)
 * Likes:    31475
 * Dislikes: 1195
 * Total Accepted:    5.2M
 * Total Submissions: 9.7M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * You are given an array prices where prices[i] is the price of a given stock
 * on the i^th day.
 *
 * You want to maximize your profit by choosing a single day to buy one stock
 * and choosing a different day in the future to sell that stock.
 *
 * Return the maximum profit you can achieve from this transaction. If you
 * cannot achieve any profit, return 0.
 *
 *
 * Example 1:
 *
 *
 * Input: prices = [7,1,5,3,6,4]
 * Output: 5
 * Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit
 * = 6-1 = 5.
 * Note that buying on day 2 and selling on day 1 is not allowed because you
 * must buy before you sell.
 *
 *
 * Example 2:
 *
 *
 * Input: prices = [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transactions are done and the max profit =
 * 0.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= prices.length <= 10^5
 * 0 <= prices[i] <= 10^4
 *
 *
 */

package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		input []int
		expected int
	}{
		{input: []int{7,1,5,3,6,4}, expected: 5},
		{input: []int{7,6,4,3,1}, expected: 0},
		{input: []int{1,4,2}, expected: 3},
		{input: []int{2,1,4}, expected: 3},
	}

	for _, testCase := range testCases {
		testCase := testCase
		result := maxProfit(testCase.input)
		fmt.Printf("expected: %v, returned: %v\n", testCase.expected, result)
	}
}

// @lc code=start
func maxProfit(prices []int) int {
	profit := 0

	boughtOn := 0
	soldOn := boughtOn + 1

	for soldOn < len(prices) {
		tempProfit := prices[soldOn] - prices[boughtOn]

		if prices[boughtOn] < prices[soldOn] {
			profit = max(tempProfit, profit)
		} else {
			boughtOn = soldOn
		}
		soldOn++
	}

	return profit
}
// @lc code=end

