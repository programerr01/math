package math

func Pow(num float64, p int) float64 {
	var result float64 = 1
	if p == 0 {
		return 1
	} else if p > 0 {
		for i := 0; i < p; i++ {
			result *= num
		}

		return result
	}else {
		for i:= 0; i > p; i-- {
			result *= num
		}

		return 1/result
	}
}


