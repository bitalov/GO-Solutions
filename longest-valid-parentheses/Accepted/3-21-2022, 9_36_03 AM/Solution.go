// https://leetcode.com/problems/longest-valid-parentheses


func longestValidParentheses(s string) int {
	var left, max, temp int
	record := make([]int, len(s))

	for i, b := range s {
		if b == '(' {
			left++
		} else if left > 0 {
			left--
			record[i] = 2
		}
	}

	for i := 0; i < len(record); i++ {
		if record[i] == 2 {
			j := i - 1
			for record[j] != 0 {
				j--
			}
			record[i], record[j] = 1, 1
		}
	}

	for _, r := range record {
		if r == 0 {
			temp = 0
			continue
		}

		temp++
		if temp > max {
			max = temp
		}
	}

	return max
}