// https://leetcode.com/problems/check-if-all-as-appears-before-all-bs

func checkString(s string) bool {
    
    b := 0
    
    for c := range s {
        
        if s[c] == 'a' && b == 1 {
             
            return false
        }
        
        if(s[c] == 'b') {
            b = 1
        }
        
    }
    
    return true
}