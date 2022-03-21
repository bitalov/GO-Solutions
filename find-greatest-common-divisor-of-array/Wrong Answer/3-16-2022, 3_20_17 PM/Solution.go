// https://leetcode.com/problems/find-greatest-common-divisor-of-array

func gcd(a int,b int) int {

    if(b == 0) {
        return a
    } else {
    
    return gcd(b , a % b)
    }
}


func findGCD(nums []int) int {
    
    ans := nums[0]
    for i := 1 ; i < len(nums) ; i++ {
        
        ans = gcd(ans , nums[i])
    }
    
    return ans
    
}