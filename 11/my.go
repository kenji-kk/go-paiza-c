package main
import f "fmt"

func main(){
    var s string
    f.Scan(&s)
    for len([]rune(s)) < 3 {
        s = "0" + s
    }
    f.Println(s)
}
