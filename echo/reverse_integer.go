package echo

const (
	MIN = -2147483648
	MAX = 2147483647
)

// ReverseInteger
//
// Given a 32-bit signed integer, reverse digits of an integer.
//
// Example 1:
//  Input: 123
//  Output: 321
// Example 2:
//  Input: -123
//  Output: -321
// Example 3:
//  Input: 120
//  Output: 21
// Note:
//
// Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
//
// https://leetcode.com/problems/reverse-integer/description/
func ReverseInteger(x int) int {
	var answer = 0
	for x != 0 {
		answer = answer*10 + x%10
		x /= 10
		if answer > MAX || answer < MIN {
			return 0
		}
	}
	return answer
}
