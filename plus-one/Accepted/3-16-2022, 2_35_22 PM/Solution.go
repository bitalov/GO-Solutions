// https://leetcode.com/problems/plus-one

func plusOne(digits []int) []int {
    sz := len(digits)
	if sz == 0 {
		return []int{1}
	}

	digits[sz-1]++

	for i := sz - 1; i > 0; i-- {
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10
		digits[i-1]++
	}

	if digits[0] > 9 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}

	return digits
}