package math 

import "testing" 

type p struct {
    num uint 
    base uint 
}

func Test_log(t *testing.T){
    sample := map[uint]p{
        1: p{2,2},
        2:p{4,2},
        3:p{27,3},
        4:p{256,4},
        5:p{100000,10},

    }

    for i , v := range sample {
    if i != Log(v.num , v.base){
        t.Errorf("%v is the log  of %v in base %v" ,Log(v.num, v.base),v.num, v.base)
    }
    }
}
