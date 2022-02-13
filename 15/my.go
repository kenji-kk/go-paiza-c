package main
import f "fmt"
func main(){
    var n int
    f.Scan(&n)
    var counter int
    for i := 0; i < n; i++{
        var m int
        f.Scan(&m)
        if m % 3 == 0{
          counter++ 
        }
    }
    f.Println(counter)
}
