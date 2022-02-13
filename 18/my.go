package main
import f "fmt"

func main(){
    var n, m int
    f.Scan(&n)
    s1 := make([]string, n)
    for i := 0; i < n; i++{
        f.Scan(&s1[i])
    }
    f.Scan(&m)
    s2 := make([]string, m)
    for i := 0; i < m; i++{
        f.Scan(&s2[i])
    }
    for _, v1 := range s1{
        for _, v2 := range s2{
            flag := false
            for k := 0; k < len([]rune(v2)); k++{
                if string(v2[k]) == v1{
                    flag = true
                    break
                }
            }  
            if flag {
                f.Println("YES")
            } else{
                f.Println("NO")
            }
        }
    }
}
