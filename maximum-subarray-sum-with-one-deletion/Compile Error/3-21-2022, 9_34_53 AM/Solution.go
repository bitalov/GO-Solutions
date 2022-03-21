// https://leetcode.com/problems/maximum-subarray-sum-with-one-deletion

func maximumSum(arr []int) int {
    func maximumSum(A []int) int {
	n := len(A)
	maxRightIs := make([]int, n)
	maxLeftIs := make([]int, n)
	maxRightIs[0] = A[0]
	maxLeftIs[n-1] = A[n-1]
	res := A[0]
	for i := 1; i < n; i++ {
		maxRightIs[i] = max(A[i], maxRightIs[i-1]+A[i])
		j := n - i - 1
		maxLeftIs[j] = max(A[j], A[j]+maxLeftIs[j+1])
		res = max(res, maxRightIs[i])
	}
	for i := 1; i < n-1; i++ {
		res = max(res, maxRightIs[i-1]+maxLeftIs[i+1])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
}