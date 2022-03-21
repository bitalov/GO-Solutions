// https://leetcode.com/problems/concatenation-of-array

func getConcatenation(nums []int) []int {
    
    a := make([]int,0)
    
    for i:= 0 ; i < len(nums) ; i++ {
        a = append(a,nums[i])
    }
    
    for i:= 0 ; i < len(nums) ; i++ {
        a = append(a,nums[i])
    }
    
    return a
}