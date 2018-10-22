package echo

// Search Insert Position
//
// Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//
// You may assume no duplicates in the array.
//
// Example 1:
//
//  Input: [1,3,5,6], 5
//  Output: 2
// Example 2:
//
//  Input: [1,3,5,6], 2
//  Output: 1
// Example 3:
//
//  Input: [1,3,5,6], 7
//  Output: 4
// Example 4:
//
//  Input: [1,3,5,6], 0
//  Output: 0
// https://leetcode.com/problems/search-insert-position/description/
func SearchInsert(nums []int, target int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if target <= nums[j] {
			i = j
			return i
		}
		i = j + 1
	}
	return i
}
