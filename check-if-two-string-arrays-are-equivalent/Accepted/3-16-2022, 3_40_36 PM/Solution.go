// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    
    var a,b strings.Builder
    for i := 0 ; i < len(word1) ; i++ {
            
            a.WriteString(word1[i])
        
        
    }
    
     for i := 0 ; i < len(word2) ; i++ {
        
         b.WriteString(word2[i])
        
    }
    
    if a.String() == b.String() {
        return true
    } else {
        return false
    }
    
    
}