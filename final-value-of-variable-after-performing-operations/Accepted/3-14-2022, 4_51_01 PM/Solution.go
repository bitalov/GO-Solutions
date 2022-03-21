// https://leetcode.com/problems/final-value-of-variable-after-performing-operations

func finalValueAfterOperations(operations []string) int {
    
    
    ans := 0
    for j := 0 ; j < len(operations) ; j++{
        
        if(operations[j] == "X++" || operations[j] == "++X"){
            ans++
        }else{
            ans--
        }
    }
    
    return ans
}