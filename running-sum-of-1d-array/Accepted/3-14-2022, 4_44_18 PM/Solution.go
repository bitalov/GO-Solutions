// https://leetcode.com/problems/running-sum-of-1d-array

func runningSum(nums []int) []int {
    
    ans := make([]int,0)
    last := 0
    for i := 0 ; i < len(nums) ; i++{
        ans = append(ans , nums[i] + last)
        last = last + nums[i]
    }
    
    return ans
}