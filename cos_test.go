package math 

import "testing" 

func Test_cos(t *testing.T){
    
    sample:= map[float64]float64{
    1:0,
    0.71: 45,
    0: 90,
    -1.0: 180,
    -0.24: 256,
    }
    for i , v:= range sample {
    if i != Cos(v){
        t.Errorf("cos of %v is %v" ,v , Cos(v))
    }
    }

}
