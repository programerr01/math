package math

import "testing"

func Test_Abs(t *testing.T) {
	sample := map[float64]float64{11: 11, 1: -1, 0: 0, 2: -2, 3: -3, 5: -5, 121: -121, 343: 343, 342: 342, 1039232: 1039232, 324322: -324322,
		0.2:   -0.2,
		0.023: -0.023,
		0.022: 0.022,
		323.1: 323.1,
		323.2: -323.2,
	}
	for i, v := range sample {
		if i != Abs(v) {
			t.Errorf("Absolute of %v is %v", v, Abs(v))
		}
	}
}

func Test_Max(t *testing.T) {
	sample := map[float64][]float64{
		1:   []float64{0, 1, -1},
		34:  []float64{34, 3, 4, -3, 2},
		-32: []float64{-323, -33, -32, -2000},
		0.1: []float64{0.001, 0.1, 0.012, 0.0001},
	}

	for i, v := range sample {
		if i != Max(v) {
			t.Errorf("Maximum of %v is %v", v, Max(v))
		}
	}
}

func Test_Min(t *testing.T) {
	sample := map[float64][]float64{
		-1:     []float64{0, 1, -1},
		-3:     []float64{34, 3, 4, -3, 2},
		-2000:  []float64{-323, -33, -32, -2000},
		0.0001: []float64{0.001, 0.1, 0.012, 0.0001},
	}

	for i, v := range sample {
		if i != Min(v) {
			t.Errorf("Maximum of %v is %v", v, Min(v))
		}
	}
}

func Test_Factorial(t *testing.T) {
	sample := map[int]int{
		0: 1,
		1: 1,
		2: 2,
		3: 6,
		4: 24,
		5: 120,
		6: 720,
	}
	for i, val := range sample {
		if val != Factorial(i) {
			t.Errorf("Factorial of %v is %v", i, Factorial(i))
		}
	}
}
