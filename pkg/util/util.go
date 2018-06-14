package util

import(
    "math"
)

func Memset(arr []int, v int) {
    if len(arr) == 0 {
        return
    }
    arr[0] = v
    for i:=1; i< len(arr); i= i*2{
        copy(arr[i:], arr[:i])
    }
}

func Min(arr ...int) int {
    mini := math.MaxInt64
    for _,n := range arr{
        if n<mini {
            mini = n
        }
    }
    return mini
}

func Max(arr ...int) int {
    maxi := math.MinInt64
    for _,n := range arr{
        if n>maxi {
            maxi = n
        }
    }
    return maxi
}



func Generate_arr(min, max int) []int{
    arr := make([]int, max-min)
    for i:= range arr {
        arr[i] = min + i
    }
    return arr
}