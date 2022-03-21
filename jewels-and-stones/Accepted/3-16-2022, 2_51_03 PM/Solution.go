// https://leetcode.com/problems/jewels-and-stones

func numJewelsInStones(jewels string, stones string) int {
    
    good := make(map[byte]bool)
    
    for i := 0 ; i < len(jewels) ; i++ {
         
        good[jewels[i]] = true
    }
    
    ret := 0
    for i := 0 ; i < len(stones) ; i++ {
        if(good[stones[i]]) {
            
            ret++
        }
    }
    
    return ret
    
}