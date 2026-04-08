package main
import "fmt"
func main() {
    var num int
    fmt.Scan(&num)

   


    for i := 1 ; i <= num; i+= 2 {
        if i%2 == 0 {
            continue
        }
     fmt.Println(i)
         }
    for i2 := num; i2 >= 0 ; i2-- {
        if i2%2 != 0 {      
            continue
        }
     fmt.Println(i2)
    }
          
        }
    


