package main
import f "fmt"

func main(){
    var n int
    f.Scan(&n)
    flag := false
    for i := 0; i < n; i++{
        var m int
        f.Scan(&m)
        if m == 7{
            flag = true
            break
        }
    }
    if flag{
        f.Println("YES")
    } else{
        f.Println("NO")
    }
}
