package main
import "fmt"
func main() {
    var c1, l1, c2, l2 int
    fmt.Scan(&c1)
    fmt.Scan(&l1)
    fmt.Scan(&c2)
    fmt.Scan(&l2)
   

    area1 := c1 * l1 //cálculos da área 1 e 2
    area2 := c2 * l2

        if area1 > area2 {
            fmt.Println(area1)
        } else {
            fmt.Println(area2)

        }
}
