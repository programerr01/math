package math 

//Function to find log(n,r) on any arbitary base r 
//Note: It not returns the int value 
func Log(n uint , r uint) uint {
    if n < r {
        return 0
    }
    
    return 1+ Log(n/r ,r)

}
