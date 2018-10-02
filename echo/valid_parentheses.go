package echo

// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
// 1. Open brackets must be closed by the same type of brackets.
//
// 2. Open brackets must be closed in the correct order.
//
// Note that an empty string is also considered valid.
// Example 1:
//  Input: "()"
//  Output: true
// Example 2:
//  Input: "()[]{}"
//  Output: true
// Example 3:
//  Input: "(]"
//  Output: false
// Example 4:
//  Input: "([)]"
//  Output: false
// Example 5:
//  Input: "{[]}"
//  Output: true
func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	left := map[byte]int{
		'(': 0,
		'[': 0,
		'{': 0,
	}
	right := map[byte]int{
		')': 0,
		']': 0,
		'}': 0,
	}
	bytes := []byte(s)
	stack := []byte{}
	for i := range bytes {
		if _, ok := left[bytes[i]]; ok {
			stack = append(stack, bytes[i])
		} else if _, ok := right[bytes[i]]; ok {
			if len(stack) == 0 {
				return false
			}
			switch stack[len(stack)-1] {
			case '(':
				if bytes[i] != ')' {
					return false
				}
			case '[':
				if bytes[i] != ']' {
					return false
				}
			case '{':
				if bytes[i] != '}' {
					return false
				}
			}
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
