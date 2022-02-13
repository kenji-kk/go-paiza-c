package main
import (
    f "fmt"
    //sc "strconv"
    )
func main(){
    var str string
    f.Scan(&str)
/*  one, _ := sc.Atoi(string(str[0]))
    two, _ := sc.Atoi(string(str[1]))
    three, _ := sc.Atoi(string(str[2]))
    four, _ := sc.Atoi(string(str[3]))
    front := sc.Itoa(one + four)
    back := sc.Itoa(two + three)  */
    ffront := ((str[0] - '0') + (str[3] - '0'))
    bback := ((str[1] - '0') + (str[2] - '0'))
    f.Print(ffront)
    f.Print(bback)
}
