package math 

// Function of Square root to calculate the square root of the function 
func Sqrt(val float64) float64 {
    z := 1.0
    if val == 0 {
        return 0.0
    }
    for i :=1 ;i <20;i++ {
        z -= (z *z -float64(val) ) /(2*z)
    }
    return z
}


