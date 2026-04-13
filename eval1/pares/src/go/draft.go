        package main
        import "fmt"
        func main() {
            var a, b int
            fmt.Scan(&a)
            fmt.Scan(&b)
            i := 0
            if a > b {
            fmt.Println("invalido")
                } else {
                    for ; a <= b ; a++ {
                        if a%2 == 0 {
                            i += a 
                        }
                    }
                    fmt.Println(i)
                  }

             
            }
        
        
            
        


        

