// https://leetcode.com/problems/number-of-good-pairs

func numIdenticalPairs(nums []int) int {
    
    ans := 0
    for i:= 0 ; i < len(nums) ; i++{
        for j:= i + 1 ; j < len(nums) ; j++{
          
            if nums[i] == nums[j]{
                ans++
            }
        }
    }
    return ans
}