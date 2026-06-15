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
		if alunos[i] % 2 == 0 { //se o índice atual for par...
			pares = append(pares, alunos[i] ) //...adicionar na lista de pares
		} else { //se não...
			impares = append(impares,alunos[i] ) //...adicionar na lista de impares
		}
	}

	fmt.Print("[ ") //abrir um parentese
	for _, v := range impares { //para cada "v" na lista de impares...
		fmt.Print(v, " ") //printe esse v "elemento" a seguir de um espaço
	}
	fmt.Print("]", "\n") //fechar o parentese e quebrar a linha.
	
	//mesma coisa só que para os pares :}
	fmt.Print("[ ")
	for _, w := range pares {
		fmt.Print(w, " ")
	}
	fmt.Print("]","\n")

}