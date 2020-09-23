package math

/// Abs function takes any integer value and returns in absolute value.
func Abs(i int) uint {
	if i < 0 {
		return uint(i * -1)
	} else {
		return uint(i)
	}
}

// Max function takes any set of number and return the maximum value.
func Max(arr []float64) float64 {
	max := arr[0]

	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	return max
}

// Min function takes any set of number and return the minimu value.
func Min(arr []float64) float64 {
	min := arr[0]

	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}

	return min
}

// Factorial function takes a int and return the factorial using Recursion
func Factorial(i int) int {
	if i == 0 || i == 1 || i < 0 {
		return 1
	} else {
		return i * Factorial(i-1)
	}

}
