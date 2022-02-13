package main
import f "fmt"

func main(){
    var n int
    f.Scan(&n)
    for i := 0; i < n; i++ {
        var (
            name string
            age int
            )
        f.Scan(&name, &age)
        age++
        f.Println(name,age)
    }
}
