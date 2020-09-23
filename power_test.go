package math

import "testing"

func Test_power(t *testing.T) {

	type p struct {
		num float64
		pow int
	}

	sample := map[float64]p{
		1:   p{1, 1},
		2:   p{2, 1},
		4:   p{2, 2},
		9:   p{3, 2},
		125: p{5, 3},
		0.5: p{2, -1},
		0.25: p{2, -2},

	}

	for i, v := range sample {
		if i != Pow(v.num, v.pow) {
			t.Errorf("%v raised to power of %v is %v", v.num, v.pow, Pow(v.num, v.pow))
		}
	}

	if 1 != Pow(322, 0) {
		t.Errorf("322 raised to power 0 is not 1")
	}

}
