 package main
import "fmt"
func main() {
    var novo, qfilhos, i int
    fmt.Scan(&novo)
    fmt.Scan(&qfilhos)
    i = 0

    for; i < qfilhos; i++ {
    fmt.Println(novo)
    novo += 2
    }
}
