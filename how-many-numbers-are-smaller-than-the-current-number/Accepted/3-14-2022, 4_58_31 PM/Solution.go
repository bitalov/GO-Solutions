// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number

func smallerNumbersThanCurrent(nums []int) []int {
    
    ans := make([]int,0)
    
    for i := 0 ; i < len(nums) ; i++{
        cnt := 0
        for j := 0 ; j < len(nums) ; j++{
            if(nums[j] < nums[i]){
                cnt++
            }
        }
        
        ans = append(ans,cnt)
    }
    
    return ans
}