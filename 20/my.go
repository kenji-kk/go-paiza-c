package main
import (
    f "fmt"
    "sort"
    )

func main(){
    var n int
    f.Scan(&n)
    slice := make([]int, n)
    for i := range slice{
        f.Scan(&slice[i])
    }
    sort.Sort(sort.IntSlice(slice))
    for _, v := range slice{
        f.Println(v)
    }
}
