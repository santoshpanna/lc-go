/*
 * @lc app=leetcode id=929 lang=golang
 *
 * [929] Unique Email Addresses
 *
 * https://leetcode.com/problems/unique-email-addresses/description/
 *
 * algorithms
 * Easy (67.30%)
 * Likes:    2700
 * Dislikes: 345
 * Total Accepted:    504.6K
 * Total Submissions: 749.6K
 * Testcase Example:  '["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]'
 *
 * Every valid email consists of a local name and a domain name, separated by
 * the '@' sign. Besides lowercase letters, the email may contain one or more
 * '.' or '+'.
 *
 *
 * For example, in "alice@leetcode.com", "alice" is the local name, and
 * "leetcode.com" is the domain name.
 *
 *
 * If you add periods '.' between some characters in the local name part of an
 * email address, mail sent there will be forwarded to the same address without
 * dots in the local name. Note that this rule does not apply to domain
 * names.
 *
 *
 * For example, "alice.z@leetcode.com" and "alicez@leetcode.com" forward to the
 * same email address.
 *
 *
 * If you add a plus '+' in the local name, everything after the first plus
 * sign will be ignored. This allows certain emails to be filtered. Note that
 * this rule does not apply to domain names.
 *
 *
 * For example, "m.y+name@email.com" will be forwarded to "my@email.com".
 *
 *
 * It is possible to use both of these rules at the same time.
 *
 * Given an array of strings emails where we send one email to each emails[i],
 * return the number of different addresses that actually receive mails.
 *
 *
 * Example 1:
 *
 *
 * Input: emails =
 * ["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]
 * Output: 2
 * Explanation: "testemail@leetcode.com" and "testemail@lee.tcode.com" actually
 * receive mails.
 *
 *
 * Example 2:
 *
 *
 * Input: emails = ["a@leetcode.com","b@leetcode.com","c@leetcode.com"]
 * Output: 3
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= emails.length <= 100
 * 1 <= emails[i].length <= 100
 * emails[i] consist of lowercase English letters, '+', '.' and '@'.
 * Each emails[i] contains exactly one '@' character.
 * All local and domain names are non-empty.
 * Local names do not start with a '+' character.
 * Domain names end with the ".com" suffix.
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
		emails				[]string
		output				int
	} {
		{emails: []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}, output: 2},
		{emails: []string{"a@leetcode.com","b@leetcode.com","c@leetcode.com"}, output: 3},
	}

	for _, tc := range testCases {
		fmt.Println("emails =", tc.emails, " output =", numUniqueEmails(tc.emails), " expected =", tc.output)
	}
}

// @lc code=start
func getParsedEmail(email string) string {
	emailAddress := strings.Split(email, "@")

	emailAddress[0] = strings.Replace(emailAddress[0], ".", "", -1)

	indexOf := strings.Index(emailAddress[0], "+")

	if indexOf >= 0 {
		return emailAddress[0][:indexOf] + "@" + emailAddress[1]
	}

	return emailAddress[0] + "@" + emailAddress[1]
}

func numUniqueEmails(emails []string) int {
	emailMap := make(map[string]bool)
	
	for i := 0; i < len(emails); i++ {
		email := getParsedEmail(emails[i])

		if _, exists := emailMap[email]; !exists {
			emailMap[email] = true
		}
	}

	count := len(emailMap)

	return count
}
// @lc code=end

