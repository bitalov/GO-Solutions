// https://leetcode.com/problems/smallest-index-with-equal-value

func smallestEqual(nums []int) int {
    
    n := len(nums)
    
    for i := 0 ; i < n ; i++ {
        
        val := i % 10
        
        if(val == nums[i]) {
            return i
        }
    }
    
    return -1
    
}