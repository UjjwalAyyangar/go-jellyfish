package util

import(
    "math"
    "strings"
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

func HasPrefix(s string, arr ...string) bool {
    for i:=0; i<len(arr); i++{
        if strings.HasPrefix(s, arr[i]){
            return true
        }
    }
    return false
}

func HasSuffix(s string, arr ...string) bool {
    for i:=0; i<len(arr); i++ {
        if strings.HasSuffix(s, arr[i]) {
           return true
        }
    }
    return false
}

func Abs(v int) int {
    if v <0 {
        return -1*v
    } else {
        return v
    }
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j:= 0,len(runes)-1; i<j; i, j = i+1, j-1 {
        runes[i],runes[j] = runes[j],runes[i]
    }
    return string(runes)
}

func Generate_arr(min, max int) []int{
    arr := make([]int, max-min)
    for i:= range arr {
        arr[i] = min + i
    }
    return arr
}

