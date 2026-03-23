package main
import "fmt"
func main() {
    var wifi, login, admin bool
    fmt.Scan(&wifi)
    fmt.Scan(&login)
    fmt.Scan(&admin)

    if wifi == false {
        fmt.Println("you must connect to wifi") //mensagem se o wifi for falso
    } else if login == false {
        fmt.Println("you need to login first") //mensagem se o login for falso
    } else if admin == false {
        fmt.Println("you must to login as admin") //mensagem se o login por admin for falso
    } else {
        fmt.Println("done") //se todos forem verdadeiros
    }
    
   
}
