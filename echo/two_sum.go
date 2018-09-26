package echo

// TwoSum
//
// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// Example:
//  Given nums = [2, 7, 11, 15], target = 9,
//  Because nums[0] + nums[1] = 2 + 7 = 9,
//  return [0, 1].
// https://leetcode.com/problems/add-two-numbers/description/
func TwoSum(nums []int, target int) []int {
	var answer []int
	answer_map := make(map[int]int)
	var length int = len(nums)

	for i := 0; i < length; i++ {
		tmp, ok := answer_map[nums[i]]
		if ok {
			answer := []int{tmp, i}
			return answer
		}
		answer_map[target-nums[i]] = i
	}
	return answer
}
