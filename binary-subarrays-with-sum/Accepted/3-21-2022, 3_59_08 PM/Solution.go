// https://leetcode.com/problems/binary-subarrays-with-sum

func numSubarraysWithSum(A []int, S int) int {
    preSum, res := 0, 0
	count := make([]int, len(A)+1)
	count[0] = 1
	for _, n := range A {
		preSum += n
		if preSum >= S {
			res += count[preSum-S]
		}
		count[preSum]++
	}
	return res
}