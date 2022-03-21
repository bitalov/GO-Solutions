// https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer

func prod(n int) int {
    
    ret := 1
    
    for n > 0 {
        ret = ret * (n % 10)
        n/=10
    }
    
    return ret
}

func sum(n int) int {
    
    ret := 0
    
    for(n > 0){
        
        ret = ret + (n % 10)
        n /= 10
    }
    
    return ret
    
}


func subtractProductAndSum(n int) int {
    
    return prod(n) - sum(n)
}