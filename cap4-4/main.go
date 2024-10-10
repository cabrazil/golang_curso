package main

import "fmt"

//Cap4-4 - Overflow
/* Um uint16, por exemplo, vai de 0 a 65535.
Que acontece se a gente tentar usar 65536?
Ou se a gente estiver em 65535 e tentar adicionar mais 1? */

var i uint16

func main() {
	i = 65535 //limite seria 0 - 65535, se forçar 65536 vai exbir overflow
	fmt.Println(i)
	i++
	fmt.Println(i)
	i++
	fmt.Println(i)

}
