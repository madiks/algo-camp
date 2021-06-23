package week01

// https://leetcode-cn.com/problems/plus-one/
// https://leetcode-cn.com/submissions/detail/189146423/

func plusOne(digits []int) []int {
	carry := true
	for p := len(digits) - 1; p >= 0; p-- {
		if !carry {
			break
		}
		if digits[p] == 9 {
			digits[p] = 0
			carry = true
		} else {
			digits[p]++
			carry = false
		}
	}
	if carry {
		digits = append([]int{1}, digits...)
	}
	return digits
}
