package main
import "fmt"
func main() {
    var B, T, trapezio, mretangulo int
    fmt.Scan(&B)
    fmt.Scan(&T)
    


        trapezio = (B+T)* 70 / 2
        mretangulo = 70 * 80

        if trapezio > mretangulo {
            fmt.Println(1)
        } else if trapezio < mretangulo {
            fmt.Println(2)
        } else {
            fmt.Println(0)
        }

}
