package math 
import (
    "fmt"
    "strconv"
)

//Function for Calculating tan of an angle 
//NOTE:- Argument should be in degree
func Tan(num float64)float64{
    sum := Sin(num) / Cos(num)

    result , _ := strconv.ParseFloat(fmt.Sprintf("%.2f" , sum),64)
    return result 

}
