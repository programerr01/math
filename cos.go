package math 

import (
    "fmt"
    "strconv"
    
)

//Function for Calaculating cos of an angle 
//NOTE: Argument should be in degrees
func Cos(num float64)float64 {
    //num = num *(pi/180)
    sum := Sin(90 -num)

    result , _ := strconv.ParseFloat(fmt.Sprintf("%.2f" , sum), 64)
    return result 

}
