package main
import f "fmt"

func main(){
    var n,m,k int
    f.Scan(&n, &m, &k)
    for i := 0; i < n; i++{
        var counter int
        for t := 0; t < m; t++{
            var ans int
            f.Scan(&ans)
            if ans == k{
                counter++
            }
        }
        f.Println(counter)
        }
    }
