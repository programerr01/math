package math 
import (
    "fmt"
    "strconv"
)

//Function for Calculating sin of an angle 
//NOTE:- Argument should in degrees 
func Sin(num float64) float64 {
    num = num * (pi/180)
    i := 0 
    j := 0
    sum_sin  := 0.0  
    for ; i <20;i= i+2 {
        sum_sin += float64(Pow(-1,j)* Pow(num , i+1))/float64(Factorial(i+1))
        j = j+1

    }
    result ,_ := strconv.ParseFloat(fmt.Sprintf("%.2f" , sum_sin),64)
    return result

}
