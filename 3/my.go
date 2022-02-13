package main
import f "fmt"

func main(){
    var n int
    f.Scan(&n)
    for i:=0; i<n; i++{
        var m int
        f.Scan(&m)
        f.Println(m)
    }
}
