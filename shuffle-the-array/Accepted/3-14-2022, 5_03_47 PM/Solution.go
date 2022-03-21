// https://leetcode.com/problems/shuffle-the-array

func shuffle(nums []int, n int) []int {
    
    ans := make([]int,0)
    
    for i := 0 ; i < len(nums) / 2 ; i++{
        
        ans = append(ans , nums[i])
        ans = append(ans , nums[i + len(nums)/2])
    }
    return ans
}