    package main
    import "fmt"
    func main() {
        var p, s, e, p1, p2, op1 int
        fmt.Scan(&p)
        fmt.Scan(&s)
        fmt.Scan(&e)
        p2 = p1 - e
        op1 = p1
        for p1 = 0 ; p1 <=  p ; p1+= s {
            if op1 != p1 {
                p2 = (p1 - e)
            fmt.Println(p1, p2)
            }
            if p1 >= p {
                fmt.Println("saiu")
            }

            
        }
    
    }
