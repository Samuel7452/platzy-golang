package main

import (
	"fmt"
)

func main() {

	//for condicional
	fmt.Print("\n*****************for condicional*************\n")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//for while
	fmt.Print("*****************for while*************\n")
	var counter int = 0
	for counter < 10 {
		fmt.Println("counter:", counter)
		counter++
	}

	//for for ever
	// counter_forever := 0
	// for {
	// 	fmt.Print(counter_forever)
	// 	counter_forever++
	// }

	// for condicional al reves

	fmt.Print("\n*****************for condicional invertido*************\n")
	for i := 10; i >= 0; i-- {
		fmt.Println(i)
	}
}
