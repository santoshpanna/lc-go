/*
 * @lc app=leetcode id=1545 lang=golang
 *
 * [1545] Find Kth Bit in Nth Binary String
 *
 * https://leetcode.com/problems/find-kth-bit-in-nth-binary-string/description/
 *
 * algorithms
 * Medium (59.26%)
 * Likes:    1179
 * Dislikes: 72
 * Total Accepted:    86.6K
 * Total Submissions: 131.1K
 * Testcase Example:  '3\n1'
 *
 * Given two positive integers n and k, the binary string Sn is formed as
 * follows:
 *
 *
 * S1 = "0"
 * Si = Si - 1 + "1" + reverse(invert(Si - 1)) for i > 1
 *
 *
 * Where + denotes the concatenation operation, reverse(x) returns the reversed
 * string x, and invert(x) inverts all the bits in x (0 changes to 1 and 1
 * changes to 0).
 *
 * For example, the first four strings in the above sequence are:
 *
 *
 * S1 = "0"
 * S2 = "011"
 * S3 = "0111001"
 * S4 = "011100110110001"
 *
 *
 * Return the k^th bit in Sn. It is guaranteed that k is valid for the given
 * n.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 3, k = 1
 * Output: "0"
 * Explanation: S3 is "0111001".
 * The 1^st bit is "0".
 *
 *
 * Example 2:
 *
 *
 * Input: n = 4, k = 11
 * Output: "1"
 * Explanation: S4 is "011100110110001".
 * The 11^th bit is "1".
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 20
 * 1 <= k <= 2^n - 1
 *
 *
 */

package main

import "fmt"

func main() {
	testCases := []struct {
		n, k int
		want byte
	}{
		{3, 1, '0'},
		{4, 11, '1'},
	}

	for _, tc := range testCases {
		got := findKthBit(tc.n, tc.k)
		fmt.Println(got, tc.want, got == tc.want)
	}
}

// @lc code=start
func invert(bString []byte) []byte {
	invertedBString := make([]byte, 0)

	for _, ch := range bString {
		if ch == '0' {
			invertedBString = append(invertedBString, '1')
		} else {
			invertedBString = append(invertedBString, '0')
		}
	}

	return invertedBString
}

func reverse(bString []byte) []byte {
	for i, j := 0, len(bString) - 1; i < j; i, j = i+1, j-1 {
		bString[i], bString[j] = bString[j], bString[i]
	}

	return bString
}

func findKthBit(n int, k int) byte {
	bString := make([]byte, 0)

	for i := 0; i <= n; i++ {
		lenOfBString := len(bString)
		if (k-1) < lenOfBString {
			return bString[k-1]
		}

		if i == 0 {
			bString = append(bString, '0')
			continue
		}

		bString = append(bString, '1')

		partBString := bString[:len(bString)-1]

		bString = append(bString, reverse(invert(partBString))...)
	}

	return '0'
}
// @lc code=end

