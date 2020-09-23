package math 
import "testing"

func Test_square_root(t *testing.T) {
    sample := map[float64]float64 {
    0: 0.0,
    1: 1.0,
    4:2.0 ,
    9: 3.0,
        16.0: 4.0,

    }
    for i , v := range sample {
        if v != Sqrt(i) {
            t.Errorf("Prime of %v is %v" , i , Sqrt(i))

        }
    }

}
