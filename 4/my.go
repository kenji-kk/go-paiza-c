package main
import f "fmt"

func main(){
    var n int
    f.Scan(&n)
    var t int = 1
    for i:=0; i<n; i++{
      var i int
      f.Scan(&i)
      if t < i {
        t = i
      }
    }
    f.Println(t) 
}
