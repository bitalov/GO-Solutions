// https://leetcode.com/problems/fibonacci-number

func fib(n int) int {
    
    f := make([]int,2)
    
    f[0] = 0
    f[1] = 1
    
    for i := 2 ; i <= n ; i++ {
        f = append(f,f[i - 1] + f[i - 2])
    }
    
    return f[n]
    
}