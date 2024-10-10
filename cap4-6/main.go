package main

import "fmt"

//Cap4-6 - Sistemas numéricos
/* Base-10: decimal, 0 e 9
Base-2: binário, 0 e 1
Base-16: hexadecimal, 0 e f
Demonstração em Go. */

var i uint16

func main() {
	a := 1000
	//exibir o valor de a em decimal, binário e hexdecimal
	fmt.Printf("%d\t%b\t%#x", a, a, a)

}
