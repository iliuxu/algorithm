package echo

import (
	"strings"
)

// Longest Common Prefix
//
// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//
// Example 1:
//   Input: ["flower","flow","flight"]
//   Output: "fl"
// Example 2:
//   Input: ["dog","racecar","car"]
//   Output: ""
// Explanation: There is no common prefix among the input strings.
// Note:
//
// All given inputs are in lowercase letters a-z.
//
// https://leetcode.com/problems/longest-common-prefix/description/
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefix string = strs[0]
	for i := 0; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			t := len(prefix) - 1
			prefix = prefix[0:t]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}
