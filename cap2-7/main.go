package main

//o pacote fmt

// Setup: strings, ints, bools.
// Strings: interpreted string literals vs. raw string literals.
// Rune literals.
// Em ciência da computação, um literal é uma notação para representar um valor fixo no código fonte.
// Format printing: documentação.
// Grupo #1: Print → standard out
// func Print(a ...interface{}) (n int, err error)
// func Println(a ...interface{}) (n int, err error)
// func Printf(format string, a ...interface{}) (n int, err error)
// Format verbs. (%v %T)
// Grupo #2: Print → string, pode ser usado como variável
// func Sprint(a ...interface{}) string
// func Sprintf(format string, a ...interface{}) string
// func Sprintln(a ...interface{}) string
// Grupo #3: Print → file, writer interface, e.g. arquivo ou resposta de servidor
// func Fprint(w io.Writer, a ...interface{}) (n int, err error)
// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
// func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

import "fmt"

func main() {
	//Strings: interpreted string literals, usando aspas duplas
	a := "oi bom dia\ncomo vai?\tespero que tudo bem."
	//raw string literals, usando acento crase
	b := `oi bom dia\ncomo vai?\tespero que tudo bem. não interpreta nada`

	x := "Oi"
	y := "bom dia"

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(x)
	fmt.Println(y)

	//Sprint, por exemplo, coloca o valor numa variável
	z := fmt.Sprint(x, " ", y)
	fmt.Println(z)

}
