package main

//valor zero
//Declaração vs. inicialização vs. atribuição de valor. Variáveis: caixas postais.
//Cada tipo em GO vai ter um valor zero diferente
//Os zeros:
//ints: 0
//floats: 0.0
//booleans: false
//strings: ""
//pointers, functions, interfaces, slices, channels, maps: nil
//Use := sempre que possível.
//Use var para package-level scope.

import "fmt"

var a int
var b float64
var c string
var d bool

func main() {

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

}
