package main
import "fmt"
func main() {
    var num int
    fmt.Scan(&num)

    leds := map[rune]int { //mapear cada led com seu respectivo número
        '1':2,
        '2':5,
        '3':5,
        '4':4,
        '5':5,
        '6':6,
        '7':3,
        '8':7,
        '9':6,
        '0':6,
    }

    for i := 0; i < num ; i++ {
        var str string
        fmt.Scan(&str)
        total := 0

        for _, v := range str {
            total += leds[v]
        }
        fmt.Printf("%d leds\n", total)
    }

}