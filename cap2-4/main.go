package main

//palavra chave var
import "fmt"

// aqui a variável y é package level scope
var y = 10

func main() {
	//variável z está declarada dentro do codeblock, fora ninguém enxerga
	z := 20
	qualquercoisa(z)

}

func qualquercoisa(x int) {
	fmt.Println(y)
	fmt.Println(x)
}
