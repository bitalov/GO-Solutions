// https://leetcode.com/problems/find-greatest-common-divisor-of-array

func gcd(a int,b int) int {

    if(b == 0) {
        return a
    } else {
    
    return gcd(b , a % b)
    }
}





func findGCD(nums []int) int {
    
    n := len(nums)
    sort.Ints(nums)
    return gcd(nums[0],nums[n - 1])
    
}