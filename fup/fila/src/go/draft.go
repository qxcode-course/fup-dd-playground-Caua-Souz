package main	
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	impares := make([]int, 0)
	pares := make([]int, 0)
	alunos := make([]int, n)
	for i := 0 ; i < n ; i++ {
		fmt.Scan(&alunos[i])
		if alunos[i] % 2 == 0 {
			alunos = append(pares, alunos[i] )
		} else {
			alunos = append(impares,alunos[i] )
			break
		}
	}
	fmt.Print(impares, "\n")
	fmt.Print(pares, "\n")

}