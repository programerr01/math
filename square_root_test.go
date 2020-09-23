package math 
import "testing"

func Test_square_root(t *testing.T) {
    sample := map[int]float64 {
    0: 0.0,
    1: 1.0,
    4:2.0 ,
    9: 3.0,

    }
    for i , v := range sample {
        if v != Square_root(i) {
            t.Errorf("Prime of %v is %v" , i , Square_root(i))

        }
    }

}
