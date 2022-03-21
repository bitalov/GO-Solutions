// https://leetcode.com/problems/length-of-last-word

func lengthOfLastWord(s string) int {
    sz := len(s)
	if sz == 0 {
		return 0
	}
	res := 0
	for i := sz - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if res != 0 {
				return res
			}
			continue
		}
		res++
	}

	return res
}