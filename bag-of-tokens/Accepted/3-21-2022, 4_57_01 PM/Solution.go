// https://leetcode.com/problems/bag-of-tokens

func bagOfTokensScore(tokens []int, P int) int {
    sort.Ints(tokens)

	maxScore, score := 0, 0
	i, j := 0, len(tokens)-1

	for i <= j {
		if P >= tokens[i] {
			P -= tokens[i]
			score++
			maxScore = score
i++
		} else if score > 0 {
			P += tokens[j]
			score--
			j--
		} else {
			break
		}
	}

	return maxScore
}