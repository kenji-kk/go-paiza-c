package main
import f "fmt"

func main(){
    var n int
    f.Scan(&n)
    var ans string = "paiza"
    for i:=0; i<n-1; i++{
        ans += " paiza"
    }
    f.Println(ans)
}
