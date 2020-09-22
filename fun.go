package math



/// Abs function takes any integer value and returns in absolute value.
func Abs(i int) uint {
	if i < 0 {
		return uint( i * -1 )
	}else {
		return uint(i) 
	}
}



// Max function takes any set of integers and return the maximum value.
func Max(arr []int) int {
	max := arr[0];

	for i:=1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	return max 
}


// Min function takes any set of integers and return the minimu value.
func Min(arr []int) int {
	min := arr[0];

	for i:=1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}

	return min 
}