package main
import f "fmt"

func main(){
    var n, last int
    f.Scan(&n)
    s := make([]int,n)
    for i := 0; i < n; i++{
        f.Scan(&s[i])
		/*  var tmp int
		    f.Scan(&tmp)
			  s[i] = tmp  */
    }
    f.Scan(&last)
    var ans int
    for i := 0; i < n; i++{
        if s[i] == last{
            ans = i + 1
            break
        }
    }
    f.Println(ans)
}
