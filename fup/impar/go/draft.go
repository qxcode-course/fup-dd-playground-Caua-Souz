package main
import "fmt"
func main() {
      var p, d1, d2, dt int
    fmt.Scan(&p)
    fmt.Scan(&d1)
    fmt.Scan(&d2)
    fmt.Scan(&dt)

    dt = (d1+d2)
    if p != 1 && dt%2 == 0 {
        fmt.Println("0")
    } else if p != 0 && dt%2 == 1{
        fmt.Println("0")
    } else {
        fmt.Println("1")
    }

  
        }
    
   
    

