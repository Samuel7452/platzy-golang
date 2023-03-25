package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// valor1 == valor2: Retorna TRUE si valor1 y valor2 son exactamente iguales.
	// valor1 != valor2: Retorna TRUE si valor1 es diferente de valor2.
	// valor1 < valor2: Retorna TRUE si valor1 es menor que valor2
	// valor1 > valor2: Retorna TRUE si valor1 es mayor que valor2
	// valor1 >= valor2: Retorna TRUE si valor1 es igual o mayor que valor2
	// valor1 <= valor2: Retorna TRUE si valor1 es menor o igual que valor2.

	myBool := false
	myBool2 := true
	fmt.Println(!myBool) // Esto retornará false

	if myBool == true {
		fmt.Print("es real")
	} else {
		fmt.Print("es falso")
	}

	//if and
	if myBool == false && myBool2 == true {
		fmt.Print("\nfalso y verdadero")

	} else {

		fmt.Print("\nno cumple")
	}

	//if or
	if myBool == false || myBool2 == true {
		fmt.Print("\nfalso o verdadero")
	}

	// convertir string a int y validar errores:

	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal("\n",err)
	}

	fmt.Print("\nValor:", value)
}
