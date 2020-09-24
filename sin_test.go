package math 

import "testing" 

func Test_sin(t *testing.T) {


    sample := map[float64]float64 {
    0.02: 1,
    0.03:2,
    1:90,
    0:0,
    

    }

    for i , v := range sample {
        if i != Sin(v) {
        t.Errorf("sin of %v is %v" , v , Sin(v))
        }
    }
}


