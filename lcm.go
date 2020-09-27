package math

//Function for finding the least common multiple of two number 
func Lcm(first_num int64 ,  second_num int64) int64 {
    var max_num  int64 = 0
    if first_num >= second_num {
        max_num = first_num
    } else {
        max_num = second_num
    }
    common_mult := max_num
    for (common_mult % first_num > 0) || (common_mult % second_num >0) {
        common_mult += max_num;
    }
    return common_mult

}
