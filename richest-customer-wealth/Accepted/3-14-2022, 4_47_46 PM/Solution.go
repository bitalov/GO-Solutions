// https://leetcode.com/problems/richest-customer-wealth

func max(a int,b int) int {
    if(a > b){
      return a
    }else{
      return b
    }
}

func maximumWealth(accounts [][]int) int {
    
    
    ans := 0
    
    for i := 0 ; i < len(accounts) ; i++{
        sum := 0
        for j := 0 ; j < len(accounts[i]) ; j++{
            
            sum = sum + accounts[i][j]
        }
        ans = max(ans , sum)
    }
    
    return ans
}