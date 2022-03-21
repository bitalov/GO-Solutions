// https://leetcode.com/problems/sum-of-all-odd-length-subarrays

func sumOddLengthSubarrays(arr []int) int {
    
    ret := 0
    for l := 0 ; l < len(arr) ; l++{
        cnt := 0
        for r := l ; r < len(arr) ; r++{
            
            sz := (r - l + 1)
            cnt = cnt + arr[r]
            if(sz % 2 == 1){
              ret+=cnt
            }
        }
    }
    
    return ret
    
}