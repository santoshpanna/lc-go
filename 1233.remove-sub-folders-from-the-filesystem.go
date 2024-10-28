/*
 * @lc app=leetcode id=1233 lang=golang
 *
 * [1233] Remove Sub-Folders from the Filesystem
 *
 * https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/description/
 *
 * algorithms
 * Medium (66.09%)
 * Likes:    883
 * Dislikes: 129
 * Total Accepted:    67.5K
 * Total Submissions: 98.5K
 * Testcase Example:  '["/a","/a/b","/c/d","/c/d/e","/c/f"]'
 *
 * Given a list of folders folder, return the folders after removing all
 * sub-folders in those folders. You may return the answer in any order.
 * 
 * If a folder[i] is located within another folder[j], it is called a
 * sub-folder of it. A sub-folder of folder[j] must start with folder[j],
 * followed by a "/". For example, "/a/b" is a sub-folder of "/a", but "/b" is
 * not a sub-folder of "/a/b/c".
 * 
 * The format of a path is one or more concatenated strings of the form: '/'
 * followed by one or more lowercase English letters.
 * 
 * 
 * For example, "/leetcode" and "/leetcode/problems" are valid paths while an
 * empty string and "/" are not.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: folder = ["/a","/a/b","/c/d","/c/d/e","/c/f"]
 * Output: ["/a","/c/d","/c/f"]
 * Explanation: Folders "/a/b" is a subfolder of "/a" and "/c/d/e" is inside of
 * folder "/c/d" in our filesystem.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: folder = ["/a","/a/b/c","/a/b/d"]
 * Output: ["/a"]
 * Explanation: Folders "/a/b/c" and "/a/b/d" will be removed because they are
 * subfolders of "/a".
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: folder = ["/a/b/c","/a/b/ca","/a/b/d"]
 * Output: ["/a/b/c","/a/b/ca","/a/b/d"]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= folder.length <= 4 * 10^4
 * 2 <= folder[i].length <= 100
 * folder[i] contains only lowercase letters and '/'.
 * folder[i] always starts with the character '/'.
 * Each folder name is unique.
 * 
 * 
 */

// @lc code=start
type TrieNode struct {
	children map[string]*TrieNode
	isFolder bool
}

func removeSubfolders(folder []string) []string {
	root := &TrieNode{children: make(map[string]*TrieNode), isFolder: false}
	for _, f := range folder {
		node := root
		for _, c := range strings.Split(f, "/")[1:] {
			if _, ok := node.children[c]; !ok {
				node.children[c] = &TrieNode{children: make(map[string]*TrieNode), isFolder: false}
			}
			node = node.children[c]
		}
		node.isFolder = true
	}
	result := []string{}
	var dfs func(node *TrieNode, path string)
	dfs = func(node *TrieNode, path string) {
		if node.isFolder {
			result = append(result, path)
			return
		}
		for c, n := range node.children {
			dfs(n, path+"/"+c)
		}
	}
	dfs(root, "")
	return result
}
// @lc code=end

