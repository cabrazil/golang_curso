package main

import "fmt"

//Cap4-8 - Iota
/* golang.org/ref/spec
Numa declaração de constantes, o identificador iota representa números sequenciais.
Na prática.
iota, iota + 1, a = iota b c, reinicia em cada const, _ */

const (
	a = iota + 1
	_
	c
	_
	e
	f
)

func main() {
	fmt.Println(a, c, e, f)

}
