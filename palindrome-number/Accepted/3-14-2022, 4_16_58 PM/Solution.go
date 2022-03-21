// https://leetcode.com/problems/palindrome-number

func isPalindrome(x int) bool {
    
    y := x
    z := 0
    
    for x > 0 {
        
        dig := x % 10
        x = x / 10
        z = z * 10 + dig
    }
    
    
    return z == y
}