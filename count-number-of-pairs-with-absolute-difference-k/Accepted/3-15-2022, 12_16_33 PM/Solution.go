// https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k

func countKDifference(nums []int, k int) int {
    
    ans:= 0
    
    for i := 0 ; i < len(nums) ; i++{
        for j := i + 1 ; j < len(nums) ; j++{
            
            diff := (nums[i] - nums[j])
            
            if diff < 0{
                diff = diff * -1
            }
            
            if diff == k {
                ans++
            }
        }
    }
    
    return ans
}