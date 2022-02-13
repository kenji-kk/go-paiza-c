package main
import f "fmt"
func main(){
  var c, s string
  f.Scan(&c, &s)
  flag := false
  for i, _ := range s {
    if string(s[i]) == c {
      flag = true
    }
  }
  if flag {
    f.Println("YES")
  } else{
    f.Println("NO")
  }
}
