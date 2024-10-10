package main

//Criando seu próprio tipo
// Grande parte dos aspectos mais avançados de Go dependem quase que exclusivamente de tipos.
// Como fundação para estas ferramentas, vamos aprender a declarar nossos próprios tipos.
// Revisando: tipos são fixos. Uma vez declarada uma variável como de um certo tipo, isso é imutável.
// type hotdog int → var b hotdog (main hotdog)
// Uma variável de tipo hotdog não pode ser atribuida com o valor de uma variável tipo int, mesmo que este seja o tipo subjacente de hotdog.

import "fmt"

// criando o próprio tipo
type hotdog int

var b hotdog

func main() {

	x := 10
	fmt.Printf("%T\n", x)
	fmt.Printf("%T", b)

	//Saida
	//int e main.hotdog

}
