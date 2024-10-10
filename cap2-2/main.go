package main

import "fmt"

func main() {
	//aqui inicia
	//notação pacote.identificador
	//essa função Println demonstra e põe na tela
	//func Println(a ...any) (n int, err error)
	//os 3 pontinhos ... indicam função variádica, trabalha com qualquer nro de argumentos
	numerodebytes, erros := fmt.Println("Hello World!", "Oi galera!", 100)
	//podemos usar _ p substituir o primeiro valor numerodebytes
	fmt.Println(numerodebytes, erros)

	//aqui finaliza
}
