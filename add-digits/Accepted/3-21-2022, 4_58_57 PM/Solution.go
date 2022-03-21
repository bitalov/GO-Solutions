// https://leetcode.com/problems/add-digits

func addDigits(num int) int {
    return (num-1)%9 + 1
}