// https://leetcode.com/problems/subarray-sum-equals-k

func subarraySum(nums []int, k int) int {
    
    ret, sum := 0 , 0
    mp := make(map[int]int, len(nums))
    
    mp[0] = 1
    
    for i := range nums {
        sum += nums[i]
        
        ret += mp[sum - k]
        
        mp[sum]++;
        
    }
    
    
    return ret
    
    
}