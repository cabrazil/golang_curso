package main

import "fmt"

//Cap4-5 - Tipo string (cadeias de caracteres)
/* Strings são sequencias de bytes.
Imutáveis.
Uma string é um "slice of bytes" (ou, em português, uma fatia de bytes).
Na prática:
%v %T
Raw string literals
Conversão para slice of bytes: []byte(x)
%#U, %#x */

func main() {

	//s := "Hello, playground" Aqui só ASCII
	//colocando caracteres UTF8
	s := "ascii éøâ 香"
	//sb := []byte(s) //slice de bytes

	//com range no string vai obter caracter por caracter
	for _, v := range s {
		fmt.Printf("%b - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println("")

	//com for ... loop vai obter byte por byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])

	}
}
