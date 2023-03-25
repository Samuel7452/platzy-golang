package main

import "fmt"

func main() {

	//printear una linea que contiene un salto de linea, exactamente igual que el print de python
	fmt.Println("Hola Mundo")

	//printear una linea igual que la anterior pero agregando variables dentro del string, %s sirve cuanso se sabe que la variable sera string y %d cuando se sabe que sera entero
	nombre := "platzy"
	cursos := 500

	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)

	// cuando no se sabe el tipo de variable que se obtendra se puede usar %v, pero la buena practica es la anterior
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf permite guardar el mismo resultado de los printf de antes pero en una variable, permitiendo printear el mensaje despues con println
	var message string = fmt.Sprintf("%v tiene más de %v cursos", nombre, cursos)
	fmt.Println(message)

	// fmt tambuien permite printear el tipo de una variable
	fmt.Printf("La variable 'nombre' es de tipo : %T\n", nombre)

}
