package main

import "fmt"

// operador curto de declaração := declara variável com tipagem automática
func main() {
	x := 10
	y := "Olá bom dia"
	z := 10.5
	//usando Printf
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n\n", z, z)
	//operador de atribuição
	//x := 20 isso não funciona
	x, w := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("w: %v, %T\n", w, w)
}
