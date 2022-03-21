// https://leetcode.com/problems/count-good-triplets

func abs(a int) int {
    
    if a < 0 {
        return -a
    }else{
        return a
    }
}


func countGoodTriplets(arr []int, a int, b int, c int) int {
    
    n := len(arr)
    ans := 0
    for i := 0 ; i < n ; i++{
        
        for j := i + 1 ; j < n ; j++ {
            
            for k := j + 1 ; k < n ; k++{
                
                pa := abs(arr[i] - arr[j])
                pb := abs(arr[j] - arr[k])
                pc := abs(arr[i] - arr[k])
                
                if(pa <= a && pb <= b && pc <= c){
                    
                    ans++
                }
            }
        }
    }
    
    return ans
    
}