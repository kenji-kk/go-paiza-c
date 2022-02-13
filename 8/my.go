package main
import f "fmt"

func main(){
    var n int
    f.Scan(&n)
    for i := 0; i < n; i++ {
        var i string
        f.Scan(&i)
        f.Println(len([]rune(i)))
    }
}
