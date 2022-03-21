// https://leetcode.com/problems/first-unique-character-in-a-string

func firstUniqChar(s string) int {
    rec := make([]int, 26)
	for _, b := range s {
		rec[b-'a']++
	}

	for i, b := range s {
		if rec[b-'a'] == 1 {
			return i
		}
	}

	return -1
}