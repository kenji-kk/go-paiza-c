package main
import f "fmt"

func main(){
    var n int 
    f.Scan(&n)
    for i := 0; i < n; i++ {
        var i int
        f.Scan(&i)
        f.Println(i)
    }
}
