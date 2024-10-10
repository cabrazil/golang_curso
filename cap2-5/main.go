package main

//explorando tipos
//se declarar uma variável x como int, será int até terminar o programa
//tipos em Go são estáticos
//Tipos de dados primitivos: disponíveis na linguagem nativamente como blocos básicos de construção
//int, string, bool
//Tipos de dados compostos: são tipos compostos de tipos primitivos, e criados pelo usuário
//slice, array, struct, map
import "fmt"

// O tipo pode ser deduzido pelo compilador
var y = 10

// Ou o tipo pode ser declarado explicitamente
var z int = 10
var w string = "isso é uma string"

// podemos declarar uma varável sem valor, mas o valor só poder ser atribuído dentro do codeblock
var k int

func main() {
	//variável k valor declarado dentro do codeblock
	k := 20
	fmt.Printf("k: %v, %T\n", k, k)

}
