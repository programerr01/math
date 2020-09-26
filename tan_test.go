package math
import "testing" 

func Test_tan(t *testing.T){

    sample := map[float64]float64{
    1 : 45,
    0:0,
    0.36: 20,
    0.57: 30,
    

    }
    for i , v := range sample {
        if i != Tan(v){
        t.Errorf("Tan of %v is %v" ,v , Tan(v))
        }
    }

}
