package main

import "fmt"

//Cap4-7 - Constantes
/* São valores imutáveis.
Podem ser tipadas ou não:
const oi = "Bom dia"
const oi string = "Bom dia"
As não tipadas só terão um tipo atribuido a elas quando forem usadas.
Ex. qual o tipo de 42? int? uint? float64?
Ou seja, é uma flexibilidade conveniente.
Na prática: int, float, string.
const x = y
const ( x = y ) */

const x = 10 //programa só vai identificar x como int, quando for usado, posso usar como int ou float64

var y = 10 //programa já identifica y como int, nesse momento de atibuição

var c int
var d float64

const (
	j = 10
	k = 20
	m = 30
)

func main() {
	c = x
	fmt.Println(c)
	d = x
	fmt.Println(d)
	fmt.Println(j, k, m)

}
