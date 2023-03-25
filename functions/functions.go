package main

import "fmt"

func normalFunction(){
	fmt.Println("hola mundo")
}

func parametersFunction(number int){
	fmt.Printf("hola mundo %d \n", number)
}

func multiple_argumentsFunction(number1, number2 int, str string){
	fmt.Println(number1, number2, str)
}

func return_multiple_argumentsFunction(number1, number2 int, str string) (a, b int, c string){
	return number1, number2, str
}

func main() {
	normalFunction()

	parametersFunction(1)
	parametersFunction(2)
	parametersFunction(3)

	multiple_argumentsFunction(1,2,"hola")

	number1, number2, str := return_multiple_argumentsFunction(1,2,"hola")
	println("multiple returns:", number1, number2, str)

	// ignorar una variable de una funcion que retorna varias variables:

	number1, number2, _ = return_multiple_argumentsFunction(1,2,"hola")
	println("multiple returns ignorando str:", number1, number2)


}