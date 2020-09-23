package math 


//Prime Function takes an int and return true or false based on the number is prime or not 
func Prime(a uint) bool {
    if a == 1 || a == 0 {
        return false
    }
    for i := a-1 ; i >= 2; i-- {
        if a % i == 0 {
            return false
        } else {
            continue
        }

    } 

    
    return true
}
