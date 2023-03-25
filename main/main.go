package main

import (
	"fmt"
	"math"
)

func main() {
	// //Declaracion de constantes
	// const pi float64 = 3.14
	// const pi2 = 3.16

	// fmt.Println("Pi:", pi)
	// fmt.Println("Pi2:", pi2)

	// //Declaracion de varaibles
	// base := 12          //Primera forma
	// var altura int = 14 //Segunda forma
	// var area int        //Go no compila si las variables no son usadas

	// fmt.Println(base, altura, area)

	// //Zero values
	// //Go asigna vaalores a variables vacías
	// var a int
	// var b float64
	// var c string
	// var d bool

	// fmt.Println(a, b, c, d)

	//Ejercicio
	//Calcular el áera del cuadrado

	//************suma****************

	x := 10
	y := 50

	suma := x + y
	fmt.Println("la suma entre X y Y es:", suma)

	//*************resta******************

	x = 10
	y = 50

	resta := y - x
	fmt.Println("la resta entre X y Y es:", resta)


	//*******************Division*********************

	x = 10
	y = 50

	division := y / x
	fmt.Println("la division entre X y Y es:", division)

	//****************Multiplicacion***********

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)

	// ************obtener pi*******************

	pi := math.Pi
	fmt.Println("pi es igual a:", pi)

	// incrementar variable

	x++
	fmt.Println("X incrementado es igual a:", x)

	// decrementar variable

	x--
	x--
	fmt.Println("X decrementado es igual a:", x)



}


