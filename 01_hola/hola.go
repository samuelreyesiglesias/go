package main

import "fmt"
//names fmt is the name of the package that contains the function Println
//include the package fmt in the program

func main() {
	mostrarHolaMundo();
}
//funcion para mostrar hola mundo
func mostrarHolaMundo(){
	fmt.Println("Hola Mundo en GoLang");
}
//execute in terminal: go run hola.go