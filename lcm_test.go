package math 

import "testing"

type int_pair struct {
    int1 int64 
    int2 int64
}


func Test_lcm(t *testing.T){
    
    sample := map[int64]int_pair{
        20: int_pair{10,20},
        228 : int_pair{12,76},
        

    }
    for i , v := range sample {
        if i != Lcm(v.int1 , v.int2){
        t.Errorf("the lcm of %v and %v is %v" ,v.int1, v.int2 , Lcm(v.int1 , v.int2))
    }
    }

}
