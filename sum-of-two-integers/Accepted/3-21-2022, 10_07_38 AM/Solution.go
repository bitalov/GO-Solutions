// https://leetcode.com/problems/sum-of-two-integers

func getSum(a int, b int) int {
	for a != 0 {
		a, b = (a&b)<<1, a^b
	}
	return b
}