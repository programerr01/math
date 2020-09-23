package math 
import "testing"

func Test_prime(t *testing.T) {
    sample := map[uint]bool {
        0:false,
        1:false ,
        2:true,
        3:true,
        4:false,
        5:true,
        6:false,
        7:true,
        28:false,
        31:true,


    }
    for i , v := range sample {
        if v != Prime(i) {
            t.Errorf("Prime of %v is %v" ,i , Prime(i))
        }

    } 
}
