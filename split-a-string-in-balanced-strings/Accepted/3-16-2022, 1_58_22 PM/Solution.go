// https://leetcode.com/problems/split-a-string-in-balanced-strings

func balancedStringSplit(s string) int {
    
       ret := 0
       cnt := 0
            
          for _, c := range s{
              
              if(c == 'R') {
                  cnt++
              } else {
                  cnt--
              }
              
              if (cnt == 0){
                  ret++
              }
          }  
        
        return ret
    
}