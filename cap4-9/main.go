package main

import "fmt"

//Cap4-9 - Deslocamento de bits
/* Deslocamento de bits é quando deslocamos digitos binários para a esquerda ou direita. */

// Exemplo de uso, demonstrar valores de KB, MB, GB e TB
const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	y := 24
	x := y << 2
	z := y >> 2
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", z)

	fmt.Println("")
	fmt.Println("binary\t\t\t\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)

}
