package main

//Conversão, não coerção
// Conversão de tipos é o que soa.
// Em Go não se diz casting, se diz conversion.
// a = int(b)
// ref/spec#Conversions, documentação

import "fmt"

// criando o próprio tipo
type hotdog int

var b hotdog = 101

func main() {

        x := 10
        fmt.Printf("%v\n", x)

		//isso é uma conversão de tipo
		x = int(b)
        fmt.Printf("%v", x)

        //Saida desse print
        //int e main.hotdog

}
