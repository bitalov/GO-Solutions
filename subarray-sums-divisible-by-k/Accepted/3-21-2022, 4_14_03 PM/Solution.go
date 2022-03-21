// https://leetcode.com/problems/subarray-sums-divisible-by-k

func subarraysDivByK(A []int, K int) int {
    count := [10000]int{}
	count[0] = 1
	prefix, res := 0, 0
	for _, a := range A {
		prefix = (prefix + a%K + K) % K
		res += count[prefix] 
		count[prefix]++
	}
	return res
}