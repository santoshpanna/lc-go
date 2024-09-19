/*
* @lc app=leetcode id=1257 lang=golang
*
* [1257] Smallest Common Region
*
* https://leetcode.com/problems/smallest-common-region/description/
*
  - algorithms
  - Medium (66.68%)
  - Likes:    466
  - Dislikes: 40
  - Total Accepted:    29.2K
  - Total Submissions: 43.2K
  - Testcase Example:  '[["Earth","North America","South America"],["North America","United States","Canada"],["United States","New York","Boston"],["Canada","Ontario","Quebec"],["South America","Brazil"]]\n' +
    '"Quebec"\n' +
    '"New York"'

*
* You are given some lists of regions where the first region of each list
* includes all other regions in that list.
*
* Naturally, if a region x contains another region y then x is bigger than y.
* Also, by definition, a region x contains itself.
*
* Given two regions: region1 and region2, return the smallest region that
* contains both of them.
*
* If you are given regions r1, r2, and r3 such that r1 includes r3, it is
* guaranteed there is no r2 such that r2 includes r3.
*
* It is guaranteed the smallest region exists.
*
*
* Example 1:
*
*
* Input:
* regions = [["Earth","North America","South America"],
* ["North America","United States","Canada"],
* ["United States","New York","Boston"],
* ["Canada","Ontario","Quebec"],
* ["South America","Brazil"]],
* region1 = "Quebec",
* region2 = "New York"
* Output: "North America"
*
*
* Example 2:
*
*
* Input: regions = [["Earth", "North America", "South America"],["North
* America", "United States", "Canada"],["United States", "New York",
* "Boston"],["Canada", "Ontario", "Quebec"],["South America", "Brazil"]],
* region1 = "Canada", region2 = "South America"
* Output: "Earth"
*
*
*
* Constraints:
*
*
* 2 <= regions.length <= 10^4
* 2 <= regions[i].length <= 20
* 1 <= regions[i][j].length, region1.length, region2.length <= 20
* region1 != region2
* regions[i][j], region1, and region2 consist of English letters.
*
*
*/
// package main

// import "fmt"

// func main() {
//   regions := [][]string{
//     {"Earth", "North America", "South America"},
//     {"North America", "United States", "Canada"},
//     {"United States", "New York", "Boston"},
//     {"Canada", "Ontario", "Quebec"},
//     {"South America", "Brazil"},
//   }
//   region1 := "Canada"
//   region2 := "South America"

//   result := findSmallestRegion(regions, region1, region2)
//   fmt.Println(result)
// }

// @lc code=start
func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
  regionsLookup := make(map[string]string)

  for i := 0; i < len(regions); i++ {
    for j := 1; j < len(regions[i]); j++ {
      regionsLookup[regions[i][j]] = regions[i][0]
    }
  }

  region1Ancestors := make(map[string]bool)
  for region1 != "" {
    region1Ancestors[region1] = true
    region1 = regionsLookup[region1]
  }

  for region2 != "" {
    if region1Ancestors[region2] {
      return region2
    }
    region2 = regionsLookup[region2]
  }

  return ""
}
// @lc code=end

