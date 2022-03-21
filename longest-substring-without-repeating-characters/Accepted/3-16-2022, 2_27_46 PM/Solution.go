// https://leetcode.com/problems/longest-substring-without-repeating-characters

func lengthOfLongestSubstring(s string) int {
    
   
    pos := [512]int{}
     ans := 0
    last := 0
    n := len(s)
    for i := 0 ; i < 512 ; i++ {
    
     pos[i] = -1
    }  
    
    for i := 0 ; i < n ; i++{
        
        if(pos[s[i]] >= last) {
            last = pos[s[i]] + 1
        } else if i + 1 - last > ans {
            ans = i + 1 - last
        }
        pos[s[i]] = i
    }
   
    
    return ans
    
}