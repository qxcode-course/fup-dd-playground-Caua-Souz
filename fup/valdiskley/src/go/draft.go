package main
import "fmt"
func main() {
    var letra string
    var rot int
    fmt.Scan(&letra, &rot)

    ind := int(letra[0] - 'a')
    ind2 := (ind + rot) % 26

    if ind2< 0 {
        ind2 += 26
    }
    result := byte(ind2 + 'a')
    fmt.Printf("%c\n", result)
}